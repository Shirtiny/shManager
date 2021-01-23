package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	// "os"

	"github.com/dgrijalva/jwt-go"
)

// RsaGenerateKeyBytes 生成私钥和公钥 file为true时 生成文件
func RsaGenerateKeyBytes() ([]byte, []byte) {
	// GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	// MarshalPKCS1PrivateKey将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	derPrivateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// Block代表PEM编码的结构, 对其进行设置
	privateBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derPrivateStream,
	}
	// 编码到内存中 私钥 []byte
	privateKeyBytes := pem.EncodeToMemory(&privateBlock)

	// 生成私钥文件
	// privateFile, err := os.Create("../private.pem")
	// defer privateFile.Close()
	// if err != nil {
	// }
	// err = pem.Encode(privateFile, &privateBlock)
	// if err != nil {
	// 	panic(err)
	// }

	// 公钥
	publicKey := privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// Block代表PEM编码的结构,
	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	// 编码到内存中 公钥 []byte
	publicKeyBytes := pem.EncodeToMemory(&publicBlock)

	// 生成公钥文件
	// publicFile, err := os.Create("../public.pem")
	// defer publicFile.Close()
	// if err != nil {
	// 	panic(err)
	// }
	// err = pem.Encode(publicFile, &publicBlock)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("生成rsa密钥对：")
	fmt.Println("私钥：\n", string(privateKeyBytes))
	fmt.Println("公钥：\n", string(publicKeyBytes))

	return privateKeyBytes, publicKeyBytes
}

// LoadRSAPrivateKeyFromDisk 从文件载入私钥
func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

// LoadRSAPublicKeyFromDisk 从文件载入公钥
func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

// RsaSignWithSha256 签名 测试用
func rsaSignWithSha256(data []byte, privateKeyBytes []byte) (signature []byte) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		panic(errors.New("私钥为空"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("私钥解析出错", err)
		panic(err)
	}

	signature, err = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("签名出错: %s\n", err)
		panic(err)
	}

	return signature
}

// ParseRSAPrivateKeyFromPEM 从pem中 解析出privateKey 测试用
func parseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	return privateKey, err
}

// ParseRSAPublicKeyFromPEM 从pem中 解析出publicKey 测试用
func parseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	return publicKey, err
}

// 测试入口函数
func genRSA() {
	privateKeyBytes, publicKeyBytes := RsaGenerateKeyBytes()
	// signature := rsaSignWithSha256([]byte("一段消息"), privateKeyBytes)
	// fmt.Println("签名成功：", string(signature))
	privateKey, _ := parseRSAPrivateKeyFromPEM(privateKeyBytes)
	fmt.Println("私钥内存：", privateKey)
	publickKey, _ := parseRSAPublicKeyFromPEM(publicKeyBytes)
	fmt.Printf("公钥内存：%+v\n", publickKey)
}
