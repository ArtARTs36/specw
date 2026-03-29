package specw

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YAMLFile[T any] struct {
	Path  string
	Value T
}

func (f *YAMLFile[T]) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected path string, got %v", n.Kind)
	}

	path := ""

	if err := n.Decode(&path); err != nil {
		return fmt.Errorf("parse file path: %w", err)
	}

	f.Path = path

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file %q: %w", path, err)
	}

	if err = yaml.Unmarshal(content, &f.Value); err != nil {
		return fmt.Errorf("decode embedded yaml: %w", err)
	}

	return nil
}
