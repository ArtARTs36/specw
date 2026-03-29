package specw

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestYAMLFile_UnmarshalYAML(t *testing.T) {
	type embedded struct {
		Name string
	}

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "config.yaml")
	expectedContent := []byte("name: artem")

	err := os.WriteFile(filePath, expectedContent, 0o600)
	require.NoError(t, err)

	type specDefinition struct {
		File YAMLFile[embedded] `yaml:"file"`
	}

	content := "file: " + filePath

	var spec specDefinition

	err = yaml.Unmarshal([]byte(content), &spec)
	require.NoError(t, err)
	assert.Equal(t, "artem", spec.File.Value.Name)
}
