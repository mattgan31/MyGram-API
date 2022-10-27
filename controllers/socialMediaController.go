package controllers

import (
	"final-project-fga/database"
	"final-project-fga/helper"
	"final-project-fga/models"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func (idb *InDB) CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.User_ID = userID
	SocialMedia.Updated_At = time.Now()
	SocialMedia.Created_At = time.Now()

	if SocialMedia.User_ID == userID {
		err := db.Debug().Create(&SocialMedia).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.Social_Media_URL,
		"user_id":          SocialMedia.User_ID,
		"created_at":       SocialMedia.Created_At,
	})
}

func (idb *InDB) GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := []models.SocialMedia{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := uint(userData["id"].(float64))

	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username")
	}).Where("user_id = ?", userID).Find(&SocialMedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func (idb *InDB) UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}
	contentType := helper.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	socialMediaID, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert",
		})
	}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.User_ID = userID
	SocialMedia.Updated_At = time.Now()
	SocialMedia.ID = uint(socialMediaID)

	err = db.Debug().Where("id=?", socialMediaID).Updates(models.SocialMedia{Name: SocialMedia.Name, Social_Media_URL: SocialMedia.Social_Media_URL}).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.Social_Media_URL,
		"user_id":          SocialMedia.User_ID,
		"updated_at":       SocialMedia.Updated_At,
	})
}

func (idb *InDB) DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	socialMediaID, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert",
		})
	}

	userID := uint(userData["id"].(float64))
	SocialMedia.User_ID = userID
	SocialMedia.ID = uint(socialMediaID)

	err = db.Debug().Where("id = ?", socialMediaID).Delete(&SocialMedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
