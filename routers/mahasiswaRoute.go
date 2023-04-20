package routers

import (
	"api-mahasiswa/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type MahasiswaRoute struct {
	Router              *gin.RouterGroup
	MahasiswaController *controllers.MahasiswaController
}

func (tr *MahasiswaRoute) SetupRoutes() {
	tr.Router.POST("/", tr.MahasiswaController.Create)
	tr.Router.GET("/", tr.MahasiswaController.Read)
	tr.Router.GET("/:id", tr.MahasiswaController.ReadById)
}

func NewMahasiswaRoute(router *gin.RouterGroup, db *sql.DB) *MahasiswaRoute {
	return &MahasiswaRoute{
		Router:              router,
		MahasiswaController: &controllers.MahasiswaController{DB: db},
	}
}
