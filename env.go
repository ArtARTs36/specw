package specw

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
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

	varValue, err := e.resolveVar(n.Value)
	if err != nil {
		return fmt.Errorf("resolve variable: %w", err)
	}

	n.Value = varValue

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

	varValue, err := e.resolveVar(v)
	if err != nil {
		return fmt.Errorf("resolve variable: %w", err)
	}

	err = json.Unmarshal([]byte(strconv.Quote(varValue)), &e.Value)
	if err != nil {
		return err
	}

	return nil
}

func (*Env[T]) resolveVar(v string) (string, error) {
	if !strings.HasPrefix(v, "${") || !strings.HasSuffix(v, "}") {
		return v, nil
	}

	varName := strings.TrimPrefix(v, "${")
	varName = strings.TrimSuffix(varName, "}")

	val, ok := os.LookupEnv(varName)
	if !ok {
		return "", fmt.Errorf("environment variable %s not found", varName)
	}

	return val, nil
}
