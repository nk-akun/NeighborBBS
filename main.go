package main

import (
	"fmt"

	"github.com/nk-akun/NeighborBBS/api"
	"github.com/nk-akun/NeighborBBS/config"
	"github.com/nk-akun/NeighborBBS/logs"
	"github.com/nk-akun/NeighborBBS/model"
	"github.com/nk-akun/NeighborBBS/util"
)

func init() {
	config.ParseConf()
	logs.InitLogger("./", 1, 1, 2, false, true) //TODO: 将log配置加在配置文件中
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci", config.GetConf().MySQL.Username,
		config.GetConf().MySQL.Password, config.GetConf().MySQL.Host, config.GetConf().MySQL.Port, config.GetConf().MySQL.DbName)
	util.OpenDB(dsn, nil, 10, 20, model.Models...)

	// util.OpenRedis(config.GetConf().Redis.Host, config.GetConf().Redis.Port, config.GetConf().Redis.Password)
}

func main() {

	// user := &model.User{
	// 	Username: "marthon",
	// 	Nickname: "用户001",
	// 	Password: "654321",
	// 	Gender:   "0",
	// }

	// util.DB().Create(user)
	defer util.CloseDB()

	api.AppRun()
}
