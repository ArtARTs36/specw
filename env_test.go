package specw

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"

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
}
