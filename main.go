package main

<<<<<<< HEAD
func main() {

=======
import (
	"TodoList/conf"
	"TodoList/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
>>>>>>> 786ab24 (go web项目基于gin+gorm开发的备忘录)
}
