// Copyright 2021 utils. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package license

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"os"
)

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

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

const (
	aeskey = "nlasfl2wrwfsnsfs131#$%fs"
)

// Create 写入文件
func Create(fileName string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	aestr, err := AesEncrypt(data, []byte(aeskey))
	if err != nil {
		return err
	}
	os.Remove(fileName)
	fp, err := os.Create(fileName)
	if err != nil {
		return err
	}
	fp.Write(aestr)
	defer fp.Close()
	return nil
}

// LicenseRead 读
func Read(filename string, v interface{}) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Create(filename, v)
	}
	//获取文件内容
	aesdata, err := AesDecrypt(data, []byte(aeskey))
	if err != nil {
		return err
	}
	return json.Unmarshal(aesdata, v)
}
