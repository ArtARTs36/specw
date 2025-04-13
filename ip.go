package specw

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"strings"

	"gopkg.in/yaml.v3"
)

type IP struct {
	net.IP
}

func (i *IP) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected string, got %q", n.Kind)
	}

	return i.unmarshalString(n.Value)
}

func (i *IP) UnmarshalJSON(data []byte) error {
	v := string(data)

	if !strings.HasPrefix(v, "\"") || !strings.HasSuffix(v, "\"") {
		return errors.New("expected string")
	}

	v = strings.Trim(v, "\"")

	return i.unmarshalString(v)
}

func (i *IP) unmarshalString(value string) error {
	addr, err := netip.ParseAddr(value)
	if err != nil {
		return fmt.Errorf("invalid ip: %w", err)
	}
	if addr.Zone() != "" {
		return fmt.Errorf("invalid ip: address with zone")
	}

	addrBytes := addr.As16()

	i.IP = addrBytes[:]

	return nil
}
