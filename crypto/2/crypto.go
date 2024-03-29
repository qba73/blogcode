package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func message(s string) (string, error) {
	hmac := hmac.New(sha256.New, []byte("our-test-key"))
	_, err := io.WriteString(hmac, s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hmac.Sum(nil)), nil
}

var usage = `Usage: crypto message`

func Main() int {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return 1
	}

	m, err := message(os.Args[1])
	if err != nil {
		fmt.Println(usage)
		return 1
	}
	fmt.Println(m)
	return 0
}
