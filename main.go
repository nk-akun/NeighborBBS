package main

import (
	"github.com/nk-akun/NeighborBBS/api"
	"github.com/nk-akun/NeighborBBS/config"
	"github.com/nk-akun/NeighborBBS/logs"
)

func initConf() {
	logs.InitLogger("./", 1, 1, 2, false)
	logs.Logger.Errorf("error!!!!!")
}

func main() {
	initConf()
	config.ParseConf()
	logs.Logger.Info(config.GetConf().Viper.GetString("host"))
	api.AppRun()
}
