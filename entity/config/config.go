package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type config struct {
	LoggerPath string
}

var configInstance *config

var configOnce sync.Once

func getConfigInstance() {
	configOnce.Do(func() {
		configInstance = &config{}

		configInstance.initLogger()
	})
}

// viper 配置初始化
func viperInit() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../etc/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("viper read error: %v", err))
	}
}

func (c *config) initLogger() {
	c.LoggerPath = viper.GetString("")

}
