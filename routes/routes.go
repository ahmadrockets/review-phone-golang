package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"final-project-sanber/controllers"
	"final-project-sanber/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", controllers.WelcomeApi)

	// auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	// changepass
	changePassMiddlewareRoute := r.Group("/changepassword")
	changePassMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	changePassMiddlewareRoute.POST("/", controllers.ChangePassword)

	// roles
	r.GET("/roles", controllers.GetAllRole)
	roleMiddlewareRoute := r.Group("/roles")
	roleMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	roleMiddlewareRoute.POST("/", controllers.CreateRole)
	roleMiddlewareRoute.PATCH("/:id", controllers.UpdateRole)
	roleMiddlewareRoute.DELETE("/:id", controllers.DeleteRole)

	// brands
	r.GET("/brands", controllers.GetAllBrand)
	brandMiddlewareRoute := r.Group("/brands")
	brandMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	brandMiddlewareRoute.POST("/", controllers.CreateBrand)
	brandMiddlewareRoute.PATCH("/:id", controllers.UpdateBrand)
	brandMiddlewareRoute.DELETE("/:id", controllers.DeleteBrand)

	// phones
	r.GET("/phones", controllers.GetAllPhone)
	r.GET("/phones/bybrand/:id", controllers.GetAllPhoneByBrand)
	r.GET("/phones/search/:keyword", controllers.GetAllPhoneByKeyword)
	r.GET("/phones/:id", controllers.GetDetailPhone)
	r.GET("/phones/fav", controllers.GetFavouritePhone)
	phoneMiddlewareRoute := r.Group("/phones")
	phoneMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	phoneMiddlewareRoute.POST("/", controllers.CreatePhone)
	phoneMiddlewareRoute.PATCH("/:id", controllers.UpdatePhone)
	phoneMiddlewareRoute.DELETE("/:id", controllers.DeletePhone)

	// pictures
	r.GET("/pictures", controllers.GetAllPicture)
	r.GET("/pictures/byphone/:id", controllers.GetAllPictureByPhoneID)
	pictureMiddlewareRoute := r.Group("/pictures")
	pictureMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	pictureMiddlewareRoute.POST("/", controllers.CreatePicture)
	pictureMiddlewareRoute.PATCH("/:id", controllers.UpdatePicture)
	pictureMiddlewareRoute.DELETE("/:id", controllers.DeletePicture)

	// reviews
	r.GET("/reviews", controllers.GetAllReview)
	r.GET("/reviews/byphone/:id", controllers.GetAllReviewByPhoneID)
	reviewMiddlewareRoute := r.Group("/reviews")
	reviewMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	reviewMiddlewareRoute.POST("/", controllers.CreateReview)
	reviewMiddlewareRoute.DELETE("/:id", controllers.DeleteReview)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
