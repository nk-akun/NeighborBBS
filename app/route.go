package app

import (
	"github.com/nk-akun/NeighborBBS/config"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func Test() {
	log = config.GetLogger()
	log.Info("this's test")
}
