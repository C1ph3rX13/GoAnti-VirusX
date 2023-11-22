package compileutils

import (
	"os"
	"path"
	"regexp"
)

const (
	binFile = "payload.bin"
	cFile   = "payload.c"
)

// BinReader 暂时无法读取Stageless类型的bin文件, 数组过大会导致加密过程卡住
func BinReader(binPath string) ([]byte, error) {
	binData, err := os.ReadFile(path.Join(binPath, binFile))
	if err != nil {
		return nil, err
	}

	return binData, nil
}

func CReader(cPath string) ([]byte, error) {
	content, err := os.ReadFile(path.Join(cPath, cFile))
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`"(.*?)"`)
	matches := re.FindAllSubmatch(content, -1)

	var result []byte
	for _, match := range matches {
		result = append(result, match[1]...)
	}

	return result, nil
}
