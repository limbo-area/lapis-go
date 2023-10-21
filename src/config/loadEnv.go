package config

import "github.com/spf13/viper"

type GlobalConfig struct {
	NodeEnv    string `mapstructure:"NODE_ENV"`
	Port       string `mapstructure:"PORT"`
	HostUrl    string `mapstructure:"HOST_URL"`
	DBUrl      string `mapstructure:"DB_URL"`
	PublicKey  string `mapstructure:"PUBLIC_KEY"`
	PrivateKey string `mapstructure:"PRIVATE_KEY"`
	ApiContext string `mapstructure:"API_CONTEXT"`
	ApiKey     string `mapstructure:"API_KEY"`
	AuthPrefix string `mapstructure:"AUTH_PREFIX"`
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
