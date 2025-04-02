package configs

import (
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type ProjectConfig struct {
	DBConfig *DBConfig
	UIUrl    string
	Port     string
}

var (
	config     *ProjectConfig
	configOnce sync.Once
)

func LoadProjectConfig() *ProjectConfig {
	_ = godotenv.Load()

	configOnce.Do(func() {
		config = &ProjectConfig{
			DBConfig: NewDBConfig(),
			UIUrl:    os.Getenv("UI_URL"),
			Port:     os.Getenv("PORT"),
		}
	})

	return config
}
