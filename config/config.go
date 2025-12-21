package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"PG_DB_HOST"`
	DBPort     string `mapstructure:"PG_DB_PORT"`
	DBUser     string `mapstructure:"PG_DB_USER"`
	DBPassword string `mapstructure:"PG_DB_PASSWORD"`
	DBName     string `mapstructure:"PG_AYEE_PORTAL_DB_NAME"`
	ServerPort string `mapstructure:"AYEE_PORTAL_SERVER_PORT"`

	KeycloakIssuer   string `mapstructure:"KEYCLOAK_ISSUER"`
	KeycloakClientID string `mapstructure:"KEYCLOAK_CLIENT_ID"`
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("PG_DB_HOST", "localhost")
	viper.SetDefault("PG_DB_PORT", "5432")
	viper.SetDefault("PG_DB_USER", "postgres")
	viper.SetDefault("PG_DB_PASSWORD", "password")
	viper.SetDefault("PG_AYEE_PORTAL_DB_NAME", "ayeeportal")
	viper.SetDefault("AYEE_PORTAL_SERVER_PORT", "8080")

	// Keycloak defaults (optional, but good to have a placeholder or expect env var)
	viper.SetDefault("KEYCLOAK_ISSUER", "")
	viper.SetDefault("KEYCLOAK_CLIENT_ID", "")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using defaults or env vars")
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal("Unable to decode into struct, ", err)
	}
}
