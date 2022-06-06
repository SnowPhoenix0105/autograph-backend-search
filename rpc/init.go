package rpc

type Config struct {
	ControllerBaseUrl string
}

func GenerateTestConfig() *Config {
	return &Config{
		ControllerBaseUrl: "http://localhost:8003",
	}
}

var globalConfig Config

func Init(config *Config) {
	globalConfig = *config
}
