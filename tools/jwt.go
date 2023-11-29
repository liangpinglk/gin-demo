package tools

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyCustomClaims struct {
	Name string `json:"name"`
	ID   uint    `json:"id"`
	jwt.RegisteredClaims
}

func JWTRegisteredClaims() jwt.RegisteredClaims {
	registeredClaims := jwt.RegisteredClaims{
		// A usual scenario is to set the expiration time relative to the current time
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "gin-demo",
		Subject:   "gin-demo",
		//ID:        "1",
		//Audience: []string{"somebody_else"},
	}
	return registeredClaims
}

func ParseJWT(tokenString string) (any, error) {

	projectSigningKey := []byte("gin-demo")
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return projectSigningKey, nil
	})
	if err != nil {
		return "", err
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims, nil
	} else {
		return "", &GinDemoError{"unknown claims type, cannot proceed"}
	}
}

func GenerateJWT(claims MyCustomClaims) (string, error) {
	projectSigningKey := []byte("gin-demo")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(projectSigningKey)
	return ss, err
}
