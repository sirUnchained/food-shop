package postgres

import (
	"fmt"
	"foodshop/configs"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitPostgres(cfg *configs.Configs) error {
	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Postgres.Host, cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Dbname, cfg.Postgres.Port)

	dbClient, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		return err
	}

	// testing db
	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	connection, _ := dbClient.DB()
	connection.Close()
}
