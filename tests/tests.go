package main

import (
	"fmt"
	"strconv"
	"time"

	aes256gcm "github.com/tgbv/aes-256-gcm"
)

func main() {
	key := "Some random key"
	data := []byte("some random data lalalal very long string abcdEFGH...." + strconv.Itoa(int(time.Now().Unix())))

	enc := aes256gcm.Encrypt([]byte(key), &data)
	fmt.Println(enc)

	dec := aes256gcm.Decrypt([]byte(key), enc)
	fmt.Println(string(dec))
}
