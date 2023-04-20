package main

import (
	"api-mahasiswa/config"
	"api-mahasiswa/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	mahasiswaRoute := routers.NewMahasiswaRoute(router.Group("/mhs"), db)
	mahasiswaRoute.SetupRoutes()

	router.Run(":8080")
}
