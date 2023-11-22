package selectfunctmpl

import (
	"fmt"
	"os"
	"strings"
)

func GetDecryptFuncFromTemplate(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %s", err)
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Decrypt") {
			// 解析出解密函数名
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				continue
			}
			decryptFunc := strings.TrimSpace(parts[1])
			return decryptFunc, nil
		}
	}

	return "", fmt.Errorf("decrypt function not found in the template file")
}

func ReplaceEncryptMethodInFile(filePath, decryptFunc, newEncryptMethod string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read program file: %s", err)
	}

	// 进行替换操作
	replacedContent := strings.ReplaceAll(string(content), decryptFunc, newEncryptMethod)

	return replacedContent, nil
}
