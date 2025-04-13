package specw

import (
	"fmt"
	"log/slog"
	"strings"

	"gopkg.in/yaml.v3"
)

type SlogLevel struct {
	Level slog.Level
}

func (l *SlogLevel) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected string, got %q", n.Kind)
	}

	val, ok := l.mapValue(n.Value)
	if !ok {
		return fmt.Errorf(
			"unexpected value: %q, possible values: [%s, %s, %s, %s]",
			n.Value,
			slog.LevelDebug,
			slog.LevelInfo,
			slog.LevelWarn,
			slog.LevelError,
		)
	}

	l.Level = val

	return nil
}

func (l *SlogLevel) mapValue(value string) (slog.Level, bool) {
	switch strings.ToLower(value) {
	case "error":
		return slog.LevelError, true
	case "warn":
		return slog.LevelWarn, true
	case "info":
		return slog.LevelInfo, true
	case "debug":
		return slog.LevelDebug, true
	default:
		return 0, false
	}
}
