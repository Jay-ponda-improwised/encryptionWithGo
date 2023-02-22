package localCrypto

import (
	"math/rand"
	"time"
)

func genKey() int64 {
	return time.Now().UnixMilli()
}

func getRowColsForEncryption(length int, key int64) (int, int) {
	rand.Seed(key)

	cols := rand.Intn((length / 2)) + 2
	rows := (length / cols)

	if cols*rows < length {
		rows += 1
	}

	return rows, cols
}

func NewEncryptedText(entry string) *CipherText {
	key := genKey()
	return &CipherText{
		Entry: encrypteIt(entry, key),
		key:   key,
	}
}

func encrypteIt(entry string, key int64) string {
	l := len(entry)
	rows, cols := getRowColsForEncryption(l, key)
	matrix := make([][]byte, rows)
	crypt := make([]byte, l)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]byte, cols)
	}

	k := 0
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if k >= l {
				break
			}
			matrix[j][i] = entry[k] ^ 1
			k++
		}
	}

	k = 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if k >= l {
				break
			}
			crypt[k] = matrix[i][j]
			k++
		}
	}

	return string(crypt)
}
