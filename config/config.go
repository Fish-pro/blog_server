package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Debug   bool         `yaml:"debug"`
	Server  *ServerConf  `yaml:"server"`
	Mongodb *MongodbConf `yaml:"mongodb"`
	Redis   *RedisConf   `yaml:"redis"`
	Auth    *AuthConf    `yaml:"auth"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s *ServerConf) GetServerAddr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type MongodbConf struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	DbName string `yaml:"dbname"`
}

type RedisConf struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Passwd string `yaml:"passwd"`
	Db     int    `yaml:"db"`
}

type AuthConf struct {
	Mongodb *MongodbConf `yaml:"mongodb"`
	Redis   *RedisConf   `yaml:"redis"`
}

func LoadConf() *Config {
	var err error
	var port string
	var cfile string

	flag.StringVar(&cfile, "c", "", "-c /config/file/path")
	flag.StringVar(&port, "p", "", "-p 8200")

	flag.Parse()

	if cfile == "" {
		fmt.Println("缺少运行的配置文件, 使用 -c /config/file/path")
		os.Exit(1)
	}

	viper.SetConfigFile(cfile)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("解析配置文件失败", err.Error())
		os.Exit(1)
	}

	var c Config

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("反序列化配置文件失败", err.Error())
		os.Exit(1)
	}

	return &c
}
