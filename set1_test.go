package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCh61(t *testing.T){
	if 37 != hamming([]byte("this is a test"), []byte("wokka wokka!!!")){
       t.Errorf("Hamming was wrong")
	}

	h := hamming([]byte{ 4,3 }, []byte{69, 17})

	if 4 != h{
		t.Errorf("Hamming was wrong: %d", h)
	}
}

func TestCh6(t *testing.T){
	input := getSet1Ch6Text()

	keyLenghts := guessKeyLenth(40, input).Sorted()[:2]

	for _, kv := range keyLenghts {

		fmt.Println("Keylen guessed:", kv.k)

		ciphers := make([]bytes.Buffer, kv.k)
		texts := make([]bytes.Buffer, kv.k)

		for i, v := range getSet1Ch6Text() {
			ciphers[i%kv.k].WriteByte(v)
		}

		var keyBuilder strings.Builder

		for i, buffer := range ciphers {
			key, text := breakOneChar(buffer.Bytes())
			texts[i].Write(text)
			keyBuilder.WriteByte(key)
		}

		fmt.Println("key:", keyBuilder.String())

		result := bytes.Buffer{}
		for i := range getSet1Ch6Text() {
			b, _ := texts[i%kv.k].ReadByte()

			result.WriteByte(b)
		}

		fmt.Println("text:", string(result.Bytes()[0:100]))
	}

}
