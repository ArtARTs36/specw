package specw

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type JSONFile[T any] struct {
	// Path is a path to JSON file.
	Path string
	// Value is decoded content of JSON file.
	Value T
}

func (f *JSONFile[T]) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected path string, got %v", n.Kind)
	}

	var path string
	if err := n.Decode(&path); err != nil {
		return fmt.Errorf("parse file path: %w", err)
	}

	return f.UnmarshalString(path)
}

func (f *JSONFile[T]) UnmarshalJSON(data []byte) error {
	var path string
	if err := json.Unmarshal(data, &path); err != nil {
		return fmt.Errorf("parse file path: %w", err)
	}

	return f.UnmarshalString(path)
}

func (f *JSONFile[T]) UnmarshalBinary(data []byte) error {
	return f.UnmarshalString(string(data))
}

func (f *JSONFile[T]) UnmarshalText(text []byte) error {
	return f.UnmarshalBinary(text)
}

func (f *JSONFile[T]) UnmarshalString(path string) error {
	f.Path = path

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file %q: %w", path, err)
	}

	if err = json.Unmarshal(content, &f.Value); err != nil {
		return fmt.Errorf("decode embedded json: %w", err)
	}

	return nil
}
