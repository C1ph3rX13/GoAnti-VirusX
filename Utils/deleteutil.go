package Utils

import (
	"os"
	"path/filepath"
)

func DeleteFiles(pattern string) error {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			return err
		}
		logger.Infof("Deleted files: %s", file)
	}

	return nil
}
