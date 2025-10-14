package specw

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/buildkite/interpolate"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Env[T any] struct {
	Value T
}

type interpolateEnv struct{}

func (e interpolateEnv) Get(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (e *Env[T]) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected string, got %q", n.Kind)
	}

	resolvedValue, err := interpolate.Interpolate(interpolateEnv{}, n.Value)
	if err != nil {
		return fmt.Errorf("interpolate value: %w", err)
	}

	changed := n.Value != resolvedValue

	n.Value = resolvedValue

	if changed {
		e.repairNode(n)
	}

	err = n.Decode(&e.Value)
	if err != nil {
		return err
	}

	return nil
}

func (e *Env[T]) UnmarshalJSON(data []byte) error {
	v := string(data)

	if !strings.HasPrefix(v, "\"") || !strings.HasSuffix(v, "\"") {
		return errors.New("expected string")
	}

	v = strings.Trim(v, "\"")

	varValue, err := interpolate.Interpolate(interpolateEnv{}, v)
	if err != nil {
		return fmt.Errorf("resolve variable: %w", err)
	}

	varValue = e.repairJSONValue(varValue)

	err = json.Unmarshal([]byte(varValue), &e.Value)
	if err != nil {
		return err
	}

	return nil
}

func (e *Env[T]) repairJSONValue(varValue string) string {
	var instance T
	expectedType := reflect.TypeOf(instance)

	switch expectedType.Kind() { //nolint:exhaustive // not need
	case reflect.Float64, reflect.Float32,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Bool:
		// skip
	default:
		varValue = strconv.Quote(varValue)
	}

	return varValue
}

func (e *Env[T]) repairNode(n *yaml.Node) {
	var typeInstance T

	expectedType := reflect.TypeOf(typeInstance)

	switch expectedType.Kind() { //nolint:exhaustive // not need
	case reflect.Float64, reflect.Float32:
		n.Tag = "!!float"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n.Tag = "!!int"
	case reflect.Bool:
		n.Tag = "!!bool"
	}
}
