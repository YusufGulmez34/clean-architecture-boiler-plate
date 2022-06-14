package config

import "github.com/spf13/viper"

type Configuration struct {
	SystemConfig SystemConfig `yaml:"system" mapstructure:"system"`
	DbConfig     DbConfig     `yaml:"db" mapstructure:"db"`
}

type SystemConfig struct {
	DbDriver        string `yaml:"dbdriver"`
	TokenExpiryTime int    `yaml:"tokenexpirytime"`
	TokenSecretKey  string `yaml:"tokensecretkey"`
}

type DbConfig struct {
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

func (c *Configuration) ReadConfigFile() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err.Error())
	}
}
