package compileutils

import (
	"fmt"
	"os"
)

func WriteStringsToFile(content, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", err)
	}

	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write content to file: %s", err)
	}

	return nil
}
