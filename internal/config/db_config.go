package config

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type DBConfig struct {
	DbUsername string `validate:"required"`
	DbPassword string `validate:"required"`
	DbNet      string `validate:"required"`
	DbAddress  string `validate:"required"`
	DbName     string `validate:"required"`
}

func LoadDbConfig() (*DBConfig, error) {
	// load the env file
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Something went wrong while loading the env vars: " + err.Error())
		return nil, err
	}

	// load the env vars and create the db config
	cfg := &DBConfig{
		DbUsername: LoadSingleEnvVar("DB_USERNAME", "root"),
		DbPassword: LoadSingleEnvVar("DB_PASSWORD", "root"),
		DbNet:      LoadSingleEnvVar("DB_NET", "tcp"),
		DbAddress:  LoadSingleEnvVar("DB_ADDRESS", "127.0.0.1:3306"),
		DbName:     LoadSingleEnvVar("DB_NAME", "auth_db"),
	}

	return cfg, nil
}

func SetupDB(dbConfig *DBConfig, logger *zap.Logger) (*sql.DB, error) {

	// create a mysql config
	cfg := mysql.NewConfig()
	cfg.User = dbConfig.DbUsername
	cfg.Passwd = dbConfig.DbPassword
	cfg.Net = dbConfig.DbNet
	cfg.Addr = dbConfig.DbAddress
	cfg.DBName = dbConfig.DbName

	// open a new db connection
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		logger.Fatal("Something went wrong while opening database connection",
			zap.String("error", err.Error()))

		return nil, err
	}

	// ping the database
	err = db.Ping()

	if err != nil {
		logger.Fatal("Something went wrong while pinging the database",
			zap.String("error", err.Error()))

		return nil, err
	}

	logger.Info("Successfully connected to the db",
		zap.String("db_name", dbConfig.DbName))

	return db, nil
}
