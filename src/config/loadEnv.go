package config

import "github.com/spf13/viper"

type GlobalConfig struct {
	NodeEnv    string `mapstructure:"NODE_ENV"`
	Port       string `mapstructure:"PORT"`
	HostUrl    string `mapstructure:"HOST_URL"`
	PublicKey  string `mapstructure:"PUBLIC_KEY"`
	ApiContext string `mapstructure:"API_CONTEXT"`
}

var config *GlobalConfig

func LoadEnv(path string) (envConfig GlobalConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&envConfig)
	if err != nil {
		return
	}

	config = &envConfig
	return
}

func GetConfig() *GlobalConfig {
	return config
}
