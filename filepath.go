package specw

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ExistsFilepath struct {
	Value string
}

func (e *ExistsFilepath) UnmarshalYAML(n *yaml.Node) error {
	if err := n.Decode(&e.Value); err != nil {
		return fmt.Errorf("decode filepath: %w", err)
	}

	_, err := os.Stat(e.Value)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file %s does not exist", e.Value)
		}
		return fmt.Errorf("stat file %q: %w", e.Value, err)
	}

	return nil
}
