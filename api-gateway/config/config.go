package config

type Config struct {
	Address                    string
	ToursServiceAdress         string
	EncounterServiceAdress     string
	StakeholdersServiceAddress string
}

func GetConfig() Config {
	return Config{
		ToursServiceAdress:         "localhost:8081",
		StakeholdersServiceAddress: "localhost:50051",
	}
}
