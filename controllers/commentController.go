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

func (idb *InDB) CreateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photo := models.Photo{}
	Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	photoID := uint(Comment.Photo_ID)

	err := db.Debug().First(&Photo, photoID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "not found",
			"message": "photo with that id not found",
		})
		return
	}

	Comment.Created_At = time.Now()
	Comment.Updated_At = time.Now()
	Comment.User_ID = userID

	err = db.Debug().Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	// }

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.Photo_ID,
		"user_id":    Comment.User_ID,
		"created_at": Comment.Created_At,
	})
}

func (idb *InDB) GetComment(c *gin.Context) {

	db := database.GetDB()
	Comment := []models.Comment{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := uint(userData["id"].(float64))

	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Title", "Caption", "Photo_URL", "User_ID")
	}).Where("user_id = ?", userID).Find(&Comment).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)

}

func (idb *InDB) UpdateComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert",
		})
	}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.Updated_At = time.Now()
	Comment.User_ID = userID
	Comment.ID = uint(commentID)

	err = db.Debug().Where("id=?", commentID).Updates(models.Comment{Message: Comment.Message}).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      Comment.ID,
		"message": Comment.Message,
		"user_id": Comment.User_ID,
	})

}

func (idb *InDB) DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert",
		})
	}

	userID := uint(userData["id"].(float64))

	Comment.User_ID = userID
	Comment.ID = uint(commentID)

	err = db.Debug().Where("id = ?", commentID).Delete(&Comment).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
