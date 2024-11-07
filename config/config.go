package config

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	DB         *sql.DB
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	config := &Config{
		DBHost:     viper.GetString("database.DB_HOST"),
		DBPort:     viper.GetString("database.DB_PORT"),
		DBUser:     viper.GetString("database.DB_USER"),
		DBPassword: viper.GetString("database.DB_PASSWORD"),
		DBName:     viper.GetString("database.DB_NAME"),
		ServerPort: viper.GetString("server.PORT"),
	}

	return config, nil
}

func (c *Config) ConnectDatabase() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	c.DB = db
	return nil
}
