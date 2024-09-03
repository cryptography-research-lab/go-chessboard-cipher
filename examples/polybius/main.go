package main

import (
	"fmt"
	chessboard_cipher "github.com/cryptography-research-lab/go-chessboard-cipher"
)

func main() {

	text := "plaintext"
	encrypt, err := chessboard_cipher.PolybiusEncrypt(text)
	if err != nil {
		panic(err)
	}
	fmt.Println(encrypt)  // Output: 353111243344155344

}
