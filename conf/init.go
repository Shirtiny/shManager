package conf

import (
	"io/ioutil"
	"os"
	"shManager/model"
	"shManager/util"

	"github.com/joho/godotenv"
)

// 在传入的目录生成公钥或私钥的文件
func generateKeyFiles(dirPath string, privateKeyBytes []byte, publicKeyBytes []byte) {
	// 判断私钥是否已经存在 如果需要安全判断, 可以使用 os.Stat 配合 os.IsNotExist
	if _, err := os.Stat(dirPath + "private.pem"); os.IsNotExist(err) {
		// 当文件不存在, 才写文件
		err := ioutil.WriteFile(dirPath+"private.pem", privateKeyBytes, 0666)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(dirPath+"public.pem", publicKeyBytes, 0666)
		if err != nil {
			panic(err)
		}
	}
}

// Init 初始化项目配置
func Init() {
	// 读取环境变量
	godotenv.Load()
	// 连接数据库
	model.ConnectDatabase("lab_1591053723:50c712fa6981_#@Aa@tcp(rm-bp1oo27t8762xhlob0o.mysql.rds.aliyuncs.com:3306)/shmysql?charset=utf8&parseTime=True&loc=UTC")
	// 生成用于签发jwt的密钥对
	privateKeyBytes, publicKeyBytes := util.RsaGenerateKeyBytes()
	// 生成密钥对文件
	generateKeyFiles("./", privateKeyBytes, publicKeyBytes)
}
