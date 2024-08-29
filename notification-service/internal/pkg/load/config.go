package load

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type EmailConfig struct {
	From     string
	Password string
}

type Config struct {
	Service ServiceConfig
	Email   EmailConfig
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Service: ServiceConfig{
			Host: viper.GetString("service.host"),
			Port: viper.GetInt("service.port"),
		},
		Email: EmailConfig{
			From: viper.GetString("email.from"),
			Password: viper.GetString("email.password"),
		},
	}
	return &cfg, nil
}
