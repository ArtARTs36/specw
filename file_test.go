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

func TestFile_UnmarshalYAML(t *testing.T) {
	t.Run("read file content", func(t *testing.T) {
		tmpDir := t.TempDir()
		filePath := filepath.Join(tmpDir, "config.txt")
		expectedContent := []byte("hello from file")

		err := os.WriteFile(filePath, expectedContent, 0o600)
		require.NoError(t, err)

		var spec struct {
			File File `yaml:"file"`
		}

		err = yaml.Unmarshal([]byte(fmt.Sprintf(`file: %q`, filePath)), &spec)
		require.NoError(t, err)
		assert.Equal(t, expectedContent, spec.File.Content)
	})

	t.Run("return error when file does not exist", func(t *testing.T) {
		var spec struct {
			File File `yaml:"file"`
		}

		err := yaml.Unmarshal([]byte(`file: /tmp/specw_missing_file`), &spec)
		require.Error(t, err)
		assert.Contains(t, err.Error(), `read file "/tmp/specw_missing_file":`)
	})
}
