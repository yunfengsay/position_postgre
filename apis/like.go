package apis

import (
	"log"
	"net/http"
	"position_postgre/tools"

	"github.com/gin-gonic/gin"
)

type UpdateLikePostBody struct {
	Target     string
	Type       string
	TargetUser string `json:"target_user"`
}

func UpdateLike(c *gin.Context) {
	data := &UpdateLikePostBody{}
	e := c.BindJSON(&data)
	log.Println("liked data ", data)
	tools.PanicError(e)
	// user, _ := c.Get("userid")
	// models.AddOrDeleteLike(data.Target, data.TargetUser, user.(string), data.Type)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
