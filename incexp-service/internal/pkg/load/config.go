package load

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type KafkaConfig struct {
	Host  string
	Port  string
	Topic string
}

type Config struct {
	Postgres PostgresConfig
	Service  ServiceConfig
	Kafka    KafkaConfig
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

		Postgres: PostgresConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetString("postgres.port"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
			Database: viper.GetString("postgres.database"),
		},
		
		Kafka: KafkaConfig{
			Host:  viper.GetString("kafka.host"),
			Port:  viper.GetString("kafka.port"),
			Topic: viper.GetString("kafka.topic"),
		},
		
	}
	return &cfg, nil
}
