package cryptoutils

import (
	"encoding/base64"
	"encoding/hex"
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
	charHeuristic := 0.72

	capsCount, spaceCount, charCount := 0, 0, 0

	for _, v := range inpRaw {
		if v >= 65 && v <= 90 {
			capsCount = capsCount + 1
		} else if v >= 97 && v <= 122 {
			charCount = charCount + 1
		} else if v == 32 {
			spaceCount = spaceCount + 1
		}
	}

	score := math.Abs(float64(capsCount)/float64(len(inpRaw)) - capsHeuristic)
	score = score + math.Abs(float64(charCount)/float64(len(inpRaw))-charHeuristic)
	score = score + math.Abs(float64(spaceCount)/float64(len(inpRaw))-spaceHeuristic)

	return score
}
