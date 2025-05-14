package specw

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnv_UnmarshallYAML(t *testing.T) {
	t.Run("string: without env", func(t *testing.T) {
		var spec struct {
			Val Env[string] `yaml:"val"`
		}

		err := yaml.Unmarshal([]byte("{val: 123}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, "123", spec.Val.Value)
	})

	t.Run("string: with env", func(t *testing.T) {
		var spec struct {
			Val Env[string] `yaml:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "123")

		err := yaml.Unmarshal([]byte("{val: '${SPECW_ENV_VAR}'}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, "123", spec.Val.Value)
	})

	t.Run("url: without env", func(t *testing.T) {
		var spec struct {
			Val Env[URL] `yaml:"val"`
		}

		err := yaml.Unmarshal([]byte("{val: http://google.ru}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, "http://google.ru", spec.Val.Value.String())
	})

	t.Run("url: with env", func(t *testing.T) {
		var spec struct {
			Val Env[URL] `yaml:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "http://google.ru")

		err := yaml.Unmarshal([]byte("{val: '${SPECW_ENV_VAR}'}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, "http://google.ru", spec.Val.Value.String())
	})

	t.Run("float64: without env", func(t *testing.T) {
		var spec struct {
			Val Env[float64] `yaml:"val"`
		}

		err := yaml.Unmarshal([]byte("{val: 5.0}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, float64(5), spec.Val.Value)
	})

	t.Run("float64: with env from int", func(t *testing.T) {
		var spec struct {
			Val Env[float64] `yaml:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "5")

		err := yaml.Unmarshal([]byte("{ val: $SPECW_ENV_VAR }"), &spec)
		require.NoError(t, err)
		assert.Equal(t, float64(5), spec.Val.Value)
	})

	t.Run("float64: with env from float", func(t *testing.T) {
		var spec struct {
			Val Env[float64] `yaml:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "5.5")

		err := yaml.Unmarshal([]byte("{ val: $SPECW_ENV_VAR }"), &spec)
		require.NoError(t, err)
		assert.Equal(t, 5.5, spec.Val.Value)
	})

	t.Run("int: with env", func(t *testing.T) {
		var spec struct {
			Val Env[int] `yaml:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "5")

		err := yaml.Unmarshal([]byte("{ val: $SPECW_ENV_VAR }"), &spec)
		require.NoError(t, err)
		assert.Equal(t, 5, spec.Val.Value)
	})

	t.Run("false: with env", func(t *testing.T) {
		var spec struct {
			Val Env[bool] `yaml:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "false")

		err := yaml.Unmarshal([]byte("{ val: $SPECW_ENV_VAR }"), &spec)
		require.NoError(t, err)
		assert.Equal(t, false, spec.Val.Value)
	})

	t.Run("true: with env", func(t *testing.T) {
		var spec struct {
			Val Env[bool] `yaml:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "true")

		err := yaml.Unmarshal([]byte("{ val: $SPECW_ENV_VAR }"), &spec)
		require.NoError(t, err)
		assert.Equal(t, true, spec.Val.Value)
	})

	t.Run("duration: with env", func(t *testing.T) {
		var spec struct {
			Val Env[Duration] `yaml:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "5s")

		err := yaml.Unmarshal([]byte("{ val: $SPECW_ENV_VAR }"), &spec)
		require.NoError(t, err)
		assert.Equal(t, Duration{Value: 5 * time.Second}, spec.Val.Value)
	})
}

func TestEnv_UnmarshallJSON(t *testing.T) {
	t.Run("string: without env", func(t *testing.T) {
		var spec struct {
			Val Env[string] `json:"val"`
		}

		err := json.Unmarshal([]byte("{\"val\": \"abcd\"}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, "abcd", spec.Val.Value)
	})

	t.Run("string: with env", func(t *testing.T) {
		var spec struct {
			Val Env[string] `json:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "abcd")

		err := json.Unmarshal([]byte("{\"val\": \"${SPECW_ENV_VAR}\"}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, "abcd", spec.Val.Value)
	})

	t.Run("url: without env", func(t *testing.T) {
		var spec struct {
			Val Env[URL] `json:"val"`
		}

		err := json.Unmarshal([]byte("{\"val\": \"http://google.ru\"}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, "http://google.ru", spec.Val.Value.String())
	})

	t.Run("url: with env", func(t *testing.T) {
		var spec struct {
			Val Env[URL] `json:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "http://google.ru")

		err := json.Unmarshal([]byte("{\"val\": \"${SPECW_ENV_VAR}\"}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, "http://google.ru", spec.Val.Value.String())
	})

	t.Run("float64: with env", func(t *testing.T) {
		var spec struct {
			Val Env[float64] `json:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "5")

		err := json.Unmarshal([]byte("{\"val\": \"${SPECW_ENV_VAR}\"}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, float64(5), spec.Val.Value)
	})

	t.Run("int: with env", func(t *testing.T) {
		var spec struct {
			Val Env[int] `json:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "5")

		err := json.Unmarshal([]byte("{\"val\": \"${SPECW_ENV_VAR}\"}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, 5, spec.Val.Value)
	})

	t.Run("false: with env", func(t *testing.T) {
		var spec struct {
			Val Env[bool] `json:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "false")

		err := json.Unmarshal([]byte("{\"val\": \"${SPECW_ENV_VAR}\"}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, false, spec.Val.Value)
	})

	t.Run("true: with env", func(t *testing.T) {
		var spec struct {
			Val Env[bool] `json:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "true")

		err := json.Unmarshal([]byte("{\"val\": \"${SPECW_ENV_VAR}\"}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, true, spec.Val.Value)
	})

	t.Run("duration: with env", func(t *testing.T) {
		var spec struct {
			Val Env[Duration] `json:"val"`
		}

		t.Setenv("SPECW_ENV_VAR", "5s")

		err := json.Unmarshal([]byte("{\"val\": \"${SPECW_ENV_VAR}\"}"), &spec)
		require.NoError(t, err)
		assert.Equal(t, Duration{Value: 5 * time.Second}, spec.Val.Value)
	})
}

func TestEnv_resolveVar(t *testing.T) {
	tests := []struct {
		Title    string
		Input    string
		Expected string
	}{
		{
			Title:    "$VAR",
			Input:    "$VAR",
			Expected: "VAR",
		},
		{
			Title:    "${VAR}",
			Input:    "${VAR}",
			Expected: "VAR",
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			env := &Env[string]{}

			assert.Equal(t, test.Expected, env.resolveVarName(test.Input))
		})
	}
}
