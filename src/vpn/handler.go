package vpn

import (
	"context"
	"sync"
	"time"

	"github.com/zerops-io/zcli/src/daemonStorage"
	"github.com/zerops-io/zcli/src/grpcApiClientFactory"
	"github.com/zerops-io/zcli/src/utils/logger"
)

type localDnsManagement string

const (
	wireguardPort  = "51820"
	vpnApiGrpcPort = ":64510"

	localDnsManagementSystemdResolve localDnsManagement = "SYSTEMD_RESOLVE"
	localDnsManagementResolveConf    localDnsManagement = "RESOLVCONF"
	localDnsManagementFile           localDnsManagement = "FILE"

	resolvFilePath          = "/etc/resolv.conf"
	resolvconfOrderFilePath = "/etc/resolvconf/interface-order"
)

type Config struct {
	VpnCheckInterval   time.Duration
	VpnCheckRetryCount int
}

type Handler struct {
	config               Config
	logger               logger.Logger
	grpcApiClientFactory *grpcApiClientFactory.Handler
	storage              *daemonStorage.Handler

	lock sync.RWMutex
}

func New(
	config Config,
	logger logger.Logger,
	grpcApiClientFactory *grpcApiClientFactory.Handler,
	daemonStorage *daemonStorage.Handler,
) *Handler {
	return &Handler{
		config:               config,
		logger:               logger,
		grpcApiClientFactory: grpcApiClientFactory,
		storage:              daemonStorage,
	}
}

func (h *Handler) Run(ctx context.Context) error {
	t := time.NewTicker(h.config.VpnCheckInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			h.checkStatus(ctx)
		}
	}
}
