package base

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

var Sed string = "3cad4cca4fc93aadedcb12bb883abf48aba63dfafcd57ac88ee77d9cd7d32b5baecb3d4b1b3e0e4beffd"

func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	rkey := sha256.Sum256(key)
	block, err := des.NewTripleDESCipher(rkey[:24])
	if err != nil {
		return nil, err
	}
	hex, err := hex.DecodeString(Sed)
	if err != nil {
		return nil, err
	}

	if len(hex) < block.BlockSize() {
		return nil, errors.New("hex too short")
	}

	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, hex[:block.BlockSize()])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	rkey := sha256.Sum256(key)
	block, err := des.NewTripleDESCipher(rkey[:24])
	if err != nil {
		return nil, err
	}
	hex, err := hex.DecodeString(Sed)
	if err != nil {
		return nil, errors.New("hex too short")
	}

	blockMode := cipher.NewCBCDecrypter(block, hex[:block.BlockSize()])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key[:8])
	if err != nil {
		return nil, err
	}

	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key[:8])
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
