package Crypto

import "errors"

func XOREncodeDecode(plainText, key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("key cannot be empty")
	}

	encodedData := make([]byte, len(plainText))
	keyLen := len(key)

	for i := 0; i < len(plainText); i++ {
		encodedData[i] = plainText[i] ^ key[i%keyLen]
	}

	return encodedData, nil
}
