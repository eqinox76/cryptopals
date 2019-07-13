package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//fmt.Println("1.1")
	//set1ch1()
	//fmt.Println("1.2")
	//set1ch2()
	//fmt.Println("1.3")
	//set1ch3()
	//fmt.Println("1.5")
	set1ch5()
	fmt.Println("1.6")

}

func set1ch1() {
	src := decodeHexStr("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	fmt.Println(src)
	str := base64.StdEncoding.EncodeToString(src)
	fmt.Println(str)
}

func set1ch2() {

	a := decodeHexStr("1c0111001f010100061a024b53535009181c")
	b := decodeHexStr("686974207468652062756c6c277320657965")

	printStr(xor(a, b))
}

func set1ch3() {
	input := decodeHexStr("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	key, text := breakOneChar(input)
	fmt.Println(key, string(text))
}

func set1ch5() {
	input := []byte(`The Guardian is a British daily newspaper. It was founded in 1821 as The Manchester Guardian, and changed its name in 1959. Along with its sister papers The Observer and The Guardian Weekly, the Guardian is part of the Guardian Media Group, owned by the Scott Trust. The trust was created in 1936 to "secure the financial and editorial independence of the Guardian in perpetuity and to safeguard the journalistic freedom and liberal values of the Guardian free from commercial or political interference".[4] The trust was converted into a limited company in 2008, with a constitution written so as to maintain for The Guardian the same protections as were built into the structure of the Scott Trust by its creators. Profits are reinvested in journalism rather than distributed to owners or shareholders.[5]The current editor is Katharine Viner: she succeeded Alan Rusbridger in 2015.[6][7] Since 2018, the paper's main newsprint sections have been published in tabloid format. As of November that year, its print edition had a daily circulation of 136,834.[8] The newspaper has an online edition, TheGuardian.com, as well as two international websites, Guardian Australia (founded in 2013) and Guardian US (founded in 2011). The paper's readership is generally on the mainstream left of British political opinion,[9][10] and its reputation as a platform for liberal and left-wing editorial (despite the high proportion of privately educated journalists writing for it) has led to the use of the "Guardian reader" and "Guardianista" as often-pejorative epithets for those of left-leaning or "politically correct" tendencies.[11][12][13] Frequent typographical errors in the paper led Private Eye magazine to dub it the "Grauniad" in the 1960s, a nickname still used today.[14]In an Ipsos MORI research poll in September 2018 designed to interrogate the public's trust of specific titles online, The Guardian scored highest for digital-content news, with 84% of readers agreeing that they "trust what [they] see in it".[15] A December 2018 report of a poll by the Publishers Audience Measurement Company (PAMCo) stated that the paper's print edition was found to be the most trusted in the UK in the period from October 2017 to September 2018. It was also reported to be the most-read of the UK's "quality newsbrands", including digital editions; other "quality" brands included The Times, The Daily Telegraph, The Independent, and the i. While The Guardian's print circulation is in decline, the report indicated that news from The Guardian, including that reported online, reaches more than 23 million UK adults each month.[16]Chief among the notable "scoops" obtained by the paper was the 2011 News International phone-hacking scandal—and in particular the hacking of the murdered English teenager Milly Dowler's phone.[17] The investigation led to the closure of the News of the World, the UK's best-selling Sunday newspaper and one of the highest-circulation newspapers in history.[18] In June 2013, The Guardian broke news of the secret collection by the Obama administration of Verizon telephone records,[19] and subsequently revealed the existence of the surveillance program PRISM after knowledge of it was leaked to the paper by the whistleblower and former NSA contractor Edward Snowden.[20] In 2016, The Guardian led an investigation into the Panama Papers, exposing then-Prime Minister David Cameron's links to offshore bank accounts. It has been named "newspaper of the year" four times at the annual British Press Awards: most recently in 2014, for its reporting on government surveillance.[21] `)

	var key = []byte("SecrET")
	encode := func(code []byte, text []byte) []byte {
		dst := make([]byte, len(text))

		for i := range text {
			dst[i] = code[i%len(code)] ^ text[i]
		}
		return dst
	}

	printStr(encode(key, input))
}
