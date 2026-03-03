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

	return u.UnmarshalString(n.Value)
}

func (u *URL) UnmarshalJSON(data []byte) error {
	v := string(data)

	if !strings.HasPrefix(v, "\"") || !strings.HasSuffix(v, "\"") {
		return errors.New("expected string")
	}

	v = strings.Trim(v, "\"")

	return u.UnmarshalString(v)
}

func (u *URL) String() string {
	return u.Value.String()
}

func (u *URL) UnmarshalString(val string) error {
	value, err := url.Parse(val)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}

	u.Value = *value

	return nil
}

func (u *URL) UnmarshalBinary(val []byte) error {
	return u.UnmarshalString(string(val))
}
