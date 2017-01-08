package fs

import (
	"os"
	"path/filepath"
	"strings"
)

func GetDirectoryFilesBySuffix(dir string, suffix string) ([]string, error) {
	stat, err := os.Stat(dir)

	if err != nil && os.IsNotExist(err) {
		return nil, err
	}

	files := []string{}
	if !stat.IsDir() && strings.HasSuffix(dir, suffix) {
		files = append(files, dir)
		return files, nil
	}

	if err := filepath.Walk(dir, func(src string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			files = append(files, src)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}
