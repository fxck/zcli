package vpn

import (
	"os/exec"
)

func (h *Handler) isVpnAlive(serverIp string) bool {

	if serverIp == "" {
		return false
	}

	for i := 0; i < 3; i++ {
		_, err := exec.Command("ping6", serverIp, "-c", "1").Output()
		if err != nil {
			continue
		}

		return true
	}

	return false
}
