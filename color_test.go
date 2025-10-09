package specw

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestColor_UnmarshalJSON(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		var cfg struct {
			Color Color `json:"color"`
		}

		err := json.Unmarshal([]byte(`{"color":"red"}`), &cfg)
		require.NoError(t, err)
	})
}

func TestColor_Hex(t *testing.T) {
	t.Run("#eeeeee", func(t *testing.T) {
		rgba := color.RGBA{
			R: 238,
			G: 238,
			B: 238,
		}

		c := &Color{
			Color: rgba,
		}

		assert.Equal(t, "#eeeeee", c.Hex())
	})
}

func TestColor_MarshalJSON(t *testing.T) {
	var spec struct {
		Color Color `json:"color"`
	}

	spec.Color = Color{
		Color: color.RGBA{
			R: 238,
			G: 238,
			B: 238,
		},
	}

	specContent, err := json.Marshal(spec)
	require.NoError(t, err)

	assert.Equal(t, "{\"color\":\"#eeeeee\"}", string(specContent))
}

func TestColor_MarshalYAML(t *testing.T) {
	var spec struct {
		Color Color `yaml:"color"`
	}

	spec.Color = Color{
		Color: color.RGBA{
			R: 238,
			G: 238,
			B: 238,
		},
	}

	specContent, err := yaml.Marshal(spec)
	require.NoError(t, err)

	assert.Equal(t, "color: '#eeeeee'\n", string(specContent))
}

func TestHexToRGBA(t *testing.T) {
	t.Run("#eee", func(t *testing.T) {
		_, err := hexToRGBA("#eee")
		require.NoError(t, err)
	})

	t.Run("#eeeeee", func(t *testing.T) {
		_, err := hexToRGBA("#eeeeee")
		require.NoError(t, err)
	})

	t.Run("#eee = #eeeeee", func(t *testing.T) {
		first, err := hexToRGBA("#eee")
		require.NoError(t, err)

		second, err := hexToRGBA("#eeeeee")
		require.NoError(t, err)

		assert.Equal(t, first, second)
	})
}
