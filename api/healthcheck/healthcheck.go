package healthcheck

import (
	"dummyCVForm/models"
	"dummyCVForm/pkg/logger"
	"dummyCVForm/utils/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleHealthCheck(c *gin.Context) {
	logger.Log.Infoln("Health check success")
	c.JSON(http.StatusOK, models.CreateResponse(c, constants.SuccessCode, constants.HealthCheck, constants.WarnHealthSuccess, nil))
}
