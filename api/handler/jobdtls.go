package handler

import (
	"dummyCVForm/models"
	"dummyCVForm/pkg/logger"
	"dummyCVForm/utils/constants"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JobControllers struct {
	Usecase models.JobDtlsUsecase
}

func NewJobControllers(g *gin.RouterGroup, JUsecase models.JobDtlsUsecase) {
	handler := &JobControllers{Usecase: JUsecase}
	g.GET("/employment/:id", handler.GetJobDetails)
}

func (j *JobControllers) GetJobDetails(c *gin.Context) {
	data, err := j.Usecase.Get(c)
	if err != nil {
		logger.Log.Errorf("[EMPLOYMENT][GET] ERROR %v for requestId: %v", err.Error(), requestid.Get(c))
		c.JSON(http.StatusInternalServerError, models.CreateResponse(c, constants.InternalServerCode, constants.InternalServerError, constants.WarnInternalError, err.Error()))
		return
	}

	logger.Log.Infof("[EMPLOYMENT][GET] success for requestId: %v", requestid.Get(c))
	c.JSON(http.StatusOK, data)
}
