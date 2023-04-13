package handler

import (
	"dummyCVForm/models"
	"dummyCVForm/pkg/logger"
	"dummyCVForm/utils/constants"
	"dummyCVForm/utils/random"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProfileControllers struct {
	Usecase models.ProfileUsecase
}

func NewProfileControllers(g *gin.RouterGroup, PUsecase models.ProfileUsecase) {
	handler := &ProfileControllers{Usecase: PUsecase}
	g.GET("/profile/:id", handler.GetProfile)
	g.POST("/profile", handler.AddProfile)
	g.PUT("/profile/:id", handler.EditProfile)
}

func (r *ProfileControllers) GetProfile(c *gin.Context) {
	id := c.Param("id")
	data, err := r.Usecase.Get(c, id)
	if err != nil {
		logger.Log.Errorf("[PROFILE][GET] ERROR %v for requestId: %v", err.Error(), requestid.Get(c))
		c.JSON(http.StatusInternalServerError, models.CreateResponse(c, constants.InternalServerCode, constants.InternalServerError, constants.WarnInternalError, err.Error()))
		return
	}

	logger.Log.Infof("[PROFILE][GET] success for requestId: %v", requestid.Get(c))
	c.JSON(http.StatusOK, data)
}

func (r *ProfileControllers) AddProfile(c *gin.Context) {
	var req models.Profile
	req.ProfileCode = random.RandNumber()

	// Bind Request
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Errorf("[PROFILE][ADD] ERROR %v for requestId: %v", err.Error(), requestid.Get(c))
		c.JSON(http.StatusInternalServerError, models.CreateResponse(c, constants.InternalServerCode, constants.InternalServerError, constants.WarnInternalError, err.Error()))
		return
	}

	err := r.Usecase.Create(c, &req)
	if err != nil {
		logger.Log.Errorf("[PROFILE][ADD] ERROR %v for requestId: %v", err.Error(), requestid.Get(c))
		c.JSON(http.StatusInternalServerError, models.CreateResponse(c, constants.InternalServerCode, constants.InternalServerError, constants.WarnInternalError, err.Error()))
		return
	}

	logger.Log.Infof("[PROFILE][ADD] success for requestId: %v", requestid.Get(c))
	c.JSON(http.StatusOK, &models.Profile{ProfileCode: req.ProfileCode})
}

func (r *ProfileControllers) EditProfile(c *gin.Context) {
	var req models.Profile
	id := c.Param("id")
	code, _ := strconv.Atoi(id)

	// Bind Request
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Errorf("[PROFILE][EDIT] ERROR %v for requestId: %v", err.Error(), requestid.Get(c))
		c.JSON(http.StatusInternalServerError, models.CreateResponse(c, constants.InternalServerCode, constants.InternalServerError, constants.WarnInternalError, err.Error()))
		return
	}

	err := r.Usecase.Update(c, &req)
	if err != nil {
		logger.Log.Errorf("[PROFILE][EDIT] ERROR %v for requestId: %v", err.Error(), requestid.Get(c))
		c.JSON(http.StatusInternalServerError, models.CreateResponse(c, constants.InternalServerCode, constants.InternalServerError, constants.WarnInternalError, err.Error()))
		return
	}

	logger.Log.Infof("[PROFILE][EDIT] success for requestId: %v", requestid.Get(c))
	c.JSON(http.StatusOK, &models.Profile{ProfileCode: code})
}
