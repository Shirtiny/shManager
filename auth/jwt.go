package auth

import (
	"crypto/rsa"
	"fmt"
	"shManager/serializer"
	"shManager/util"

	"github.com/dgrijalva/jwt-go"
)

const (
	// jwt多少秒后过期
	timeout = 3600
)

// Claims 自定jwt载体
type Claims struct {
	jwt.StandardClaims
	serializer.User
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
func CreateJwt(user serializer.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Unix() + timeout,
			Issuer:    "test",
		},
	})
	jwt, err := token.SignedString(Rsa.PrivateKey)
	if err != nil {
		return "", err
	}
	return jwt, err
}

// ParseJwt 解析jwt RS
func ParseJwt(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Rsa.PublicKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		fmt.Printf("解析jwt：%+v", claims)
		return nil, err
	}
	return claims, nil
}
