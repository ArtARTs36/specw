package specw

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var fileReader = os.ReadFile

type File struct {
	Path    string
	Content []byte
}

func SetFileReader(fn func(path string) ([]byte, error)) {
	if fn == nil {
		panic("specw: SetFileReader called with nil func")
	}

	fileReader = fn
}

func (f *File) UnmarshalYAML(n *yaml.Node) error {
	path := ""

	if err := n.Decode(&path); err != nil {
		return fmt.Errorf("parse file path: %w", err)
	}

	content, err := fileReader(path)
	if err != nil {
		return fmt.Errorf("read file %q: %w", path, err)
	}

	f.Path = path
	f.Content = content

	return nil
}

func (f *File) MarshalYAML() (interface{}, error) {
	return f.Path, nil
}

func (f *File) IsEmpty() bool {
	return len(f.Content) == 0
}

func (f *File) String() string {
	return string(f.Content)
}
