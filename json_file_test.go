package specw

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestJSONFile_UnmarshalYAML(t *testing.T) {
	type embedded struct {
		Name string `json:"name"`
	}

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "config.json")
	expectedContent := []byte(`{"name":"artem"}`)

	err := os.WriteFile(filePath, expectedContent, 0o600)
	require.NoError(t, err)

	type specDefinition struct {
		File JSONFile[embedded] `yaml:"file"`
	}

	content := "file: " + filePath

	var spec specDefinition

	err = yaml.Unmarshal([]byte(content), &spec)
	require.NoError(t, err)
	assert.Equal(t, "artem", spec.File.Value.Name)
}

func TestJSONFile_UnmarshalJSON(t *testing.T) {
	type embedded struct {
		Name string `json:"name"`
	}

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "config.json")
	expectedContent := []byte(`{"name":"artem"}`)

	err := os.WriteFile(filePath, expectedContent, 0o600)
	require.NoError(t, err)

	type specDefinition struct {
		File JSONFile[embedded] `json:"file"`
	}

	content := `{"file":"` + filePath + `"}`

	var spec specDefinition

	err = json.Unmarshal([]byte(content), &spec)
	require.NoError(t, err)
	assert.Equal(t, "artem", spec.File.Value.Name)
}

func TestJSONFile_UnmarshalBinary(t *testing.T) {
	type embedded struct {
		Name string `json:"name"`
	}

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "config.json")
	expectedContent := []byte(`{"name":"artem"}`)

	err := os.WriteFile(filePath, expectedContent, 0o600)
	require.NoError(t, err)

	var file JSONFile[embedded]

	err = file.UnmarshalBinary([]byte(filePath))
	require.NoError(t, err)
	assert.Equal(t, filePath, file.Path)
	assert.Equal(t, "artem", file.Value.Name)
}
