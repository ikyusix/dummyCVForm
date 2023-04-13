package handler

import (
	"dummyCVForm/models"
	"dummyCVForm/pkg/logger"
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
		return
	}

	logger.Log.Infof("[EMPLOYMENT][GET] success for requestId: %v", requestid.Get(c))
	c.JSON(http.StatusOK, data)
}
