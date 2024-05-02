package api

import (
	"time"

	v1 "api-gateway/api/handlers/v1"
	"api-gateway/api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	grpcClients "api-gateway/internal/infrastructure/grpc_service_client"
	"api-gateway/internal/pkg/config"
)

type RouteOption struct {
	Config         *config.Config
	Logger         *zap.Logger
	ContextTimeout time.Duration
	Service        grpcClients.ServiceClient
}

// NewRoute
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRoute(option RouteOption) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	HandlerV1 := v1.New(&v1.HandlerV1Config{
		Config:         option.Config,
		Logger:         option.Logger,
		ContextTimeout: option.ContextTimeout,
		Service:        option.Service,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	router.Use(middleware.Tracing)

	apiV1 := router.Group("/v1")

	// clients
	apiV1.POST("/client", HandlerV1.CreateClient)
	apiV1.PUT("/client", HandlerV1.UpdateClient)
	apiV1.DELETE("/client/:id", HandlerV1.DeleteClient)
	apiV1.GET("/client/:id", HandlerV1.GetClient)

	// jobs
	apiV1.POST("/job", HandlerV1.CreateJob)
	apiV1.PUT("/job", HandlerV1.UpdateJob)
	apiV1.DELETE("/job/:id", HandlerV1.DeleteJob)
	apiV1.GET("/job/:id", HandlerV1.GetJob)
	apiV1.POST("/job/add-client", HandlerV1.AddClientToJob)
	apiV1.DELETE("/job/remove-client", HandlerV1.RemoveClientFromJob)
	apiV1.POST("/jobs/client-jobs", HandlerV1.GetClientsWithJob)

	url := ginSwagger.URL("swagger/doc.json")
	apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
