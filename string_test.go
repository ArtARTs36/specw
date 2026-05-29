package specw

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestString_UnmarshalYAML(t *testing.T) {
	var spec struct {
		String String `yaml:"string"`
	}

	err := yaml.Unmarshal([]byte(`string: ""`), &spec)
	require.Error(t, err)
	assert.Equal(t, "provided empty string", err.Error())
}
