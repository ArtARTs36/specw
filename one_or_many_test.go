package specw

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestOneOrMany_UnmarshalYAML(t *testing.T) {
	t.Run("scalar: one", func(t *testing.T) {
		type specDefinition struct {
			Values OneOrMany[string] `yaml:"values"`
		}

		content := "{values: a}"

		var spec specDefinition

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("scalar: many", func(t *testing.T) {
		type specDefinition struct {
			Values OneOrMany[string] `yaml:"values"`
		}

		content := "{values: [a, b]}"

		var spec specDefinition

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a", "b"}, spec.Values.Value)
	})

	t.Run("mapping: one", func(t *testing.T) {
		type specDefinitionNested struct {
			A string `yaml:"a"`
		}

		type specDefinition struct {
			Values OneOrMany[specDefinitionNested] `yaml:"values"`
		}

		content := "{values: {a: 1}}"

		var spec specDefinition

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		assert.Equal(t, []specDefinitionNested{
			{
				A: "1",
			},
		}, spec.Values.Value)
	})

	t.Run("mapping: many", func(t *testing.T) {
		type specDefinitionNested struct {
			A string `yaml:"a"`
		}

		type specDefinition struct {
			Values OneOrMany[specDefinitionNested] `yaml:"values"`
		}

		content := "{values: [{a: 1}, {a: 2}]}"

		var spec specDefinition

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		assert.Equal(t, []specDefinitionNested{
			{
				A: "1",
			},
			{
				A: "2",
			},
		}, spec.Values.Value)
	})
}

func TestOneOrMany_UnmarshalJSON(t *testing.T) {
	t.Run("scalar: one", func(t *testing.T) {
		type specDefinition struct {
			Values OneOrMany[string] `json:"values"`
		}

		content := "{\"values\": \"a\"}"

		var spec specDefinition

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("scalar: many", func(t *testing.T) {
		type specDefinition struct {
			Values OneOrMany[string] `json:"values"`
		}

		content := "{\"values\": [\"a\", \"b\"]}"

		var spec specDefinition

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a", "b"}, spec.Values.Value)
	})

	t.Run("slice: one", func(t *testing.T) {
		type specDefinitionNested struct {
			A string `json:"a"`
		}

		type specDefinition struct {
			Values OneOrMany[specDefinitionNested] `json:"values"`
		}

		content := "{\"values\": {\"a\": \"1\"}}"

		var spec specDefinition

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		assert.Equal(t, []specDefinitionNested{
			{
				A: "1",
			},
		}, spec.Values.Value)
	})

	t.Run("slice: many", func(t *testing.T) {
		type specDefinitionNested struct {
			A string `json:"a"`
		}

		type specDefinition struct {
			Values OneOrMany[specDefinitionNested] `json:"values"`
		}

		content := "{\"values\": [{\"a\": \"1\"}, {\"a\": \"2\"}]}"

		var spec specDefinition

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		assert.Equal(t, []specDefinitionNested{
			{
				A: "1",
			},
			{
				A: "2",
			},
		}, spec.Values.Value)
	})
}
