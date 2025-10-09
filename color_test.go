package specw

import (
	"encoding/json"
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
