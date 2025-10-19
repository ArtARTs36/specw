package specw

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestBoolObject_UnmarshalYAML(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		var spec struct {
			Value BoolObject[struct {
				Title string `yaml:"title"`
			}] `yaml:"value"`
		}

		content := `value: false`

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		assert.Nil(t, spec.Value.Object)
	})

	t.Run("true", func(t *testing.T) {
		var spec struct {
			Value BoolObject[struct {
				Title string `yaml:"title"`
			}] `yaml:"value"`
		}

		content := `value: true`

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		assert.NotNil(t, spec.Value.Object)
	})

	t.Run("object", func(t *testing.T) {
		var spec struct {
			Value BoolObject[struct {
				Title string `yaml:"title"`
			}] `yaml:"value"`
		}

		content := `value: {title: abc}`

		err := yaml.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		require.NotNil(t, spec.Value.Object)
		assert.Equal(t, "abc", spec.Value.Object.Title)
	})
}

func TestBoolObject_UnmarshalJSON(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		var spec struct {
			Value BoolObject[struct {
				Title string `json:"title"`
			}] `json:"value"`
		}

		content := `{"value": false}`

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		assert.Nil(t, spec.Value.Object)
	})

	t.Run("true", func(t *testing.T) {
		var spec struct {
			Value BoolObject[struct {
				Title string `json:"title"`
			}] `json:"value"`
		}

		content := `{"value": true}`

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		assert.NotNil(t, spec.Value.Object)
	})

	t.Run("object", func(t *testing.T) {
		var spec struct {
			Value BoolObject[struct {
				Title string `json:"title"`
			}] `json:"value"`
		}

		content := `{"value": {"title": "abc"}}`

		err := json.Unmarshal([]byte(content), &spec)
		require.NoError(t, err)

		require.NotNil(t, spec.Value.Object)
		assert.Equal(t, "abc", spec.Value.Object.Title)
	})
}
