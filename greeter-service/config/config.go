package config



type Config struct {
	Address string
}

func GetConfig() Config {
	return Config{
		Address: "localhost:8085",
	}
}
