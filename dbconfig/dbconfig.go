package dbconfig

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	Driver   string
}

func NewDBConfig(d string) *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Driver:   d,
	}
}

func (c *DBConfig) GetConnString() (string, string) {
	return c.Driver, fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.Host, c.Port, c.Database, c.User, c.Password)
}
