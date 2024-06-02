package config

type Config struct {
	Address                    string
	ToursServiceAdress         string
	EncounterServiceAdress     string
	StakeholdersServiceAddress string
	FollowerServiceAddress     string
}

func GetConfig() Config {
	return Config{
		ToursServiceAdress:         "localhost:8081",
		StakeholdersServiceAddress: "localhost:50051",
		EncounterServiceAdress:     "localhost:50052",
		FollowerServiceAddress:     "localhost:50053",
		Address:                    "localhost:8000",
	}
}
