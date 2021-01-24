package conf

import (
	"io/ioutil"
	"os"

	"shManager/auth"
	"shManager/model"
	"shManager/util"

	"github.com/joho/godotenv"
)

// 密钥对是否存在数据库中
var keyInDb = true

// 密钥对文件的存放目录
var keyDirPath = "./tmp/"

// 私钥文件名
var privateKeyFileName = "private.pem"

// 公钥文件名
var publicKeyFileName = "public.pem"

// PathExists 判断路径是否存在
func PathExists(path string) (bool, error) {
	// 如果需要安全判断,  可以使用 os.Stat 配合 os.IsNotExist
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 在传入的目录生成公钥或私钥的文件
func generateKeyFiles(dirPath string) {
	// 如果文件夹不存在
	if dirExist, _ := PathExists(dirPath); !dirExist {
		// 创建文件夹
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// 如果私钥文件已经存在 不再继续操作
	if privateKeyFileExist, _ := PathExists(dirPath + privateKeyFileName); privateKeyFileExist {
		return
	}
	// 生成用于签发jwt的密钥对
	privateKeyBytes, publicKeyBytes := util.RsaGenerateKeyBytes()
	// 当私钥文件不存在时, 才生成文件 覆盖原
	err := ioutil.WriteFile(dirPath+privateKeyFileName, privateKeyBytes, 0666)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(dirPath+publicKeyFileName, publicKeyBytes, 0666)
	if err != nil {
		panic(err)
	}
}

// Init 初始化项目配置
func Init() {
	// 读取环境变量
	godotenv.Load()
	// 连接数据库
	model.ConnectDatabase("lab_1591053723:50c712fa6981_#@Aa@tcp(rm-bp1oo27t8762xhlob0o.mysql.rds.aliyuncs.com:3306)/shmysql?charset=utf8&parseTime=True&loc=UTC")
	// 是否从数据库中读取密钥对
	if keyInDb {
		// 从数据库生成密钥对
		privateKeyStr, publicKeyStr := model.KeyGenerate()
		// 通过字符串初始化密钥对
		auth.InitRsa(privateKeyStr, publicKeyStr)
	} else {
		// 生成密钥对文件
		generateKeyFiles(keyDirPath)
		// 从文件初始化jwt的rsa密钥对
		auth.InitRsaByFile(keyDirPath+privateKeyFileName, keyDirPath+publicKeyFileName)
	}
}
