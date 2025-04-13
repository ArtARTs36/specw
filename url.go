package specw

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"gopkg.in/yaml.v3"
)

type URL struct {
	Value url.URL
}

func (u *URL) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected string, got %q", n.Kind)
	}

	value, err := url.Parse(n.Value)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}

	u.Value = *value

	return err
}

func (u *URL) UnmarshalJSON(data []byte) error {
	v := string(data)

	if !strings.HasPrefix(v, "\"") || !strings.HasSuffix(v, "\"") {
		return errors.New("expected string")
	}

	v = strings.Trim(v, "\"")

	value, err := url.Parse(v)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}

	u.Value = *value

	return nil
}

func (u *URL) String() string {
	return u.Value.String()
}
