package main

import (
	"fmt"

	"github.com/JerryJeager/mingle-backend/cmd"
	"github.com/JerryJeager/mingle-backend/config"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	fmt.Println("starting the Mingle server...")
	cmd.ExecuteApiRoutes()
}