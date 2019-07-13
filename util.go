package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/bits"
	"strings"
)

func decodeHex(src []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	return dst
}

func decodeHexStr(src string) []byte {
	return decodeHex([]byte(src))
}

func printStr(data []byte) {
	fmt.Println(hex.EncodeToString(data))
}

func xor(a []byte, b []byte) []byte {
	dst := make([]byte, len(a))

	for i := range a {
		dst[i] = a[i] ^ b[i]
	}

	return dst
}

func hamming(a, b []byte) int {
	if len(a) != len(b) {
		log.Fatal("a b are not the same length", len(a), len(b))
	}

	diff := 0
	for i := range a {
		c := a[i] ^ b[i]
		diff += bits.OnesCount8(c)
		//fmt.Println(a[i], "^", b[i] , "=", c , "bits:", bits.OnesCount8(c), strconv.FormatInt(int64(c), 2))
	}
	return diff
}

func guessKeyLenth(max int, data []byte) Histo {

	histo := NewHisto()

	for i := 2; (i <= max) && (i*8 < len(data)); i++ {
		k1 := data[0:i]
		k2 := data[i : 2*i]
		k3 := data[2*i : 3*i]
		k4 := data[3*i : 4*i]
		//fmt.Println(string(k1), "|", string(k2), "|", string(k3), "|" ,string(k4))
		localScore := 0
		localScore += hamming(k1, k2)
		localScore += hamming(k1, k3)
		localScore += hamming(k1, k4)
		localScore += hamming(k2, k1)
		localScore += hamming(k2, k3)
		localScore += hamming(k2, k4)
		localScore += hamming(k3, k1)
		localScore += hamming(k3, k2)
		localScore += hamming(k3, k4)
		localScore += hamming(k4, k1)
		localScore += hamming(k4, k2)
		localScore += hamming(k4, k3)

		s := float64(localScore) / float64(12 * i)
		histo.Values[i] = s
	}

	return histo
}

func breakOneChar(input []byte) (byte, []byte) {
	distribution := "!-',.ULDRHSNIOATEuldrhs nioate"

	decode := func(code byte, cipher []byte) []byte {
		dst := make([]byte, len(cipher))

		for i := range cipher {
			dst[i] = code ^ cipher[i]
		}
		return dst
	}

	score := func(input []byte) int {


		sum := 0
		for _, i := range input {
			sum += strings.Index(distribution, string(i))
		}
		return sum
	}

	max := 0
	var result []byte
	var key byte

	for c :=0 ; c < 255; c++ {
		text := decode(byte(c), input)
		score := score(text)
		//fmt.Println(c, score, string(text[:100]))
		if max < score {
			max = score
			result = text
			key = byte(c)
			//fmt.Println(">>>", score)
		}
	}

	//fmt.Println(key, "===============================================")

	return key, result
}
