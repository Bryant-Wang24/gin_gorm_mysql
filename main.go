package main

import (
	"gin/database"
	"gin/routers"
)

func main() {
	database.InitMySQL()
	routers.SetupRouter()
}
