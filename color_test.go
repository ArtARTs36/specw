package specw

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestColor_UnmarshalJSON(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		var cfg struct {
			Color Color `json:"color"`
		}

		err := json.Unmarshal([]byte(`{"color":"red"}`), &cfg)
		require.NoError(t, err)
		require.Equal(t, "red", cfg.Color.Raw)
	})
}
