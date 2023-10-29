package main

import (
	"github.com/fitriadwiseptiani/task-5-pbi-btpns-Fitria-Dwi-Septiani/controllers/usercontroller"
	"github.com/fitriadwiseptiani/task-5-pbi-btpns-Fitria-Dwi-Septiani/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/users/register", usercontroller.Index)
	r.GET("", usercontroller.Show)
	r.POST("", usercontroller.Create)
	r.PUT("", usercontroller.Update)
	r.DELETE("", usercontroller.Delete)

	r.Run()

}
