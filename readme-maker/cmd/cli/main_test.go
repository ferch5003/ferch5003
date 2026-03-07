package main

import (
	"os"
	"path/filepath"
	"testing"

	nasatest "github.com/ferch5003/ferch5003/readme-maker/internal/nasa/nasatest"
	"github.com/stretchr/testify/require"
)

func TestRunSuccessful(t *testing.T) {
	// Given
	templateContent := `# Test README
Title: {{.Nasa.APOD.Title}}
Date: {{.Nasa.APOD.Date}}
`
	readmeContent := `# My README
`
	tempDir := t.TempDir()
	templateFile := filepath.Join(tempDir, "README.md.tpl")
	readmeFile := filepath.Join(tempDir, "README.md")

	err := os.WriteFile(templateFile, []byte(templateContent), 0644)
	require.NoError(t, err)

	err = os.WriteFile(readmeFile, []byte(readmeContent), 0644)
	require.NoError(t, err)

	// Create mock NASA server
	server := nasatest.NewServer()
	defer server.Close()
	t.Setenv("NASA_BASE_URL", server.URL)
	t.Setenv("NASA_API_KEY", "test")

	// Override the dependencies to use temp files
	origTemplateFile := templatePath
	origReadmeFile := readmePath
	templatePath = templateFile
	readmePath = readmeFile

	// When
	err = run()

	// Then
	require.NoError(t, err)

	// Restore
	templatePath = origTemplateFile
	readmePath = origReadmeFile
}

func TestRunReadError(t *testing.T) {
	// Given
	tempDir := t.TempDir()
	// Create readme file but no template file
	readmeFile := filepath.Join(tempDir, "README.md")
	err := os.WriteFile(readmeFile, []byte("# Test"), 0644)
	require.NoError(t, err)

	// Create mock NASA server (needed for dependencies)
	server := nasatest.NewServer()
	defer server.Close()
	t.Setenv("NASA_BASE_URL", server.URL)
	t.Setenv("NASA_API_KEY", "test")

	// Override paths - template file doesn't exist (but readme does)
	origTemplateFile := templatePath
	origReadmeFile := readmePath
	templatePath = filepath.Join(tempDir, "nonexistent.tpl")
	readmePath = readmeFile

	// When
	err = run()

	// Then
	require.Error(t, err)

	// Restore
	templatePath = origTemplateFile
	readmePath = origReadmeFile
}

func TestRunTemplateParseError(t *testing.T) {
	// Given - invalid template (missing closing brace)
	templateContent := `# Test README
{{.Nasa.APOD.Title}
`
	readmeContent := `# My README
`
	tempDir := t.TempDir()
	templateFile := filepath.Join(tempDir, "README.md.tpl")
	readmeFile := filepath.Join(tempDir, "README.md")

	err := os.WriteFile(templateFile, []byte(templateContent), 0644)
	require.NoError(t, err)

	err = os.WriteFile(readmeFile, []byte(readmeContent), 0644)
	require.NoError(t, err)

	// Create mock NASA server
	server := nasatest.NewServer()
	defer server.Close()
	t.Setenv("NASA_BASE_URL", server.URL)
	t.Setenv("NASA_API_KEY", "test")

	// Override the dependencies to use temp files
	origTemplateFile := templatePath
	origReadmeFile := readmePath
	templatePath = templateFile
	readmePath = readmeFile

	// When
	err = run()

	// Then
	require.Error(t, err)

	// Restore
	templatePath = origTemplateFile
	readmePath = origReadmeFile
}

func TestRunWriteError(t *testing.T) {
	// Given - readme file is a directory (can't write to it)
	templateContent := `# Test README
Title: {{.Nasa.APOD.Title}}
`
	readmeContent := `# My README
`
	tempDir := t.TempDir()
	templateFile := filepath.Join(tempDir, "README.md.tpl")
	readmeFile := filepath.Join(tempDir, "README.md")

	err := os.WriteFile(templateFile, []byte(templateContent), 0644)
	require.NoError(t, err)

	err = os.WriteFile(readmeFile, []byte(readmeContent), 0644)
	require.NoError(t, err)
	// Make readmeFile a directory to cause write error
	_ = os.Remove(readmeFile)
	err = os.Mkdir(readmeFile, 0755)
	require.NoError(t, err)

	// Create mock NASA server
	server := nasatest.NewServer()
	defer server.Close()
	t.Setenv("NASA_BASE_URL", server.URL)
	t.Setenv("NASA_API_KEY", "test")

	// Override the dependencies to use temp files
	origTemplateFile := templatePath
	origReadmeFile := readmePath
	templatePath = templateFile
	readmePath = readmeFile

	// When
	err = run()

	// Then
	require.Error(t, err)

	// Restore
	templatePath = origTemplateFile
	readmePath = origReadmeFile
}
