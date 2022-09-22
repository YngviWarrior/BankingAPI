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
	GenerateJWT(ip string) (tokenString string, err error)
	VerifyJWT(http.ResponseWriter, *http.Request) error
}

type Claims struct {
	Ip string `json:"ip"`
	jwt.RegisteredClaims
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func (*Jwt) GenerateJWT(ip string) (tokenString string, err error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Ip: ip,
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

func (*Jwt) VerifyJWT(writer http.ResponseWriter, request *http.Request) error {
	if len(request.Header["Auth-Token"]) == 0 {
		writer.WriteHeader(http.StatusUnauthorized)
		return errors.New("no auth token")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(request.Header["Auth-Token"][0], claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			writer.WriteHeader(http.StatusUnauthorized)
			return err
		}

		writer.WriteHeader(http.StatusBadRequest)
		return err
	}

	if !token.Valid {
		writer.WriteHeader(http.StatusUnauthorized)
		return err
	}

	return nil
}
