package modules

import (
	"net"
)

type IP struct {
	Prefix    string
	Suffix    string
	Interface []string
	NoIp      string `yaml:"no_ip"`
}

/**
 * @TODO: Support IPv6, flag in the config to choose which to display
 */
func (ip IP) Print(args []string) string {
	for _, wantedInterface := range ip.Interface {
		iface, err := net.InterfaceByName(wantedInterface)
		if err != nil {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}

	return ip.NoIp
}

func (ip IP) GetPrefix() string {
	return ip.Prefix
}

func (ip IP) GetSuffix() string {
	return ip.Suffix
}
