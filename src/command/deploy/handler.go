package deploy

import (
	"github.com/zerops-io/zcli/src/service/httpClient"
	"github.com/zerops-io/zcli/src/service/logger"
	"github.com/zerops-io/zcli/src/service/zipClient"
	"github.com/zerops-io/zcli/src/zeropsApiProtocol"
)

type Config struct {
}

type RunConfig struct {
	ProjectName      string
	ServiceStackName string
	PathsForPacking  []string
	WorkingDir       string
	ZipFilePath      string
}

type Handler struct {
	config        Config
	httpClient    *httpClient.Handler
	zipClient     *zipClient.Handler
	logger        logger.Logger
	apiGrpcClient zeropsApiProtocol.ZeropsApiProtocolClient
}

func New(
	config Config,
	httpClient *httpClient.Handler,
	zipClient *zipClient.Handler,
	logger logger.Logger,
	apiGrpcClient zeropsApiProtocol.ZeropsApiProtocolClient,
) *Handler {
	return &Handler{
		config:        config,
		httpClient:    httpClient,
		zipClient:     zipClient,
		logger:        logger,
		apiGrpcClient: apiGrpcClient,
	}
}
