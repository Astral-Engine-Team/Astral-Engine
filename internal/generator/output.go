package generator

import (
	"os"
	"path/filepath"
)

func WriteFile(path string, data string) error {
	dir := filepath.Dir(path)
	os.MkdirAll(dir, os.ModePerm)
	return os.WriteFile(path, []byte(data), 0644)
}
