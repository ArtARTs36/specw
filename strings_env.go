package specw

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buildkite/interpolate"
	"gopkg.in/yaml.v3"
	"strings"
)

type EnvStrings struct {
	Value []string
}

func (s *EnvStrings) UnmarshalYAML(n *yaml.Node) error {
	switch n.Kind { //nolint:exhaustive // other types not supported
	case yaml.ScalarNode:
		v, err := interpolate.Interpolate(interpolateEnv{}, n.Value)
		if err != nil {
			return fmt.Errorf("interpolate: %w", err)
		}

		s.Value = strings.Split(v, ",")
		return nil
	case yaml.SequenceNode:
		for _, child := range n.Content {
			v, err := interpolate.Interpolate(interpolateEnv{}, child.Value)
			if err != nil {
				return fmt.Errorf("interpolate: %w", err)
			}

			s.Value = append(s.Value, v)
		}
		return nil
	}

	return fmt.Errorf("unexpected node type: %q", n.Kind)
}

func (s *EnvStrings) UnmarshalJSON(n []byte) error {
	if bytes.HasPrefix(n, []byte{'"'}) {
		var v string

		if err := json.Unmarshal(n, &v); err != nil {
			return err
		}

		v, err := interpolate.Interpolate(interpolateEnv{}, v)
		if err != nil {
			return fmt.Errorf("interpolate: %w", err)
		}

		s.Value = strings.Split(v, ",")
		return nil
	}

	var values []string

	if err := json.Unmarshal(n, &values); err != nil {
		return err
	}

	for i, value := range values {
		v, err := interpolateEnvExpression(value)
		if err != nil {
			return fmt.Errorf("interpolate: %w", err)
		}

		values[i] = v
	}

	s.Value = values
	return nil
}
