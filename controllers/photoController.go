package controllers

import (
	"final-project-fga/database"
	"final-project-fga/helper"
	"final-project-fga/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func (idb *InDB) CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	Photo.Created_At = time.Now()
	Photo.Updated_At = time.Now()
	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.User_ID = userID

	err := db.Debug().Create(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_URL,
		"user_id":    Photo.User_ID,
		"created_at": Photo.Created_At,
	})

}

func (idb *InDB) UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.User_ID = userID
	Photo.ID = uint(photoId)

	Photo.Updated_At = time.Now()
	err := db.Model(&Photo).Where("id = ?", photoId).Updates(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_URL,
		"user_id":    Photo.User_ID,
		"updated_at": Photo.Updated_At,
	})
}

func (idb *InDB) DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Model(&Photo).Where("id = ?", photoId).Where("user_id = ?", userID).Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo Successfully Deleted",
	})
}

func (idb *InDB) GetPhoto(c *gin.Context) {
	db := database.GetDB()
	Photos := []models.Photo{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := uint(userData["id"].(float64))

	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username")
	}).Where("user_id = ?", userID).Find(&Photos).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photos)
}
