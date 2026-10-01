package specw

import (
	"encoding/json"
	"errors"
	"fmt"
    "reflect"

	"go.yaml.in/yaml/v3"
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

	return fmt.Errorf("unexpected type: %q", n.Tag)
}

func (o *BoolObject[O]) UnmarshalJSON(data []byte) error {
	str := string(data)
	switch str {
	case "true":
		var err error
		o.Object, err = o.instance()
		return err
	case "false":
		o.Object = nil
		return nil
	}

	return json.Unmarshal(data, &o.Object)
}

func (o *BoolObject[O]) instance() (*O, error) {
	var instance O

	val, ok := reflect.New(reflect.TypeOf(instance)).Interface().(*O)
	if !ok {
		return nil, errors.New("unable to create object")
	}

	return val, nil
}
