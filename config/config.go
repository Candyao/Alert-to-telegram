package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Service struct{
		Port string `yaml:"port"`
		Token string `yaml:"token"`
	}`yaml:"service"`
	App struct{
		ChatID string `mapstructure:"chat_id"`
		BotToken string `mapstructure:"bot_token"`
	}`yaml:"app"`
	TimeZone string `mapstructure:"time_zone"`
	Templatepath string `mapstructure:"template_path"`
	TimeOutdata string `mapstructure:"time_outdata"`
}

type Secret struct {
	Mysql struct{
		Driver string `yaml:"driver"`
		Connections struct{
			Database string `yaml:"database"`
			Passwd string `yaml:"passwd"`
			Port string `yaml:"port"`
			Host string `yaml:"host"`
			User string `yaml:"user"`
		}`yaml:"connections"`
		Logmod bool `mapstructure:"log_mod"`
		MaxIdleConns int `mapstructure:"max_idle_conns"`
		MaxOpenConns int `mapstructure:"max_open_conns"`

	}`yaml:"mysql"`
}

type Public struct {
	 Secret
	 Config
}

var Configmanager=&Public{}

func (public *Public) LoadConfig (configFile,secretFile *string) {
	conf := viper.New()
	secret:= viper.New()
	if configFile == nil {
		conf.SetConfigFile("./config/config.yaml")
	} else {
		conf.SetConfigFile(*configFile)
	}
	if secretFile == nil {
		secret.SetConfigFile("./secret/secret.yaml")
	} else {
		secret.SetConfigFile(*secretFile)
	}
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig();err != nil {
		log.Printf("read config error %v",err)
	}
	if err := conf.Unmarshal(&public.Config); err != nil {
		log.Printf("unmarshal config error %v",err)
	}
	secret.SetConfigType("yaml")
	if err := secret.ReadInConfig();err != nil {
		log.Printf("read secret error %v",err)
	}
	if err := secret.Unmarshal(&public.Secret);err != nil {
		log.Printf("unmarshal secert error %v",err)
	}
}
