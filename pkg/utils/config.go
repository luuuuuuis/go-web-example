/*
@File    :   config.go
@Time    :   2024/07/18 21:38:01
@Author  :   Luis
@Contact :   luis9527@163.com
*/



package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/pkg/errors"
)

const (
	ConfigFilePath = "./conf"
	ConfigFileName = "app"
	ConfigFileType = "yaml"
)

var (
	Config *viper.Viper
)

func InitConfig() {
	configParsed, err := ParseConfig(ConfigFilePath, ConfigFileName, ConfigFileType)
	if err != nil {
		fmt.Printf("read config error,%v \n", err)
		os.Exit(1)
	}
	Config = configParsed
}

/**
 * 通用的读取配置文件的方法，传入路径和文件名以及类型，返回一个Viper的指针
 */
func ParseConfig(filePath, fileName, configType string) (config *viper.Viper, err error) {
	config = viper.New()
	config.AddConfigPath(filePath)
	config.SetConfigName(fileName)
	config.SetConfigType(configType)
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Not found config file: " + err.Error())
		} else {
			return nil, errors.New("read config file error: " + err.Error())
		}
	}
	// fmt.Printf("config.GetStringMap(\"server\"): %v\n", config.GetString("server"))
	// fmt.Printf("config.GetStringMap(\"log\"): %v\n", config.GetString("log"))
	return
}
