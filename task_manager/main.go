package main

import (
	"taskManager/cmd"
	"taskManager/database"
)

func main() {
	database.Mysql()
	defer database.DB.Close()
	cmd.Serve()
}