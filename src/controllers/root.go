package src_controllers

import (
	"net/http"
	"runtime"
	src_services "vazir_hessab/src/services"

	"github.com/gin-gonic/gin"
)

func Get_Version(ctx *gin.Context) {
	server_version := src_services.Get_Version()
	go_version := runtime.Version()[2:]

	json := []string{
		server_version, go_version,
	}

	ctx.JSON(http.StatusOK, json)
}
