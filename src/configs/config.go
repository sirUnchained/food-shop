package configs

import (
	"encoding/json"
	"io"
	"os"
)

type Configs struct {
	Server   server
	Postgres postgres
	Redis    redis
	Jwt      jwt
}

type server struct {
	Port string
	Host string
}

type postgres struct {
	Port     string
	Host     string
	Username string
	Password string
	Dbname   string
}

type redis struct {
	Port string
	Host string
}

type jwt struct {
	SecretKey string
	ExpiresIn string
}

func GetConfigs() Configs {
	env := os.Getenv("APP_ENV")
	address := getFileAddress(&env)
	cfg := readConfigFileAndGet(&address)
	return cfg
}

func getFileAddress(env *string) string {
	if *env == "production" {
		return "configs/config.prod.json"

	}
	return "configs/config.dev.json"
}

func readConfigFileAndGet(path *string) Configs {
	var cfg Configs
	jsonFile, err := os.Open(*path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &cfg)
	return cfg
}
