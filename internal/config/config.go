package config
import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	DSN string
}

var AppConfig *Config

func LoadConfig(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err:=viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
}