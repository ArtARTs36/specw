package specw

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

type Duration struct {
	Value time.Duration
}

func (d *Duration) UnmarshalYAML(n *yaml.Node) error {
	var v interface{}
	if err := n.Decode(&v); err != nil {
		return err
	}
	switch value := v.(type) {
	case int:
		d.Value = time.Duration(value)
		return nil
	case string:
		var err error
		d.Value, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unexpected type: %T", value)
	}
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Value = time.Duration(value)
		return nil
	case string:
		var err error
		d.Value, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unexpected type: %T", value)
	}
}
