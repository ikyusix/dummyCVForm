package router

import (
	"dummyCVForm/pkg/controllers"
	"dummyCVForm/pkg/middleware"
	"dummyCVForm/pkg/postgres"
	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// CreateRouter is a function to create initiation router
func CreateRouter(isDev bool) *gin.Engine {

	//host := config.MyConfig.Host
	//port := config.MyConfig.ServerPort
	//url := ginSwagger.URL(host + port + "/swagger/doc.json")

	// Create path url
	router := gin.New()
	//pprof.Register(router)
	// Use middleware
	router.Use(middleware.Secure(isDev))
	router.Use(middleware.CORSMiddleware())
	router.Use(requestid.New())
	router.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		controllers.HandlePanic(c, err)
	}))
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Use(healthcheck.Default())
	router.Use(middleware.RequestLoggerActivity())
	return router
}

func InitRoute(router *gin.Engine) {
	_, err := postgres.GetConnectionDB()
	if err != nil {
		return
	}
	//
	//v1private := router.Group("/v1.0/private")
	//
	//// repositories
	//rr := rptRepo.NewReportingRepositories(db)
	//
	//// usecase
	//ru := rptUc.NewReportingUsecase(rr)
	//
	//// handler
	//rptHdlr.NewReportingControllers(v1private, ru)

}
