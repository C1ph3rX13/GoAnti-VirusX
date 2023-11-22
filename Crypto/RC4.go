package Crypto

import "crypto/rc4"

func RC4encrypt(rawData, key []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(rawData))
	c.XORKeyStream(ciphertext, rawData)
	return ciphertext, err
}

func RC4decrypt(ciphertext, key []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	rawData := make([]byte, len(ciphertext))
	c.XORKeyStream(rawData, ciphertext)
	return rawData, err
}
