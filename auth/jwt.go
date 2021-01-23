package auth

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	jwt.StandardClaims
	Foo string `json:"foo"`
}

// CreateJwt 创建jwt SigningMethodRS256
func CreateJwt(privateKey *rsa.PrivateKey) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims{
		Foo: "my Foo",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Unix() + 800,
			Issuer:    "test",
		},
	})
	ss, err := token.SignedString(privateKey)
	fmt.Printf("生成jwt： %v %v", ss, err)
	return ss
}

// ParseJwt 解析jwt RS
func ParseJwt(tokenString string, publicKey *rsa.PublicKey) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if claims, ok := token.Claims.(*claims); ok && token.Valid {
		fmt.Printf("解析jwt：%+v", claims)
	} else {
		fmt.Println(err)
	}
}
