package specw

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHexColor_FromHex(t *testing.T) {
	t.Run("#eee", func(t *testing.T) {
		c := HexColor{}

		err := c.FromHex("#eee")
		require.NoError(t, err)
	})

	t.Run("#eeeeee", func(t *testing.T) {
		c := HexColor{}

		err := c.FromHex("#eee")
		require.NoError(t, err)
	})
}
