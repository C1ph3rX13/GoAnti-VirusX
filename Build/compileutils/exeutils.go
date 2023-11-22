package compileutils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExeBuilderCmdStartAndWait(goFileName, exeFileName string) error {
	cmd := exec.Command("go", "build", "-ldflags=-s -w", "-o", exeFileName, "-trimpath", goFileName)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("failed to build executable: %s: %w", stderr.String(), err)
	}

	return nil
}

func ExeBuilderCombinedOutput(goFileName, exeFileName string) error {
	cmd := exec.Command("go", "build", "-ldflags=-s -w", "-o", exeFileName, "-trimpath", goFileName)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to build executable: %s: %w", string(output), err)
	}

	return nil
}
