package usercontroller

import (
	"net/http"

	"github.com/fitriadwiseptiani/task-5-pbi-btpns-Fitria-Dwi-Septiani/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func Show(c *gin.Context) {

}
func Create(c *gin.Context) {

}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
