package vpn

import (
	"net"
	"os/exec"

	"github.com/zerops-io/zcli/src/utils"

	"github.com/zerops-io/zcli/src/utils/cmdRunner"
)

func (h *Handler) cleanDns(dnsIp net.IP, dnsManagement localDnsManagement) error {
	if dnsManagement == localDnsManagementResolveConf {
		cmd := exec.Command("resolvconf", "-d", "wg0")
		_, err := cmdRunner.Run(cmd)
		if err != nil {
			return err
		}
	}

	if dnsManagement == localDnsManagementFile {
		err := utils.RemoveFirstLine(resolvFilePath, "nameserver "+dnsIp.String())
		if err != nil {
			return err
		}

	}

	return nil
}
