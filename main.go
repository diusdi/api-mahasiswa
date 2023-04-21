package main

import (
	"api-mahasiswa/config"
	"api-mahasiswa/routers"
	"log"

	_ "api-mahasiswa/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Mahasiswa
// @version 1.0
// @description API untuk mengatur data mahasiswa Jobhun. Untuk source code dapat dilihat di https://github.com/diusdi/api-mahasiswa

// @license.name Licensi MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	router := gin.Default()
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}
	mahasiswaRoute := routers.NewMahasiswaRoute(router.Group("/mhs"), db)
	mahasiswaRoute.SetupRoutes()

	// route dokumentasi
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
