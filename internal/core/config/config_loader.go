package config

const defaultEnv = "development"

func LoadConfig(env string) {
	if env == "" {
		env = defaultEnv
	}

	//fmt.Sprintf(".env.%s", env)
	//godotenv.Load()
}
