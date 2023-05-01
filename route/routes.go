package route

import (
	controllers "web-api/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes(router *gin.Engine) {

	/***BASEPATH OF AN API. NOTE:THIS SHOULDN'T BE CHANGED***/
	api := router.Group("/MocktestAPI")

	/***ADD THE ROUTES HERE***/
	api.POST("/saveUsers", controllers.CreateUsers)
	api.GET("/readUsers", controllers.ReadUsers)
	api.PUT("/updateUsers", controllers.UpdateUsers)
	api.DELETE("/deleteUsers", controllers.DeleteUsers)
	api.POST("/savePosts", controllers.CreatePost)
	api.GET("/readPosts", controllers.ReadPost)
	api.GET("/findPosts", controllers.FindPostsByUser)

	router.Run(viper.GetString("server.port"))
}
