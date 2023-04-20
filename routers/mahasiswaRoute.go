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
	tr.Router.PUT("/:id", tr.MahasiswaController.Update)
	tr.Router.DELETE("/:id", tr.MahasiswaController.Delete)
}

func NewMahasiswaRoute(router *gin.RouterGroup, db *sql.DB) *MahasiswaRoute {
	return &MahasiswaRoute{
		Router:              router,
		MahasiswaController: &controllers.MahasiswaController{DB: db},
	}
}
