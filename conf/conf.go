package conf

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

// JWConfig 教务在线配置
type JWConfig struct {
	UserName string
	PassWord string
}

// Config 基础配置
type Config struct {
	Host string
	Port int
	MAIN string
	JW   JWConfig
}

// ProConf 新建实例
var ProConf = new(Config)

// INIT 初始化函数
func INIT() {

	// 取项目地址
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(path + "/conf") // 设置读取的文件路径
	viper.SetConfigName("conf")         // 设置读取的文件名
	viper.SetConfigType("yaml")         // 设置文件的类型

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		log.Fatalln(err)
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(ProConf); err != nil {
		log.Fatalln(err)
	}

	// 拼接最终运行地址
	ProConf.MAIN = fmt.Sprintf("%s:%d", ProConf.Host, ProConf.Port)
}
