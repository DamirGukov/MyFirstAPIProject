package main

import (
	"FirstAPIProject/docs"
	"FirstAPIProject/server"
	"github.com/sirupsen/logrus"
)

// @title           MyBestAPIProject
// @version         1.0
// @description     This my first Project
// @termsOfService  http://swagger.io/terms/
// @BasePath /
func main() {
	docs.SwaggerInfo.Host = ""

	err := server.Start()
	if err != nil {
		logrus.Fatal("failed to start server", err)
	}
}
