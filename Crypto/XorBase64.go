package Crypto

import (
	"encoding/base64"
)

func XORBase64Encrypt(plainText, key []byte) (string, error) {
	xorEncode, err := XOREncodeDecode(plainText, key)
	if err != nil {
		return "", nil
	}

	base64Encode := base64.StdEncoding.EncodeToString(xorEncode)

	return base64Encode, err
}

func XORBase64Decrypt(cipherText string, key []byte) ([]byte, error) {

	base64Decode, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	xorDecode, err := XOREncodeDecode(base64Decode, key)
	if err != nil {
		return nil, err
	}

	return xorDecode, nil
}
