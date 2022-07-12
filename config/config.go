package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Configuration struct {
		DatabaseSettings
		JwtSettings
	}
	// 数据库配置
	DatabaseSettings struct {
		DatabaseURI  string
		DatabaseName string
		Username     string
		Password     string
	}
	// jwt配置
	JwtSettings struct {
		SecretKey string
	}
	// reader
	ConfigReader struct {
		configFile string
		v          *viper.Viper
	}
)

// 获得所有配置
func GetAllConfigValues(configFile string) (configuration *Configuration, err error) {
	// 获取配置
	cfgReader := newConfigReader(configFile)

	// 读取配置并解析
	if err = cfgReader.v.ReadInConfig(); err != nil {
		fmt.Printf("配置文件读取失败 : %s", err)
		return nil, err
	}

	err = cfgReader.v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("解析配置文件到结构体失败 : %s", err)
		return nil, err
	}

	return configuration, err
}

// 实例化configReader
func newConfigReader(configFile string) (cfgReader *ConfigReader) {
	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)
	cfgReader = &ConfigReader{
		configFile: configFile,
		v:          v,
	}
	return cfgReader
}
