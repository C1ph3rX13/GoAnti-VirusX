package compileutils

import (
	"GoShellCodeLoader/Utils"
	"os"
	"path/filepath"
)

func IsWinresExist() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	winresDir := filepath.Join(dir, "winres")
	if _, err = os.Stat(winresDir); os.IsNotExist(err) {
		err = os.Mkdir(winresDir, os.ModePerm)
		if err != nil {
			return err
		}
		Utils.GetLogger().Warningf("Directory created successfully: %s", winresDir)
	} else if err != nil {
		return err
	}

	return nil
}
