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
	minScore, likelyDecode := cryptoutils.FindBestSingleXor(inpRaw)
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
		inpRaw := cryptoutils.HexToByte(scanner.Text())
		score, text := cryptoutils.FindBestSingleXor(inpRaw)
		if score < minScore {
			minScore = score
			likelyDecode = cryptoutils.ByteToString(text)
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

func Ch6(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	var inpBytes []byte
	for scanner.Scan() {
		inpBytes = append(inpBytes, cryptoutils.Base64ToByte(scanner.Text())...)
	}

	var likelyKeySize int = 0
	var minScore float64 = 100

	for guessKeysize := 2; guessKeysize <= 40; guessKeysize++ {
		// Desired number of keysize chunks
		numKeys := 6

		keys := make([][]byte, numKeys)
		for i := 0; i < numKeys; i++ {
			start := i * guessKeysize
			end := (i + 1) * guessKeysize
			// Ensure we don't go out of bounds
			if end > len(inpBytes) {
				break
			}
			keys[i] = inpBytes[start:end]
		}

		// Compute the avg score for all combinations of keys
		score := 0.0
		for i := 0; i < len(keys); i++ {
			for j := i + 1; j < len(keys); j++ {
				score += float64(cryptoutils.HammingDistance(keys[i], keys[j])) / float64(guessKeysize)
			}
		}

		if score < minScore {
			minScore = score
			likelyKeySize = guessKeysize
		}
	}

	out := make([]byte, len(inpBytes))
	for i := range likelyKeySize {
		var compositeInp []byte
		// Collect transposed chunks into slices
		for j := i; j < len(inpBytes); j += likelyKeySize {
			compositeInp = append(compositeInp, inpBytes[j])
		}
		_, estimate := cryptoutils.FindBestSingleXor(compositeInp)

		// Reconstruct transposed chunks
		idx := 0
		for k := i; k < len(inpBytes); k += likelyKeySize {
			out[k] = estimate[idx]
			idx++
		}
	}

	return cryptoutils.ByteToString(out)
}

func Ch7(filepath string, keyPlain string) string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	var inpBytes []byte
	for scanner.Scan() {
		inpBytes = append(inpBytes, cryptoutils.Base64ToByte(scanner.Text())...)
	}

	key := cryptoutils.StringToByte(keyPlain)
	out := cryptoutils.DecryptAesEcb(inpBytes, key)

	return cryptoutils.ByteToString(out)
}

func Ch8(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	blockSize := 16
	blockSet := make(map[[16]byte]bool)
	for scanner.Scan() {
		lineBytes := cryptoutils.HexToByte(scanner.Text())
		for i := 0; i < len(lineBytes); i += blockSize {
			var block [16]byte
			copy(block[:], lineBytes[i:i+blockSize])
			_, ok := blockSet[block]
			if ok {
				return cryptoutils.ByteToHex(lineBytes)
			}
			blockSet[block] = true
		}
	}

	return "No valid block found"
}
