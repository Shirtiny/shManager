package auth

import (
	"crypto/rsa"
	"fmt"
	"shManager/util"

	"github.com/dgrijalva/jwt-go"
)

// Claims 自定jwt载体
type Claims struct {
	jwt.StandardClaims
	Foo string `json:"foo"`
}

// RsaKey rsa密钥对bytes
type RsaKey struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// Rsa 密钥对
var Rsa *RsaKey

// InitRsa 初始化rsa密钥对
func InitRsa(privateKeyStr, publicKeyStr string) {
	Rsa = &RsaKey{
		PrivateKey: util.LoadRSAPrivateKeyFromBytes([]byte(privateKeyStr)),
		PublicKey:  util.LoadRSAPublicKeyFromBytes([]byte(publicKeyStr)),
	}
}

// InitRsaByFile 通过文件初始化签名和验证jwt需要用到的rsa密钥对
func InitRsaByFile(privateKeyFilePath, publicKeyFilePath string) {
	Rsa = &RsaKey{
		PrivateKey: util.LoadRSAPrivateKeyFromDisk(privateKeyFilePath),
		PublicKey:  util.LoadRSAPublicKeyFromDisk(publicKeyFilePath),
	}
}

// CreateJwt 创建jwt SigningMethodRS256
func CreateJwt() string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, Claims{
		Foo: "my Foo",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Unix() + 60,
			Issuer:    "test",
		},
	})
	ss, err := token.SignedString(Rsa.PrivateKey)
	fmt.Printf("生成jwt： %v %v", ss, err)
	return ss
}

// ParseJwt 解析jwt RS
func ParseJwt(tokenString string) *Claims {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Rsa.PublicKey, nil
	})
	claims, ok := token.Claims.(*Claims)

	if ok && token.Valid {
		fmt.Printf("解析jwt：%+v", claims)
		return claims
	}
	fmt.Println(err)
	return claims
}
