package cfg

import (
	"NPP/logUtils"
	"github.com/spf13/viper"
)

var (
	V *viper.Viper
)

func InitViper() {
	V = viper.New()
	V.SetConfigType("yaml")
	V.SetConfigFile("")
	V.AddConfigPath("./")
	V.SetConfigName("cfg.yaml")
	if err := V.ReadInConfig(); err != nil {
		logUtils.Println("err: " + err.Error())
	}
}
