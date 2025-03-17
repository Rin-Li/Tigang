package main

import (
	"Tigang/conf"
	"Tigang/repository/cache"
	"Tigang/routes"
)

func main(){
	conf.Init()
	cache.InitRedis()


	r := routes.NewRouter()
	_ = r.Run(":3000")
}