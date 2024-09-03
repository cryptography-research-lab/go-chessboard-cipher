package main

import (
	"fmt"
	chessboard_cipher "github.com/cryptography-research-lab/go-chessboard-cipher"
)

func main() {

	text := "plaintext"
	// 使用自定义的棋盘布局来加密
	encrypt, err := chessboard_cipher.PolybiusEncrypt(text, chessboard_cipher.Chessboard55_J)
	if err != nil {
		panic(err)
	}
	fmt.Println(encrypt)  // Output: 353111243344155344


}
