package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port               string `mapstructure:"PORT"`
	DbUrl              string `mapstructure:"DB_URL"`
	ClientUrl          string `mapstructure:"CLIENT_URL"`
	AwsAccessKeyId     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AwsRegion          string `mapstructure:"AWS_REGION"`
	AwsBucketName      string `mapstructure:"AWS_BUCKET_NAME"`
}

func LoadConfig() (AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	config := AppConfig{
		Port:               viper.GetString("PORT"),
		DbUrl:              viper.GetString("DB_URL"),
		ClientUrl:          viper.GetString("CLIENT_URL"),
		AwsAccessKeyId:     viper.GetString("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
		AwsRegion:          viper.GetString("AWS_REGION"),
		AwsBucketName:      viper.GetString("AWS_BUCKET_NAME"),
	}

	return config, nil
}
