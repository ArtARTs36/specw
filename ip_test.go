package specw

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIP_UnmarshallYAML(t *testing.T) {
	var spec struct {
		IP IP `yaml:"ip"`
	}

	content := "{ip: 192.168.0.1}"

	err := yaml.Unmarshal([]byte(content), &spec)
	require.NoError(t, err)
	assert.Equal(t, "192.168.0.1", spec.IP.String())
}

func TestIP_UnmarshallJSON(t *testing.T) {
	var spec struct {
		IP IP `json:"ip"`
	}

	content := "{\"ip\": \"192.168.0.1\"}"

	err := json.Unmarshal([]byte(content), &spec)
	require.NoError(t, err)
	assert.Equal(t, "192.168.0.1", spec.IP.String())
}
