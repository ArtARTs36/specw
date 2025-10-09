package specw

import (
	"bytes"
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type OneOrMany[T any] struct {
	Value []T
}

func (o *OneOrMany[T]) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind == yaml.ScalarNode || n.Kind == yaml.MappingNode {
		var val T

		if err := n.Decode(&val); err != nil {
			return err
		}

		o.Value = []T{val}

		return nil
	}

	return n.Decode(&o.Value)
}

func (o *OneOrMany[T]) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(data, []byte{'['}) {
		return json.Unmarshal(data, &o.Value)
	}

	var val T
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	o.Value = []T{val}

	return nil
}
