package specw

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Env[T any] struct {
	Value T
}

func (e *Env[T]) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected string, got %q", n.Kind)
	}

	varValue, isVar, err := e.resolveVarValue(n.Value)
	if err != nil {
		return fmt.Errorf("resolve variable: %w", err)
	}

	n.Value = varValue

	if isVar {
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

	varValue, _, err := e.resolveVarValue(v)
	if err != nil {
		return fmt.Errorf("resolve variable: %w", err)
	}

	err = json.Unmarshal([]byte(strconv.Quote(varValue)), &e.Value)
	if err != nil {
		return err
	}

	return nil
}

func (e *Env[T]) resolveVarValue(v string) (string, bool, error) {
	if !strings.HasPrefix(v, "$") {
		return v, false, nil
	}

	varName := e.resolveVarName(v)

	val, ok := os.LookupEnv(varName)
	if !ok {
		return "", true, fmt.Errorf("environment variable %s not found", varName)
	}

	return val, true, nil
}

func (*Env[T]) resolveVarName(v string) string {
	varName := strings.TrimPrefix(v, "$")
	varName = strings.TrimPrefix(varName, "{")
	varName = strings.TrimSuffix(varName, "}")

	return varName
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
