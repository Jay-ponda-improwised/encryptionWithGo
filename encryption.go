package localCrypto

import (
	"errors"
	"math/rand"
	"time"
)

func GenKey() int64 {
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

func Encrypt(entry string) (CipherText, error) {
	if entry == "" {
		return CipherText{}, errors.New("EmptyEntryFound")
	}

	key := GenKey()
	enc, _ := EncrypteIt(entry, key)
	return CipherText{
		Entry: enc,
		key:   key,
	}, nil
}

func EncrypteIt(entry string, key int64) (string, error) {
	l := len(entry)
	if l == 0 {
		return entry, errors.New("EmptyEntryFound")
	}
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

	return string(crypt), nil
}
