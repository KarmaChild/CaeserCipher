package main

import (
	"CaeserCipher/cipher"
	"fmt"
)

func main() {
	c := cipher.Cipher{}
	//rounds := c.RandomRounds()
	//text := "Hello World"
	enc := c.Encrypt("aa", 20)
	dec := c.Decrypt(enc, 20)
	fmt.Printf("encrypt %s\n", enc)
	fmt.Printf("decrypt %s\n", dec)
}
