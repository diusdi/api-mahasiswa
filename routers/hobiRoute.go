package routers

import (
	"api-mahasiswa/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type HobiRoute struct {
	Router              *gin.RouterGroup
	HobiController *controllers.HobiController
}

func (tr *HobiRoute) SetupRoutes() {
	tr.Router.POST("/", tr.HobiController.Create)
	tr.Router.GET("/", tr.HobiController.Read)
	tr.Router.PUT("/:id", tr.HobiController.Update)
}

func NewHobiRoute(router *gin.RouterGroup, db *sql.DB) *HobiRoute {
	return &HobiRoute{
		Router:              router,
		HobiController: &controllers.HobiController{DB: db},
	}
}
