package utils

import (
	"path/filepath"
	"path"
	"os"
)

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func RunPath() string {
	exe, _ := filepath.Abs(os.Args[0])
	return path.Dir(exe)
}

func RunFile() string {
	exe, _ := filepath.Abs(os.Args[0])
	return exe
}

func ChangeFileExt(file string, newExt string) string {
	ext := path.Ext(file)
	return file[0:len(file)-len(ext)] + newExt
}
