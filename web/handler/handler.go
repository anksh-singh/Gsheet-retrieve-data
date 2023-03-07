package handler

import (
	"sheet-retrieve/config"
	grpcClient "sheet-retrieve/pkg/grpc/client"
	"sheet-retrieve/pkg/grpc/proto/pb"
	"sheet-retrieve/utils"
	"time"

	"go.uber.org/zap"
)

type handler struct {
	grpcClient pb.BetaUsersClient
	config     *config.Config
	logger     *zap.SugaredLogger
	util       *utils.UtilConf
}

const (
	//defaultTimeOutMills API time out
	defaultTimeOutMills = 10 * time.Second
	// TimeOut25 TimeOut25Seconds
	TimeOut25 = 25 * time.Second
	//TimeOutOneMinute Special case for bridge APIs
	TimeOutOneMinute = 60 * time.Second
	//defaultSuccessMsg A default API success response
	defaultSuccessMsg = "OK"
)

func NewHandler(config *config.Config, logger *zap.SugaredLogger, gc *grpcClient.GrpcClientManager) *handler {
	utilConf := utils.NewUtils(logger, config)

	return &handler{
		grpcClient: gc.GetGrpcClient(),
		config:     config,
		logger:     logger,
		util:       utilConf,
	}
}
