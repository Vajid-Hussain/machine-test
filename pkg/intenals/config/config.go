package config

import "github.com/spf13/viper"

type Config struct {
	DB     SqlDatabase
	JWT    JWTConfig
	Server Server
}

type SqlDatabase struct {
	Name     string `mapstructure:"Name"`
	Host     string `mapstructure:"Host"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Port     string `mapstructure:"Port"`
	URL      string `mapstructure:"URL"`
}

type JWTConfig struct {
	SecretKeyUser  string `mapstructure:"JWT_SECRET_USER"`
	SecretKeyAdmin string `mapstructure:"JWT_SECRET_ADMIN"`
	ExpirationTime int64  `mapstructure:"JWT_EXPIRATION"`
}

type Server struct {
	Port string `mapstructure:"ServerPort"`
}

func InitConfig() (*Config, error) {
	var c Config
	viper.SetConfigType("json")
	viper.SetConfigName("env")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c.DB)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c.JWT)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c.Server)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
