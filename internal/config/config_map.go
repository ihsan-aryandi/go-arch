package config

type Conf struct {
	App  *AppConfig
	HTTP *HTTPConfig
}

type AppConfig struct {
	Version string
}

type HTTPConfig struct {
	Port string
}
