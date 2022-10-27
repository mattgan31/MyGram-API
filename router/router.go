package router

import (
	"final-project-fga/controllers"
	"final-project-fga/database"
	"final-project-fga/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	db := database.GetDB()
	router := gin.Default()
	ctr := controllers.New(db)

	userRouter := router.Group("/users")
	{

		userRouter.POST("/register", ctr.UserRegister)
		userRouter.POST("/login", ctr.UserLogin)
		userRouter.PUT("/:userId", middleware.Authentication(), middleware.UserAuthorization(), ctr.UserUpdate)
		userRouter.DELETE("/:userId", middleware.Authentication(), middleware.UserAuthorization(), ctr.UserDelete)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", ctr.CreatePhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), ctr.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), ctr.DeletePhoto)
		photoRouter.GET("/", ctr.GetPhoto)
	}
	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", ctr.CreateComment)
		commentRouter.GET("/", ctr.GetComment)

		commentRouter.PUT("/:commentId", middleware.Authentication(), middleware.CommentAuthorization(), ctr.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.Authentication(), middleware.CommentAuthorization(), ctr.DeleteComment)
	}
	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication(), middleware.SocialMediaAuthorization())
		socialMediaRouter.POST("/", ctr.CreateSocialMedia)
		socialMediaRouter.GET("/", ctr.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", ctr.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", ctr.DeleteSocialMedia)
	}

	return router
}
