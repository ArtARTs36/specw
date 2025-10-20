package specw

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvStrings_UnmarshalYAML(t *testing.T) {
	t.Run("scalar: separated one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `yaml:"values"`
		}

		content := "{values: 'a,b'}"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a", "b"}, spec.Values.Value)
	})

	t.Run("scalar: simple one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `yaml:"values"`
		}

		content := "{values: a}"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("scalar: env one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `yaml:"values"`
		}

		t.Setenv("SPECW_ENV_VAR", "a")

		content := "{values: $SPECW_ENV_VAR}"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("sequence: simple one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `yaml:"values"`
		}

		content := "{values: [a]}"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("sequence: env one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `yaml:"values"`
		}

		t.Setenv("SPECW_ENV_VAR", "a")

		content := "{values: [$SPECW_ENV_VAR]}"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("sequence: env and single one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `yaml:"values"`
		}

		t.Setenv("SPECW_ENV_VAR", "a")

		content := "{values: [b, $SPECW_ENV_VAR]}"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"b", "a"}, spec.Values.Value)
	})
}

func TestEnvStrings_UnmarshalJSON(t *testing.T) {
	t.Run("scalar: separated one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `json:"values"`
		}

		content := "{\"values\": \"a,b\"}"

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a", "b"}, spec.Values.Value)
	})

	t.Run("scalar: simple one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `json:"values"`
		}

		content := "{\"values\": \"a\"}"

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("scalar: env one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `json:"values"`
		}

		t.Setenv("SPECW_ENV_VAR", "a")

		content := "{\"values\": \"$SPECW_ENV_VAR\"}"

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("sequence: simple one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `json:"values"`
		}

		content := "{\"values\": [\"a\"]}"

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("sequence: env one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `json:"values"`
		}

		t.Setenv("SPECW_ENV_VAR", "a")

		content := "{\"values\": [\"$SPECW_ENV_VAR\"]}"

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"a"}, spec.Values.Value)
	})

	t.Run("sequence: env and single one", func(t *testing.T) {
		var spec struct {
			Values EnvStrings `json:"values"`
		}

		t.Setenv("SPECW_ENV_VAR", "a")

		content := "{\"values\": [\"b\", \"$SPECW_ENV_VAR\"]}"

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, []string{"b", "a"}, spec.Values.Value)
	})
}
