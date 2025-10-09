package specw

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
)

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
