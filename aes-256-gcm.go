// Quick implementation of AES-256-GCM in Go.
//
// Nonce is random per each encryption.
//
// Key is passed through MD5 to output 32 bytes regardless key's size.
package aes256gcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Random bytes nonce
func generateNonce() []byte {
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic("Could not make nonce.")
	}

	return nonce
}

// Encrypts bytes with a key. Key can be of any size
func Encrypt(key []byte, data *[]byte) string {
	nonce := generateNonce()
	keyMD5 := fmt.Sprintf("%x", md5.Sum(key))

	block, err := aes.NewCipher([]byte(keyMD5))
	if err != nil {
		panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	sealed := aesgcm.Seal(nil, nonce, *data, nil)
	out := append(nonce, sealed...)

	return fmt.Sprintf("%x", out)
}

// Decrypts data with a key. Key can be of any size
func Decrypt(key []byte, data string) []byte {
	decData, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	nonce := decData[:12]
	actualData := decData[12:]
	keyMD5 := fmt.Sprintf("%x", md5.Sum(key))

	block, err := aes.NewCipher([]byte(keyMD5))
	if err != nil {
		panic(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	out, err := aesgcm.Open(nil, nonce, actualData, nil)
	if err != nil {
		panic(err)
	}

	return out
}
