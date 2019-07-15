package main

import (
	"crypto/aes"
	"fmt"
	"strings"
	"testing"
)

func TestCh61(t *testing.T) {
	if 37 != hamming([]byte("this is a test"), []byte("wokka wokka!!!")) {
		t.Errorf("Hamming was wrong")
	}

	h := hamming([]byte{4, 3}, []byte{69, 17})

	if 4 != h {
		t.Errorf("Hamming was wrong: %d", h)
	}
}

func TestCh6(t *testing.T) {
	input := getSet1Ch6Text()

	keyLenghts := guessKeyLenth(40, input).Sorted()[:2]

	for _, kv := range keyLenghts {

		fmt.Println("Keylen guessed:", kv.k)
		key, plain := breakVarChar(kv.k, getSet1Ch6Text())

		fmt.Println("key:", key)
		fmt.Println("text:", string(plain[0:100]))
	}

}

func TestCh7(t *testing.T) {
	cipher, _ := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	buffer := make([]byte, 16)
	var plaintext []byte
	input := getSet1Ch7Text()
	for len(input) > 0 {
		cipher.Decrypt(buffer, input)
		plaintext = append(plaintext, buffer...)
		input = input[16:]
	}

	if ! strings.Contains(string(plaintext), "Well that's my DJ Deshay cuttin' all them Z's") {
		t.Errorf("Decryption did not work: %s", string(plaintext))
	}
}

func TestCh8(t *testing.T) {
	//texts := getSet1Ch8Texts()
	//
	//for _, text := range texts {
	//	//keyLength := guessKeyLenth(40, text).Sorted()[0]
	//	//fmt.Println("keylen", keyLength.k)
	//	key, plain := breakVarChar(16, text)
	//
	//	fmt.Println("key:", key)
	//	fmt.Println("text:", string(plain[0:20]))
	//}

}
