package config

import (
	"fmt"
	//"github.com/caarlos0/env"
	"github.com/eduardogpg/gonv"
)

type DatabaseConfig struct {
	username string
	password string
	host     string
	port     int
	database string
	debug    bool
}

type Config interface {
	url() string
}

type ServerConfig struct {
	host  string
	port  int
	debug bool
}

var database *DatabaseConfig
var server *ServerConfig

func init() {
	database = &DatabaseConfig{}
	database.username = gonv.GetStringEnv("USERNAME", "root")
	database.password = gonv.GetStringEnv("PASSWORD", "usuario")
	database.host = gonv.GetStringEnv("HOST", "localhost")
	database.port = gonv.GetIntEnv("PORT", 3306)
	database.database = gonv.GetStringEnv("DATABASE", "project_go_web")
	database.debug = gonv.GetBoolEnv("DEBUG", true)

	server = &ServerConfig{}
	server.host = gonv.GetStringEnv("HOST", "localhost")
	server.port = gonv.GetIntEnv("PORT", 8000)
	server.debug = gonv.GetBoolEnv("DEBUG", true)
}

func Debug() bool {
	return server.debug
}

func (this *DatabaseConfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", this.username, this.password, this.host, this.port, this.database)
}

func UrlDatabase() string {
	return database.url()
}

func UrlServer() string {
	return server.url()
}

func (this *ServerConfig) url() string {
	return fmt.Sprintf("%s:%d", this.host, this.port)
}

func ServerPort() int {
	return server.port
}

func DirTemplate() string {
	return "templates/**/*.html"
}

func DirTemplateError() string {
	return "templates/error.html"
}
