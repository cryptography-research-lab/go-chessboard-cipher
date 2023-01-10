package chessboard_cipher

import (
	"testing"
)

func TestChessboardRandom55(t *testing.T) {
	random55 := NewRandomChessboard()
	t.Log("随机棋盘： \n" + random55.String())
}
