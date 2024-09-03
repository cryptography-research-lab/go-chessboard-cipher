package chessboard_cipher

import (
	"fmt"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
)

// PolybiusEncrypt Polybius棋盘加密，传入的text必须全是英文字母组成
func PolybiusEncrypt(plaintext string, chessboard ...Chessboard) (string, error) {

	// 如果没有传递棋盘的话，则使用默认的棋盘
	chessboard = variable_parameter.SetDefaultParam(chessboard, DefaultChessboard)

	encryptMap := chessboard[0].ToMapForEncrypt()
	runes := make([]rune, len(plaintext)*2)
	for index, char := range plaintext {

		// 转换为大写字母
		char = ToUpperCase(char)
		// 必须是大写字母，否则的话认为是不合法的输入
		if !IsAlpha(char) {
			return "", ErrEncryptText
		}

		point := encryptMap[char]
		if point == nil {
			return "", ErrEncryptText
		}
		runes[index*2] = rune(point.X + '1')
		runes[index*2+1] = rune(point.Y + '1')
	}
	return string(runes), nil
}

// PolybiusDecrypt Polybius棋盘解密，传入的text必须全是数字组成，否则返回错误
func PolybiusDecrypt(ciphertext string, chessboard ...Chessboard) (string, error) {

	// ciphertext的长度必须是偶数
	if len(ciphertext)&1 != 0 {
		return "", fmt.Errorf("the length of the ciphertext must be even")
	}

	// 如果没有传递棋盘的话，则使用默认的棋盘
	chessboard = variable_parameter.SetDefaultParam(chessboard, DefaultChessboard)

	// 开始按照位置找字符
	runes := make([]rune, len(ciphertext)/2)
	for i := 0; i < len(ciphertext); i += 2 {

		// 减去1是减去加密时加上来的一个偏移量
		x := int(ciphertext[i] - '1')
		y := int(ciphertext[i+1] - '1')

		// x轴棋盘越界检查
		if x < 0 || x >= len(chessboard[0]) {
			return "", fmt.Errorf("the value %d at the position %d of the input ciphertext is out of bounds", i, x)
		}

		// y轴棋盘越界检查
		if y < 0 || y >= len(chessboard[0][x]) {
			return "", fmt.Errorf("the value %d at the position %d of the input ciphertext is out of bounds", i+1, y)
		}

		runes[i/2] = chessboard[0][x][y]
	}

	return string(runes), nil
}
