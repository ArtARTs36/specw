package specw

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestExistsFilepath_UnmarshalYAML(t *testing.T) {
	t.Run("success when file exists", func(t *testing.T) {
		tmpDir := t.TempDir()
		filePath := filepath.Join(tmpDir, "config.yaml")

		err := os.WriteFile(filePath, []byte("content"), 0o600)
		require.NoError(t, err)

		var spec struct {
			Path ExistsFilepath `yaml:"path"`
		}

		err = yaml.Unmarshal([]byte(fmt.Sprintf("path: %q", filePath)), &spec)
		require.NoError(t, err)
		assert.Equal(t, filePath, spec.Path.Value)
	})

	t.Run("error when file does not exist", func(t *testing.T) {
		var spec struct {
			Path ExistsFilepath `yaml:"path"`
		}

		err := yaml.Unmarshal([]byte("path: /tmp/specw_missing_file"), &spec)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "file /tmp/specw_missing_file does not exist")
	})

	t.Run("error when yaml value cannot decode to string", func(t *testing.T) {
		var spec struct {
			Path ExistsFilepath `yaml:"path"`
		}

		err := yaml.Unmarshal([]byte("path: {nested: true}"), &spec)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "decode filepath:")
	})
}
