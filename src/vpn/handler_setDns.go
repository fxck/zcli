package vpn

import (
	"net"
	"os/exec"
	"strings"

	"github.com/zerops-io/zcli/src/utils"

	"github.com/zerops-io/zcli/src/utils/cmdRunner"
)

func (h *Handler) setDns(dnsIp net.IP, dnsManagement localDnsManagement) error {
	var err error

	if dnsManagement == localDnsManagementSystemdResolve {
		_, err = cmdRunner.Run(exec.Command("systemd-resolve", "--set-dns="+dnsIp.String(), `--set-domain=local`, "--interface=wg0"))
		if err != nil {
			return err
		}
	}

	if dnsManagement == localDnsManagementResolveConf {
		err := utils.SetFirstLine(resolvconfOrderFilePath, "wg*")
		if err != nil {
			return err
		}

		cmd := exec.Command("resolvconf", "-a", "wg0")
		cmd.Stdin = strings.NewReader("nameserver " + dnsIp.String())
		_, err = cmdRunner.Run(cmd)
		if err != nil {
			return err
		}
	}

	if dnsManagement == localDnsManagementFile {
		err := utils.SetFirstLine(resolvFilePath, "nameserver "+dnsIp.String())
		if err != nil {
			return err
		}

	}

	return nil
}
