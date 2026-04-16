package specw

import (
	"encoding/json"
	"fmt"
	"time"
	"unicode"

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

func (d *Duration) UnmarshalBinary(bytes []byte) error {
	return d.UnmarshalString(string(bytes))
}

func (d *Duration) UnmarshalText(bytes []byte) error {
	return d.UnmarshalBinary(bytes)
}

func (d *Duration) UnmarshalString(val string) error {
	number, isNumber := extractNumber(val)
	if isNumber {
		d.Value = time.Duration(number)
		return nil
	}

	var err error
	d.Value, err = time.ParseDuration(val)
	if err != nil {
		return fmt.Errorf("parse duration: %w", err)
	}

	return nil
}

func extractNumber(val string) (int, bool) {
	if val == "" {
		return 0, false
	}

	number := 0

	for _, c := range val {
		if !unicode.IsDigit(c) {
			return 0, false
		}

		if number == 0 {
			number = int(c - '0')
		} else {
			number = number*10 + int(c-'0') //nolint:mnd // not need
		}
	}

	return number, true
}
