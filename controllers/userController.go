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
)

const (
	MinCost     int = 4
	MaxCost     int = 40
	DefaultCost int = 10
)

var (
	// validate *validator.Validate
	appJSON = "application/json"
)

func (idb *InDB) UserRegister(c *gin.Context) {

	db := database.GetDB()
	contenType := helper.GetContentType(c)
	_, _ = db, contenType
	User := models.User{}

	if contenType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.Created_At = time.Now()
	User.Updated_At = time.Now()

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"password": User.Password,
		"username": User.Username,
	})

}

func (idb *InDB) UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBind(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (idb *InDB) UserUpdate(c *gin.Context) {

	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	User := models.User{}

	userId, _ := strconv.Atoi(c.Param("userId"))

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User.Updated_At = time.Now()
	User.ID = uint(userID)
	err := db.Model(&User).Where("id = ?", userId).Updates(&User).First(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.Updated_At,
	})
}

func (idb *InDB) UserDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	err := db.Debug().Where("id = ?", userID).Delete(&User).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo Successfully Deleted",
	})
}
