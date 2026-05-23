package specw

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

type String string

func (s *String) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected string, got %T", n.Kind)
	}

	if n.Value == "" {
		return errors.New("empty string")
	}
	
	return nil
}
