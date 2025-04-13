package specw

import (
	"encoding/json"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestURL_UnmarshallYAML(t *testing.T) {
	var spec struct {
		URL URL `yaml:"url"`
	}

	content := "{url: http://google.com}"

	err := yaml.Unmarshal([]byte(content), &spec)
	require.NoError(t, err)
	assert.Equal(t, "http://google.com", spec.URL.String())
}

func TestURL_UnmarshallJSON(t *testing.T) {
	var spec struct {
		URL URL `json:"url"`
	}

	content := "{\"url\": \"http://google.com\"}"

	err := json.Unmarshal([]byte(content), &spec)
	require.NoError(t, err)
	assert.Equal(t, "http://google.com", spec.URL.String())
}
