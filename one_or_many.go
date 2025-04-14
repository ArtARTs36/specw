package specw

import "gopkg.in/yaml.v3"

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
