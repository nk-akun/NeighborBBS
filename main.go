package main

import (
	"github.com/nk-akun/NeighborBBS/api"
	"github.com/nk-akun/NeighborBBS/config"
	"github.com/nk-akun/NeighborBBS/logs"
)

func init() {
	config.ParseConf()
	logs.InitLogger("./", 1, 1, 2, false)
	//logs.Logger.Info(config.GetConf().Viper.GetString("host"))
	// logs.Logger.Errorf("error!!!!!")

}

func main() {

	api.AppRun()
}
