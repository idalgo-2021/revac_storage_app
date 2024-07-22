package config

import (
	"encoding/json"
	"os"
)

type ResumeServiceConfig struct {
	MaxResumesPerUser               int  `json:"maxResumesPerUser"`               // Значение(максимальное) того, сколько резюме разрешено иметь соискателю
	ControlQntResumesPerUserEnabled bool `json:"controlQntResumesPerUserEnabled"` // Флаг использования контроля параметра MaxResumesPerUser
}

type DBConnectionConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
}

type HTTPServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	DBConnectionConfig  DBConnectionConfig  `json:"dbConnectionConfig"` // Адрес БД, с которой работает сервис
	HTTPServerConfig    HTTPServerConfig    `json:"HTTPServerConfig"`   // Параметры запускаемого HTTP-сервера
	ResumeServiceConfig ResumeServiceConfig `json:"ResumeServiceConfig"`
}

func LoadConfig(c *Config, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(c); err != nil {
		return err
	}

	return nil
}
