package initialize

import (
	"github.com/spf13/viper"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func InitConfig() {

}
