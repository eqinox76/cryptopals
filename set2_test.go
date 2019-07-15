package main

import (
	"crypto/aes"
	"reflect"
	"strings"
	"testing"
)

func TestCh9(t *testing.T) {
	padded := pkscPad([]byte("YELLOW SUBMARINE"), 20)
	if ! reflect.DeepEqual(padded, []byte("YELLOW SUBMARINE")) {
		t.Errorf("padding '%s'", padded)
	}
}

func TestCh10(t *testing.T) {
	key, _ := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	buffer := make([]byte, 16)
	var plaintext []byte
	cipher := getCh10Text()

	for len(cipher) > 0 {
		key.Decrypt(buffer, cipher[:16])
		plaintext = append(plaintext, xor(buffer, iv)...)

		iv = cipher[:16]
		cipher = cipher[16:]
	}

	if ! strings.Contains(string(plaintext), "Well that's my DJ Deshay cuttin' all them Z's") {
		t.Errorf("Decryption did not work: %s", string(plaintext))
	}
}

func 10
