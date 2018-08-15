package apis

import (
	"net/http"
	. "position_postgre/tools"

	"github.com/gin-gonic/gin"
)

func GetQiniuTokenApi(c *gin.Context) {
	token := GetQiniuToken()
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"token": token,
	})
}
