package config

import "os"

func LoadConfig() *Conf {
	return &Conf{
		App: &AppConfig{
			Version: getEnv("APP_VERSION", "0.0.1"),
		},
		HTTP: &HTTPConfig{
			Port: getEnv("HTTP_PORT", "8000"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
