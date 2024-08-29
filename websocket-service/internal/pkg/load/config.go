package load

import "github.com/spf13/viper"

type KafkaConfig struct {
	Host  string
	Port  string
	Topic string
}

type ServiceConfig struct {
	Host string
	Port int
}

type Config struct {
	Kafka    KafkaConfig
	Service  ServiceConfig
	CertFile string
	KeyFile  string
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Kafka: KafkaConfig{
			Host:  viper.GetString("kafka.host"),
			Port:  viper.GetString("kafka.port"),
			Topic: viper.GetString("kafka.topic"),
		},

		Service: ServiceConfig{
			Host: viper.GetString("service.host"),
			Port: viper.GetInt("service.port"),
		},

		CertFile: viper.GetString("file.cert"),
		KeyFile: viper.GetString("file.key"),
	}
	return &cfg, nil
}
