package main

import (
	"github.com/nk-akun/NeighborBBS/config"
)

func initConf() {
	config.InitLogger("./", 1, 1, 2, false)
	log := config.GetLogger()
	log.Errorf("error!!!!!")
}

func main() {
	initConf()
	config.ParseConf()
	config.GetLogger().Info(config.GetConf().Viper.GetString("host"))
}
