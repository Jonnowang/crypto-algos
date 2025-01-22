package set1

import (
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

func Ch3(inp string) string {
	inpRaw := cryptoutils.HexToByte(inp)

	var minScore float64 = 100
	var likelyDecode []byte
	for i := byte(65); i < 122; i++ {
		out := cryptoutils.ByteXor(inpRaw, []byte{i})
		if cryptoutils.EvalPlainText(out) < minScore {
			minScore = cryptoutils.EvalPlainText(out)
			likelyDecode = out
		}
	}
	return cryptoutils.ByteToString(likelyDecode)
}
