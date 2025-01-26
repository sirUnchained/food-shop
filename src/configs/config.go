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

var cfg Configs

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
	Db   int
}

type jwt struct {
	AccessSecret          string
	RefreshSecret         string
	AccessTokenExpiresIn  int
	RefreshTokenExpiresIn int
}

func GetConfigs() *Configs {
	return &cfg
}

func InitConfigs() {
	env := os.Getenv("APP_ENV")
	address := getFileAddress(&env)
	res := readConfigFileAndGet(&address)
	cfg = res
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
