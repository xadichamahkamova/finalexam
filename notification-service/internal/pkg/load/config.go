package load

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port string
}

type EmailConfig struct {
	From     string
	Password string
}

type KafkaConfig struct {
	Host  string
	Port  string
	Topic string
}

type Config struct {
	Service  ServiceConfig
	Kafka    KafkaConfig
	Email    EmailConfig
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

		Service: ServiceConfig{
			Host: viper.GetString("service.host"),
			Port: viper.GetString("service.port"),
		},

		Kafka: KafkaConfig{
			Host:  viper.GetString("kafka.host"),
			Port:  viper.GetString("kafka.port"),
			Topic: viper.GetString("kafka.topic"),
		},

		Email: EmailConfig{
			From:     viper.GetString("email.from"),
			Password: viper.GetString("email.password"),
		},

		CertFile: viper.GetString("file.cert"),
		KeyFile:  viper.GetString("file.key"),
	}
	return &cfg, nil
}
