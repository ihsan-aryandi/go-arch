package envloader

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Load(env string) error {
	fileName := fmt.Sprintf(".env.%s", env)

	if err := godotenv.Load(fileName); err != nil {
		return godotenv.Load()
	}

	return nil
}
