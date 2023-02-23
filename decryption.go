package localCrypto

import (
	"errors"
	"math/rand"
)

func getRowColsForDecryption(length int, key int64) (int, int) {
	rand.Seed(key)

	rows := rand.Intn((length / 2)) + 2
	cols := (length / rows)

	if cols*rows < length {
		cols += 1
	}

	return rows, cols
}

func (d CipherText) Decrypt() (string, error) {
	length := len(d.Entry)

	if length == 0 {
		return "", errors.New("EmptyEntryFound")
	}
	rows, cols := getRowColsForDecryption(length, d.key)
	decodeIt := []byte(d.Entry)
	matrix := make([][]byte, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]byte, cols)
	}

	k := 0
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if k >= length {
				break
			}
			matrix[j][i] = decodeIt[k] ^ 1
			k++
		}
	}

	simpleText := []byte{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] != '\x00' {
				simpleText = append(simpleText, matrix[i][j])
			}
		}
	}
	return string(simpleText)
}

func (c CipherText) PrintKey() int64 {
	return c.key
}
