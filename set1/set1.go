package set1

import (
	"bufio"
	"os"

	"github.com/Jonnowang/crypto-algos/cryptoutils"
)

func Ch1(inp string) string {
	hexBytes := cryptoutils.HexToByte(inp)
	charBase64 := cryptoutils.ByteToBase64(hexBytes)
	return charBase64
}

func Ch2(inp1 string, inp2 string) string {
	hexBytes1 := cryptoutils.HexToByte(inp1)
	hexBytes2 := cryptoutils.HexToByte(inp2)

	hexOut := cryptoutils.ByteXor(hexBytes1, hexBytes2)
	return cryptoutils.ByteToHex(hexOut)
}

func Ch3(inp string) (float64, string) {
	inpRaw := cryptoutils.HexToByte(inp)

	var minScore float64 = 100
	var likelyDecode []byte
	for i := byte(0); i < 255; i++ {
		out := cryptoutils.ByteXor(inpRaw, []byte{i})
		if cryptoutils.EvalPlainText(out) < minScore {
			minScore = cryptoutils.EvalPlainText(out)
			likelyDecode = out
		}
	}
	return minScore, cryptoutils.ByteToString(likelyDecode)
}

func Ch4(filepath string) (float64, string) {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	var minScore float64 = 100
	var likelyDecode string

	for scanner.Scan() {
		score, text := Ch3(scanner.Text())
		if score < minScore {
			minScore = score
			likelyDecode = text
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return minScore, likelyDecode
}

func Ch5(inp string, key string) string {
	inpRaw := cryptoutils.StringToByte(inp)
	keyRaw := cryptoutils.StringToByte(key)

	out := cryptoutils.ByteXor(inpRaw, keyRaw)
	return cryptoutils.ByteToHex(out)
}
