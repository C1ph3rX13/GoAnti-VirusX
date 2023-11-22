package Crypto

import (
	"github.com/eknkc/basex"
)

var base85 *basex.Encoding

func init() {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~"
	base85, _ = basex.NewEncoding(alphabet)
}

func Base85Encode(plainText []byte) (string, error) {
	// Base85 编码
	encodedData := base85.Encode(plainText)
	return encodedData, nil
}

func Base85Decode(encodedData string) ([]byte, error) {
	// Base85 解码
	decodedData, err := base85.Decode(encodedData)
	if err != nil {
		return nil, err
	}

	return decodedData, nil
}
