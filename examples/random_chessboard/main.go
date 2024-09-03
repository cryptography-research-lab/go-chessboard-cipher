package main

import (
	"fmt"
	chessboard_cipher "github.com/cryptography-research-lab/go-chessboard-cipher"
)

func main() {

	// 随机生成一个棋盘
	chessboard := chessboard_cipher.NewRandomChessboard()
	fmt.Println(chessboard.String())
	// Output:
	// [
	//        [ Y, C, I, E, R ]
	//        [ W, O, Z, U, F ]
	//        [ K, X, L, B, Q ]
	//        [ N, T, G, M, A ]
	//        [ H, D, V, S, P ]
	// ]

	// 可以基于这个给定的棋盘去加密，这样子只要棋盘不泄露别人就无法解密
	text := "plaintext"
	encrypt, err := chessboard_cipher.PolybiusEncrypt(text, chessboard)
	if err != nil {
		panic(err)
	}
	fmt.Println(encrypt)  // Output: 553345134142143242

	// 使用同一个棋盘来解密
	decrypt, err := chessboard_cipher.PolybiusDecrypt(encrypt, chessboard)
	if err != nil {
		panic(err)
	}
	fmt.Println(decrypt)  // Output: PLAINTEXT
}
