package specw

import (
	"testing"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDuration_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		Title    string
		Content  string
		Expected Duration
	}{
		{
			Title:   "int: 0",
			Content: "{interval: 0}",
			Expected: Duration{
				Value: time.Duration(0),
			},
		},
		{
			Title:   "int: 5",
			Content: "{interval: 5}",
			Expected: Duration{
				Value: time.Duration(5),
			},
		},
		{
			Title:   "string: 5ms",
			Content: "{interval: 5ms}",
			Expected: Duration{
				Value: 5 * time.Millisecond,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			var spec struct {
				Interval Duration `yaml:"interval"`
			}

			err := yaml.Unmarshal([]byte(test.Content), &spec)
			require.NoError(t, err)
			assert.Equal(t, test.Expected, spec.Interval)
		})
	}
}
