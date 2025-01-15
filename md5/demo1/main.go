package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("1qaz!QAZ"))
	fmt.Printf("%x", h.Sum(nil))
	fmt.Printf("%x", h.Sum(nil))
}
