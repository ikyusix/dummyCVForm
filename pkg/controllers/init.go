package controllers

import (
	"dummyCVForm/models"
	"dummyCVForm/pkg/logger"
	"dummyCVForm/pkg/middleware"
	"dummyCVForm/pkg/postgres"
	"dummyCVForm/utils/config"
	"dummyCVForm/utils/constants"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	logger.Log.Printf("Starting %s on DEVELOPMENT Environment", config.MyConfig.AppName)
	middleware.SetupLogger()
	err := postgres.InitDBConnection()
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
}

func HandleNoRoutes(c *gin.Context) {
	c.JSON(http.StatusNotFound, models.CreateResponse(c, fmt.Sprintf("%v", http.StatusNotFound), constants.UndefinedProcess, constants.WarnUndefinedProcess, nil))
}

func HandleNoMethod(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, models.CreateResponse(c, fmt.Sprintf("%v", http.StatusMethodNotAllowed), constants.UndefinedMethods, constants.WarnUndefinedMethod, nil))
}

func HandlePanic(c *gin.Context, err interface{}) {
	logger.Log.Error(err.(error).Error())
	c.JSON(http.StatusInternalServerError, models.CreateResponse(c, fmt.Sprintf("%v", http.StatusInternalServerError), constants.FAILED, constants.WarnInternalError, nil))
}
