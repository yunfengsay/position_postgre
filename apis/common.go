package apis

import (
	"github.com/gin-gonic/gin"
	//"position_postgre/db"
	//"time"
	"net/http"
)

type AddCommentType struct {
	LocationId int32 `json: "location_id"` // location的id
	To         int32 `json: "to"`          // location 的 user id
	From       int32 `json: "from"`        // 评论者的id
	UserAvater string        `json: "user_avater"` // 评论者的头像
	UserName   string        `json: "user_name"`   // 评论者的昵称
}

func IndexApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": 1})
}

func AddComment(c *gin.Context) {
	// comment := &models.Comments{}
	//postData := &AddCommentType{}
}
