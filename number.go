package specw

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
)

type PositiveNumber[T int | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64 | time.Duration] struct { //nolint:lll // not need
	Value T
}

func (e *PositiveNumber[T]) UnmarshalYAML(n *yaml.Node) error {
	err := n.Decode(&e.Value)
	if err != nil {
		return fmt.Errorf("decode number: %w", err)
	}

	if e.Value <= 0 {
		return errors.New("value must be positive")
	}

	return nil
}
