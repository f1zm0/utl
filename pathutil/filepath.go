package pathutil

import (
	"path/filepath"
	"strings"
)

// Filepath is a structure that holds the main attributes of a file path.
type Filepath struct {
	IsDir     bool
	Dirpath   string
	Filepath  string
	Extension string
}

// ParseFilepath parses the provided filepath, and extracts several information about it.
// Returns the dirpath, filepath without extension, and the file extension.
func ParseFilepath(f string) *Filepath {
	f = filepath.Clean(f)
	return &Filepath{
		IsDir:     strings.HasSuffix(f, "/"),
		Dirpath:   filepath.Dir(f),
		Extension: strings.TrimPrefix(filepath.Ext(f), "."),
		Filepath:  strings.TrimSuffix(filepath.Base(f), filepath.Ext(f)),
	}
}
