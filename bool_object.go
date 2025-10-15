package specw

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"reflect"
)

type BoolObject[O any] struct {
	Object *O
}

func (o *BoolObject[O]) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind == yaml.ScalarNode {
		switch n.Value {
		case "true":
			var err error
			o.Object, err = o.instance()
			return err
		case "false":
			o.Object = nil
			return nil
		}

		return fmt.Errorf("unexpected value: %s", n.Value)
	}

	if n.Kind == yaml.MappingNode {
		if err := n.Decode(&o.Object); err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("unexpected type: %q", n.Kind)
}

func (o *BoolObject[O]) instance() (*O, error) {
	var instance O

	val, ok := reflect.New(reflect.TypeOf(instance)).Interface().(*O)
	if !ok {
		return nil, errors.New("unable to create object")
	}

	return val, nil
}
