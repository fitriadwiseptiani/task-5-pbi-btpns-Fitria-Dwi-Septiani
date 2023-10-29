package main

import (
	"github.com/fitriadwiseptiani/task-5-pbi-btpns-Fitria-Dwi-Septiani/controllers/usercontroller"
	"github.com/fitriadwiseptiani/task-5-pbi-btpns-Fitria-Dwi-Septiani/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/users", usercontroller.Index)
	r.GET("/users/:id", usercontroller.Show)
	r.POST("/users/register", usercontroller.Create)
	r.PUT("/users/:id", usercontroller.Update)
	r.DELETE("/users", usercontroller.Delete)

	r.Run()

}
