package web

import (
	"sheet-retrieve/config"
	grpcClient "sheet-retrieve/pkg/grpc/client"
	"sheet-retrieve/web/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.uber.org/zap"
)

func InitRoutes(config *config.Config, logger *zap.SugaredLogger, route *gin.Engine) {
	grpcClientManager := grpcClient.NewGrpcClientManager(config, logger)
	webHandler := handler.NewHandler(config, logger, grpcClientManager)

	groupRoute := route.Group("/api")
	groupRoute.GET("/retrieve-sheet", webHandler.RetrieveSheet)

	//swagger api
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
