package chessboard_cipher

import (
	"fmt"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"strings"
)

// PolybiusEncrypt Polybius棋盘加密，传入的text必须全是英文字母组成
func PolybiusEncrypt(text string, chessboard ...Chessboard) (string, error) {

	// 如果没有传递棋盘的话，则使用默认的棋盘
	chessboard = variable_parameter.SetDefaultParam(chessboard, Chessboard55_I)

	encryptMap := chessboard[0].ToMapForEncrypt()
	buff := strings.Builder{}
	for _, char := range text {
		point := encryptMap[char]
		if point == nil {
			return "", ErrEncryptText
		}
		buff.WriteString(fmt.Sprintf("%d%d", point.X+1, point.Y+1))
	}
	return buff.String(), nil
}

// PolybiusDecrypt Polybius棋盘解密，传入的text必须全是数字组成，否则返回错误
func PolybiusDecrypt(text string, chessboard ...Chessboard) (string, error) {
	panic("")
}
