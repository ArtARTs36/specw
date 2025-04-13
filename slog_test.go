package specw

import (
	"encoding/json"
	"log/slog"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlog_UnmarshallYAML(t *testing.T) {
	type specDefinition struct {
		LogLevel SlogLevel `yaml:"log_level"`
	}

	tests := []struct {
		Title    string
		Content  string
		Expected slog.Level
	}{
		{
			Title:    "debug",
			Content:  "{log_level: debug}",
			Expected: slog.LevelDebug,
		},
		{
			Title:    "info",
			Content:  "{log_level: info}",
			Expected: slog.LevelInfo,
		},
		{
			Title:    "warn",
			Content:  "{log_level: warn}",
			Expected: slog.LevelWarn,
		},
		{
			Title:    "error",
			Content:  "{log_level: error}",
			Expected: slog.LevelError,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			var spec specDefinition

			err := yaml.Unmarshal([]byte(test.Content), &spec)
			require.NoError(t, err)

			assert.Equal(t, test.Expected, spec.LogLevel.Value)
		})
	}
}

func TestSlog_UnmarshalJSON(t *testing.T) {
	type specDefinition struct {
		LogLevel SlogLevel `json:"log_level"`
	}

	tests := []struct {
		Title    string
		Content  string
		Expected slog.Level
	}{
		{
			Title:    "debug",
			Content:  "{\"log_level\": \"debug\"}",
			Expected: slog.LevelDebug,
		},
		{
			Title:    "info",
			Content:  "{\"log_level\": \"info\"}",
			Expected: slog.LevelInfo,
		},
		{
			Title:    "warn",
			Content:  "{\"log_level\": \"warn\"}",
			Expected: slog.LevelWarn,
		},
		{
			Title:    "error",
			Content:  "{\"log_level\": \"error\"}",
			Expected: slog.LevelError,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			var spec specDefinition

			err := json.Unmarshal([]byte(test.Content), &spec)
			require.NoError(t, err)

			assert.Equal(t, test.Expected, spec.LogLevel.Value)
		})
	}
}
