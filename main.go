package main

import (
	"fmt"

	"github.com/JerryJeager/mingle-backend/cmd"
	"github.com/JerryJeager/mingle-backend/config"
)

func init() {
	config.ConnectToDB()
	config.LoadEnv()
}

func main() {
	fmt.Println("starting the Mingle server...")
	cmd.ExecuteApiRoutes()
}