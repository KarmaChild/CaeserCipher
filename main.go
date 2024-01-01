package main

import "fmt"

func main() {
	cipher := Cipher{}
	rounds := cipher.randomRounds
	fmt.Printf("%d\n", rounds)
}
