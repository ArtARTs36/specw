package specw

import (
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPositiveNumber_UnmarshalYAML(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var spec struct {
			Value PositiveNumber[int] `yaml:"value"`
		}

		content := "value: 10"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)
		assert.Equal(t, 10, spec.Value.Value)
	})

	t.Run("zero", func(t *testing.T) {
		var spec struct {
			Value PositiveNumber[int] `yaml:"value"`
		}

		content := "value: 0"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.Error(t, err)
		assert.ErrorContains(t, err, "must be positive")
	})

	t.Run("negative", func(t *testing.T) {
		var spec struct {
			Value PositiveNumber[int] `yaml:"value"`
		}

		content := "value: -5"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.Error(t, err)
		assert.ErrorContains(t, err, "must be positive")
	})

	t.Run("decode error", func(t *testing.T) {
		var spec struct {
			Value PositiveNumber[int] `yaml:"value"`
		}

		content := "value: abc"

		err := yaml.Unmarshal([]byte(content), &spec)
		require.Error(t, err)
		assert.ErrorContains(t, err, "decode number")
	})
}
