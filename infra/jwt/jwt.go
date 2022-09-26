package jwt

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct{}

type JwtInterface interface {
	GenerateJWT(id uint64, admin bool, ip string) (tokenString string, err error)
	VerifyJWT(http.ResponseWriter, *http.Request) (uint64, error)
}

type Info struct {
	Id   uint64 `json:"id"`
	Ip   string `json:"ip"`
	Type bool   `json:"type"`
}

type Claims struct {
	Data Info
	jwt.RegisteredClaims
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func (*Jwt) GenerateJWT(id uint64, admin bool, ip string) (tokenString string, err error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Data: Info{id, ip, admin},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(secretKey)

	if err != nil {
		return
	}

	return
}

func (*Jwt) VerifyJWT(writer http.ResponseWriter, request *http.Request) (userId uint64, err error) {
	if request.Header["Auth-Token"][0] == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		return 0, errors.New("no auth token")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(request.Header["Auth-Token"][0], claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		userId = claims.Data.Id
	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			writer.WriteHeader(http.StatusUnauthorized)
			return 0, err
		}

		writer.WriteHeader(http.StatusBadRequest)
		return 0, err
	}

	if !token.Valid {
		writer.WriteHeader(http.StatusUnauthorized)
		return 0, err
	}

	return userId, nil
}
