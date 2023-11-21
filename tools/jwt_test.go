package tools

import (
	"fmt"
	"testing"
)

func TestGenerateJWT1(t *testing.T) {
	claims := MyCustomClaims{
		"liangping",
		1,
		JWTRegisteredClaims(),
	}
	ss, _ := GenerateJWT(claims)
	info, err := ParseJWT(ss)
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}
