# encryptionWithGo

## Installation

- Install latest version

```
go get -u github.com/Jay-ponda-improwised/encryptionWithGo
```

- Install spacific version

```
go get -u github.com/Jay-ponda-improwised/encryptionWithGo@v1.1.0
```

#

## Imports and package use

- Import package and use alice localCrypto

```
import localCrypto "github.com/Jay-ponda-improwised/encryptionWithGo"
```

#

## Core info about structure

### Structure

- A struct `CipherText` used to store (encrypted ext with key)

1. Entry (It stores encrypted text)
2. key (It stores key) <br/><br/>

### Encryption

1. Encrypt(entry string) (CipherText, error)
   - params:
     - `entry (string)`: Entry takes string which you want to encrypt.
   - returns:
     - `CipherText, error`: Which is struct with encrypted-string and key, error

<br/>

2. EncrypteIt(entry string, key int64) (string, error)

   - params:
     - `entry (string), key (int64)`: Entry takes string which you want to encrypt, custom key
   - returns:
     - `string, error`: encrypted string, error

<br/>

### Decryption

1. (d CipherText) Decrypt() (string, error)

   - returns:
     - `string, error`: plain text (decrypted text), error

<br/>

4. DecryptIt(encEntry string, key int64) (string, error)

   - params:
     - `entry (string), key (int64)`: Entry takes string which you want to encrypt.
   - returns:
     - `string, error`: plain text (decrypted text), error

<br/>

#

## Errors

- `EmptyEntryFound`: If given entry is empty ("")
