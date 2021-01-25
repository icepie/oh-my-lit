package conf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/spf13/viper"
)

// JWConfig 教务在线配置
type JWConfig struct {
	UserName string
	PassWord string
	RefInt   int
}

// Config 基础配置
type Config struct {
	Host string
	Port int
	JW   JWConfig
}

// ProConf 新建实例
var ProConf = new(Config)

// MAIN 主地址配置
var MAIN string

// initConfig 初始化配置
func initConfig(cpath string) error {

	b, err := yaml.Marshal(ProConf)
	if err != nil {
		return err
	}

	f, err := os.Create(cpath)
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString(string(b))

	return nil
}

// INIT 初始化函数
func INIT() {

	// 取项目地址
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	cpath := path + "/conf"
	cfpath := cpath + "/conf.yaml"

	viper.AddConfigPath(cpath)  // 设置读取的文件路径
	viper.SetConfigName("conf") // 设置读取的文件名
	viper.SetConfigType("yaml") // 设置文件的类型

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		log.Println(errors.New("Can not read the config file, will recreate it! "))
		// 初始化配置
		ProConf.Port = 8088
		ProConf.JW.RefInt = 1800
		if err = initConfig(cpath + "/conf.yaml"); err != nil { // 重新初始化配置文件
			log.Fatalln(err)
		}
		log.Println(errors.New("Please edit the \"" + cfpath + "\"， then restart lit-edu-go"))
		os.Exit(1)
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(ProConf); err != nil {
		log.Fatalln(err)
	}

	// 拼接最终运行地址
	MAIN = fmt.Sprintf("%s:%d", ProConf.Host, ProConf.Port)
}
