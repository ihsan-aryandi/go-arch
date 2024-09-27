package envloader

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv(env string) error {
	if err := godotenv.Load(fmt.Sprintf(".env.%s", env)); err != nil {
		return godotenv.Load()
	}

	return nil
}
