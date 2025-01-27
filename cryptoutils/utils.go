package cryptoutils

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
)

func HexToByte(inp string) []byte {
	rawInp := []byte(inp)
	out := make([]byte, hex.DecodedLen(len(rawInp)))
	n, err := hex.Decode(out, rawInp)
	if err != nil {
		panic(err)
	}
	return out[:n]
}

func StringToByte(inp string) []byte {
	return []byte(inp)
}

func Base64ToByte(inp string) []byte {
	dec, err := base64.StdEncoding.DecodeString(inp)
	if err != nil {
		panic(err)
	}
	return dec
}

// string is just for pprint
func ByteToBase64(inp []byte) string {
	enc := base64.StdEncoding.EncodeToString(inp)
	return enc
}

func ByteToHex(inp []byte) string {
	out := hex.EncodeToString(inp)
	return out
}

func ByteToString(inp []byte) string {
	return string(inp[:])
}

func ByteXor(inp []byte, key []byte) []byte {
	if len(inp) < len(key) {
		panic("Input lengths don't match!")
	}

	out := make([]byte, len(inp))
	for i := 0; i < len(inp); i++ {
		out[i] = inp[i] ^ key[i%len(key)]
	}
	return out
}

func EvalPlainText(inpRaw []byte) float64 {
	capsHeuristic := 0.05
	spaceHeuristic := 0.18
	charHeuristic := 0.78

	capsCount, spaceCount, charCount, deduction := 0, 0, 0, 0

	for _, v := range inpRaw {
		if v >= 65 && v <= 90 {
			capsCount = capsCount + 1
		} else if v >= 97 && v <= 122 {
			charCount = charCount + 1
		} else if v == 32 {
			spaceCount = spaceCount + 1
		} else {
			deduction = deduction + 1
		}
	}

	score := math.Abs(float64(capsCount)/float64(len(inpRaw)) - capsHeuristic)
	score = score + math.Abs(float64(charCount)/float64(len(inpRaw))-charHeuristic)
	score = score + math.Abs(float64(spaceCount)/float64(len(inpRaw))-spaceHeuristic)
	score = score + float64(deduction)/50

	return score
}

func FindBestSingleXor(inp []byte) (float64, []byte) {
	var minScore float64 = 100
	var likelyDecode []byte
	for i := byte(0); i < 255; i++ {
		out := ByteXor(inp, []byte{i})
		if EvalPlainText(out) < minScore {
			minScore = EvalPlainText(out)
			likelyDecode = out
		}
	}
	return minScore, likelyDecode
}

func HammingDistance(inp1 []byte, inp2 []byte) int16 {
	var dist int16 = 0

	for i := range len(inp1) {
		for mask := byte(128); mask != 0; mask >>= 1 {
			if (inp1[i] & mask) == (inp2[i] & mask) {
				// do nothing
			} else {
				dist = dist + 1
			}
		}
	}
	return dist
}

func DecryptAesEcb(inp []byte, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	decrypted := make([]byte, len(inp))
	size := 16

	for bs, be := 0, size; bs < len(inp); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], inp[bs:be])
	}

	return decrypted
}
