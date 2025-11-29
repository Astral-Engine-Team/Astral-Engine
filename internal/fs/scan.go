package fs

import (
	"os"
)

func ListItemFiles() ([]os.DirEntry, error) {
	return os.ReadDir("content/items")
}
