package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileStorage_ReadSuccessful(t *testing.T) {
	tmpDir := t.TempDir()
	fileIn := filepath.Join(tmpDir, "input.txt")
	content := "test content"

	err := os.WriteFile(fileIn, []byte(content), 0644)
	require.NoError(t, err)

	storage := New(fileIn, "output.txt")

	result, err := storage.Read()

	require.NoError(t, err)
	require.Equal(t, content, result)
}

func TestFileStorage_ReadErrorFileNotFound(t *testing.T) {
	storage := New("nonexistent.txt", "output.txt")

	_, err := storage.Read()

	require.Error(t, err)
	require.Contains(t, err.Error(), "no such file")
}

func TestFileStorage_WriteSuccessful(t *testing.T) {
	tmpDir := t.TempDir()
	fileOut := filepath.Join(tmpDir, "output.txt")

	storage := New("input.txt", fileOut)
	content := "test content"

	err := storage.Write(content)

	require.NoError(t, err)

	data, err := os.ReadFile(fileOut)
	require.NoError(t, err)
	require.Equal(t, content, string(data))
}

func TestFileStorage_WriteErrorReadOnly(t *testing.T) {
	storage := New("input.txt", "/root/readonly/output.txt")

	err := storage.Write("content")

	require.Error(t, err)
}
