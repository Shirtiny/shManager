package model

import (
	"errors"
	"shManager/util"
)

// Key 密钥
type Key struct {
	ID         uint   `gorm:"primary_key"`
	PrivateKey string `gorm:"type:text;not null;"`
	PublicKey  string `gorm:"type:text;not null"`
}

// KeyExist key是否已经存在
func KeyExist() bool {
	count := 0
	DB.Table("keys").Count(&count)
	return count > 0
}

// KeyGet 查询密钥
func KeyGet() (string, string) {
	keyPair := Key{}
	// 直接返回第一个
	DB.First(&keyPair)
	return keyPair.PrivateKey, keyPair.PublicKey
}

// KeyAdd 添加密钥
func KeyAdd(pri string, pub string) error {
	exist := KeyExist()
	if exist {
		return nil
	}
	// key不存在时
	if err := DB.Create(&Key{
		PrivateKey: pri,
		PublicKey:  pub,
	}).Error; err != nil {
		return errors.New("插入密钥对失败：" + err.Error())
	}
	return nil
}

// KeyGenerate 根据情况在数据库存入密钥对 返回拿到的key值
func KeyGenerate() (string ,string) {
	exist := KeyExist()
	// 如果数据库中已有key 返回库中第一个key值
	if exist {
		return KeyGet()
	}
	// 不存在则 生成用于签发jwt的密钥对
	privateKeyBytes, publicKeyBytes := util.RsaGenerateKeyBytes()
	privateKeyStr := string(privateKeyBytes)
	publicKeyStr := string(publicKeyBytes)
	// 添加到数据库中
	err := KeyAdd(privateKeyStr, publicKeyStr)
	if err != nil {
		panic(err)
	}
	return privateKeyStr, publicKeyStr
}
