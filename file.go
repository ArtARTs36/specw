package specw

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type File struct {
	Path    string
	Content []byte
}

func (f *File) UnmarshalYAML(n *yaml.Node) error {
	path := ""

	if err := n.Decode(&path); err != nil {
		return fmt.Errorf("parse file path: %w", err)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file %q: %w", path, err)
	}

	f.Path = path
	f.Content = content

	return nil
}

func (f *File) IsEmpty() bool {
	return len(f.Content) == 0
}

func (f *File) String() string {
	return string(f.Content)
}
