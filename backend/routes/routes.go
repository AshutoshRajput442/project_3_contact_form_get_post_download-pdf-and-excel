package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/submit", handlers.SubmitForm)
	router.GET("/data", handlers.GetData)
	router.GET("/download/csv", handlers.DownloadCSV)
	router.GET("/download/pdf", handlers.DownloadPDF)
}
