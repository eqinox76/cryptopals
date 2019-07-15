package main

import (
	"crypto/aes"
	"math/rand"
)

func pkscPad(input []byte, padTo int) []byte{
	pad := padTo - len(input)
	padTo -= len(input)
	result := input
	for i := 0; i < padTo; i++{
		result = append(result, byte(pad))
	}

	return result
}

func ECBEncrypt(input []byte, key string) []byte{
	keyCipher, _ := aes.NewCipher([]byte(key))
	buffer := make([]byte, 16)
	var cipher []byte
	for len(input) > 0 {
		keyCipher.Encrypt(buffer, input)
		cipher = append(cipher, buffer...)
		input = input[16:]
	}
	return cipher
}

func ECBCBCencrypt(input []byte) []byte{

	key := make([]byte, 16)
	rand.Read(key)

	token1 := make([]byte, rand.Intn(6) + 5)
	rand.Read(token1)
	token2 := make([]byte, rand.Intn(6) + 5)
	rand.Read(token2)

	input = append(token1, input...)
	input = append(input, token2...)

	if rand.Intn(1) == 0 {
		return ECBEncrypt(input, string(key))
	} else {
		return nil
	}
}