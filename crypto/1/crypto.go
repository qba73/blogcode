package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	m, err := message("test@acme.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}

func message(s string) (string, error) {
	hmac := hmac.New(sha256.New, []byte("our-test-key"))
	_, err := io.WriteString(hmac, s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hmac.Sum(nil)), nil
}
