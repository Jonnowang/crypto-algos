package main

import (
	"log"

	"github.com/Jonnowang/crypto-algos/set1"
)

func main() {
	log.Println("Challenge 1")
	log.Print(set1.Ch1("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))

	log.Println("Challenge 2")
	log.Print(set1.Ch2("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))

	log.Println("Challenge 3")
	_, likely_decode := set1.Ch3("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	log.Print(likely_decode)

	log.Println("Challenge 4")
	_, likely_decode = set1.Ch4("./set1/data/ch4.txt")
	log.Print(likely_decode)

	log.Println("Challenge 5")
	log.Print(set1.Ch5("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", "ICE"))
}
