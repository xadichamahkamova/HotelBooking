package load

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type KafkaConfig struct {
	Host  string
	Port  string
	Topic string
}

type Config struct {
	ApiGateway          ServiceConfig
	UserService         ServiceConfig
	HotelService        ServiceConfig
	BookingService      ServiceConfig
	NotificationService ServiceConfig
	TokenKey            string
	Kafka               KafkaConfig
	CertFile            string
	KeyFile             string
}

func Load(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		ApiGateway: ServiceConfig{
			Host: viper.GetString("api_gateway.host"),
			Port: viper.GetInt("api_gateway.port"),
		},
		UserService: ServiceConfig{
			Host: viper.GetString("services.user_service.host"),
			Port: viper.GetInt("services.user_service.port"),
		},
		HotelService: ServiceConfig{
			Host: viper.GetString("services.hotel_service.host"),
			Port: viper.GetInt("services.hotel_service.port"),
		},
		BookingService: ServiceConfig{
			Host: viper.GetString("services.booking_service.host"),
			Port: viper.GetInt("services.booking_service.port"),
		},
		NotificationService: ServiceConfig{
			Host: viper.GetString("services.notification_service.host"),
			Port: viper.GetInt("services.notification_service.port"),
		},

		TokenKey: viper.GetString("token.key"),

		Kafka: KafkaConfig{
			Host:  viper.GetString("kafka.host"),
			Port:  viper.GetString("kafka.port"),
			Topic: viper.GetString("kafka.topic"),
		},
		
		CertFile: viper.GetString("file.cert"),
		KeyFile: viper.GetString("file.key"),
	}
	return &cfg, nil
}
