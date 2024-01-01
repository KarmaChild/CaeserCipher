package main

import (
	"CaeserCipher/cipher"
	"fmt"
)

func main() {
	c := cipher.Cipher{}
	rounds := c.RandomRounds()
	text := "You can edit this code!"
	enc := c.Encrypt(text, rounds)
	dec := c.Decrypt(enc, rounds)
	fmt.Printf("encrypt %s\n", enc)
	fmt.Printf("decrypt %s\n", dec)
}
