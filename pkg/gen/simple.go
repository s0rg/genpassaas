package gen

import "bytes"

const simpleAlphabetBlocks = 3

var (
	alphaLower = []byte("abcdefghijklmnopqrstuvwxyz")
	alphaUpper = bytes.ToUpper(alphaLower)
	extra      = [][]byte{
		[]byte("0123456789"),
		[]byte("!@#$%&*-+()[]{}=?"),
	}
)

type bytesSwapper []byte

func (c bytesSwapper) Len() int      { return len(c) }
func (c bytesSwapper) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

func Simple(length int) (rv string) {
	alphabetSize := len(alphaLower) * simpleAlphabetBlocks

	alphabet := make([]byte, 0, alphabetSize)

	alphabet = append(alphabet, alphaLower...)
	alphabet = append(alphabet, alphaUpper...)

	for _, ext := range extra {
		alphabet = append(alphabet, ext...)
	}

	shuffle(bytesSwapper(alphabet))

	var (
		max  = len(alphabet)
		pass = make([]byte, length)
	)

	for i := 0; i < length; i++ {
		pass[i] = alphabet[mustRandInt(max)]
	}

	shuffle(bytesSwapper(pass))

	return string(pass)
}
