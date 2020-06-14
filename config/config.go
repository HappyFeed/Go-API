package config

import (
	"fmt"
	//"github.com/caarlos0/env"
	"github.com/eduardogpg/gonv"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

var database *DatabaseConfig

func init() {
	database = &DatabaseConfig{}
	database.Username = gonv.GetStringEnv("USERNAME", "root")
	database.Password = gonv.GetStringEnv("PASSWORD", "usuario")
	database.Host = gonv.GetStringEnv("HOST", "localhost")
	database.Port = gonv.GetIntEnv("PORT", 3306)
	database.Database = gonv.GetStringEnv("DATABASE", "project_go_web")
}

func (this *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", this.Username, this.Password, this.Host, this.Port, this.Database)
}

func GetUrlDatabase() string {
	return database.url()
}
