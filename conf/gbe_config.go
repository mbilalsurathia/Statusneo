package conf

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var (
	configPath     = "."
	configFileName = "conf.json"
)

type GbeConfig struct {
	DataSource DataSourceConfig `json:"dataSource"`
	Rest       RestServer       `json:"rest"`
	Logger     Logger           `json:"logger"`
	Env        string           `json:"env"`
	Email      Email            `json:"emailConfig"`
}

type Logger struct {
	Path  string `yaml:"path" json:"path"`
	Level string `yaml:"level" json:"level"`
	Name  string `yaml:"name" json:"name"`
	Mode  string `yaml:"mode" json:"mode"`
}

type DataSourceConfig struct {
	Host              string `json:"host"`
	Port              string `json:"port"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	SslMode           string `json:"sslMode"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
	Retries           int    `json:"retries"`
	Mode              int    `json:"mode"`
}

type Email struct {
	IsEnabled bool   `json:"isEnabled"`
	SmtpHost  string `json:"host"`
	SmtpPort  string `json:"port"`
	From      string `json:"from"`
	Password  string `json:"password"`
}

type RestServer struct {
	Addr string `json:"addr"`
}

var config GbeConfig
var configOnce sync.Once

func GetConfig() *GbeConfig {
	configOnce.Do(func() {

		bytes, err := os.ReadFile(configPath + "/" + configFileName)
		log.Println(configPath + "/" + configFileName)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})

	return &config
}
