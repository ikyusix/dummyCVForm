package handler

import (
	"dummyCVForm/models"
	"dummyCVForm/pkg/logger"
	"dummyCVForm/utils/constants"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controllers struct {
	PUsecase models.ProfileUsecase
}

func NewControllersProfile(g *gin.RouterGroup, PUsecase models.ProfileUsecase) {
	handler := &Controllers{PUsecase: PUsecase}
	g.GET("/profile/:id", handler.GetProfile)
}

func (r *Controllers) GetProfile(c *gin.Context) {
	id := c.Param("id")
	data, err := r.PUsecase.Get(c, id)
	if err != nil {
		logger.Log.Errorf("[PROFILE][GET] ERROR %v for requestId: %v", err.Error(), requestid.Get(c))
		c.JSON(http.StatusInternalServerError, models.CreateResponse(c, constants.InternalServerCode, "[GetProfile] ERROR", constants.WarnInternalError, nil))
		return
	}

	logger.Log.Infoln("[PROFILE][GET] success")
	c.JSON(http.StatusOK, data)
}
