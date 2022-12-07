package main

import (
	"TodoList/conf"
	"TodoList/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
