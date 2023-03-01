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
	GenerateJWT(id uint64, admin bool, ip string) (accessTokenString string, refreshTokenString string, err error)
	VerifyJWT(http.ResponseWriter, *http.Request) (uint64, error)
	VerifyToRefreshJWT(accessToken, refreshToken string) (uint64, string, error)
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

func (*Jwt) GenerateJWT(id uint64, admin bool, ip string) (accessTokenString string, refreshTokenString string, err error) {
	accessExpirationTime := time.Now().Add(10000 * time.Minute)

	claims := &Claims{
		Data: Info{id, ip, admin},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: accessExpirationTime},
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err = accessToken.SignedString(secretKey)

	if err != nil {
		return
	}

	refreshExpirationTime := time.Now().Add(10 * time.Minute)

	claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: refreshExpirationTime},
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err = refreshToken.SignedString(secretKey)

	if err != nil {
		return
	}

	return
}

func (j *Jwt) VerifyJWT(writer http.ResponseWriter, request *http.Request) (userId uint64, err error) {
	if request.Header.Get("Auth-Token") == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		return 0, errors.New("no token found")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(request.Header.Get("Auth-Token"), claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || token == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return 0, err
	}

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
		return userId, err
	}

	return userId, nil
}

func (j *Jwt) VerifyToRefreshJWT(authToken string, refreshToken string) (userId uint64, ip string, err error) {
	if authToken == "" && refreshToken == "" {
		return 0, "", errors.New("no token found")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || token == nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		userId = claims.Data.Id
		ip = claims.Data.Ip
	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, "", err
		}

		return 0, "", err
	}

	// if !token.Valid {
	// 	return userId, ip, nil
	// }

	return userId, ip, nil
}
