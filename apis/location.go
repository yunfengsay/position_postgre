package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"position_postgre/modelStruct"
	"position_postgre/models"
	"position_postgre/tools"
)

func AddLocationApi(c *gin.Context) {
	locationForm := &modelStruct.AddLocationApiForm{}
	err := c.BindJSON(locationForm)
	tools.PanicError(err)
	session, _ := c.Get("session")

	user := models.FindUserBySession(session.(string))
	models.AddLocation(locationForm, user)
}

func GetLocationsApi(c *gin.Context) {

}

func NeerLocation(c *gin.Context) {
	location_request := &modelStruct.LocationRequest{}
	err := c.BindJSON(location_request)
	tools.PanicError(err)
	locations := models.NeerLocation(location_request)
	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"locations": locations,
	})
}
