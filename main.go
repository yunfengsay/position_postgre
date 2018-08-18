package main

import (
	"log"
	"position_postgre/conf"
	"position_postgre/db"

	"position_postgre/apis"

	"github.com/gin-gonic/gin"
)

func AuthNeedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		log.Println("token ", token)
		if token == "" {
			c.AbortWithStatus(400)
		}
		// id := models.GetUserIdByToken(token)
		// c.Set("userid", id)
		c.Next()
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", apis.IndexApi)
	user := router.Group("/user")
	user.POST("/signup", apis.AddUserApi)
	user.POST("/wx/login", apis.WXLogin)
	location := router.Group("/location")

	location.POST("/add_location", AuthNeedLogin(), apis.AddLocationApi)
	// location.POST("/get_locations", apis.GetLocationsApi)
	// location.GET("/get_location/:id", apis.GetPageByIdApi)
	// location.DELETE("/delete_location/:id", apis.DeleteLocation)
	// location.POST("/comment", AuthNeedLogin(), apis.AddComment)

	router.GET("/get_upload_token", AuthNeedLogin(), apis.GetQiniuTokenApi)
	// router.POST("/like/update", AuthNeedLogin(), apis.UpdateLike)
	// router.DELETE("/user/delete", AuthNeedLogin(), apis.DeleteUserApi)

	return router
}

func main() {
	defer db.DB.Close()

	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println(db.DB)

	gin.SetMode(gin.ReleaseMode)
	// MongoSession := db.MongoSession
	// defer MongoSession.Close()
	router := initRouter()
	log.Println("😄")
	err := router.Run(conf.ConfigContext.ServerPort)
	log.Println("err ", err)
}
