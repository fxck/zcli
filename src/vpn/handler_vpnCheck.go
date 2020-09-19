package vpn

import "context"

func (h *Handler) checkStatus(ctx context.Context) {

	data := h.storage.Data()

	if data.ProjectId != "" {
		if !h.isVpnAlive(data.ServerIp) {
			err := h.StartVpn(
				ctx,
				data.GrpcApiAddress,
				data.GrpcVpnAddress,
				data.Token,
				data.ProjectId,
			)
			if err != nil {
				h.logger.Error(err)
			}
		}
	}
}
