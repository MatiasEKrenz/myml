package myml

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myml/src/api/services/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"strconv"
)

/*func User (context *gin.Context) {

	//context.String(200, "pong")
	userID := context.Param("id")
	user := myml.GetUsersParams(userID)
	stringUser := user.Nickname
	context.String(200, stringUser)

}*/

const (
	paramUserID = "id"
)

// el controlador no puede devolver nada, ni errores
func GetUser(c *gin.Context) {

	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	user, apiErr := myml.GetUserFromAPI(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetInfo(c *gin.Context) {

	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	general, apiErr := myml.GetGeneralInfo(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}

	c.JSON(http.StatusOK, general)
}
