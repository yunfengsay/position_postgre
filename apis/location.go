package apis

import (
	"github.com/gin-gonic/gin"
	"position_postgre/modelStruct"
	"position_postgre/models"
	"position_postgre/tools"
)

func AddLocationApi(c *gin.Context) {
	locationForm := &modelStruct.AddLocationApiForm{}
	err := c.BindJSON(locationForm)
	tools.PanicError(err)
	models.AddLocation(locationForm)
}
