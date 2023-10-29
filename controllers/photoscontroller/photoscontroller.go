package photoscontroller

import (
	"encoding/json"
	"net/http"

	"github.com/fitriadwiseptiani/task-5-pbi-btpns-Fitria-Dwi-Septiani/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var photoss []models.Photos

	models.DB.Find(&photoss)
	c.JSON(http.StatusOK, gin.H{"photoss": photoss})
}
func Show(c *gin.Context) {
	var photos models.Photos
	id := c.Param("id")

	if err := models.DB.First(&photos, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"photos": photos})
}
func Create(c *gin.Context) {
	var photos models.Photos
	if err := c.ShouldBindJSON(&photos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&photos)
	c.JSON(http.StatusOK, gin.H{"photos": photos})
}
func Update(c *gin.Context) {
	var photos models.Photos
	id := c.Param("id")

	if err := c.ShouldBindJSON(&photos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&photos).Where("id = ?", id).Updates(&photos).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data photos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}
func Delete(c *gin.Context) {
	var photos models.Photos

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&photos, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data photos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
