package compileutils

import (
	"GoShellCodeLoader/Crypto"
	"fmt"
	"path/filepath"
)

func EncWriter(binDir, fileName string, key []byte) (string, error) {
	binRaw, err := BinReader(binDir)
	if err != nil {
		return "", fmt.Errorf("read binary file failed: %w", err)
	}

	encrypt, err := Crypto.XORRC4HexBase85Encrypt(binRaw, key)
	if err != nil {
		return "", fmt.Errorf("encrypt binary failed: %w", err)
	}

	err = WriteStringsToFile(encrypt, filepath.Join(binDir, fileName))
	if err != nil {
		return "", fmt.Errorf("write To file failed: %w", err)
	}

	return encrypt, nil
}
