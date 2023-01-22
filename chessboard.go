package chessboard_cipher

import (
	"github.com/golang-infrastructure/go-slice"
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"strings"
)

// Chessboard 定义一个棋盘，后面加密的时候就基于这个棋盘来加密，这个棋盘可以是任意大小，但是必须是每个位置仅有一个字符，并且字符不能重复
// 这个棋盘实际上就是密码表，只不过这个密码表是公开的
type Chessboard [][]rune

// Chessboard55_I 默认棋盘，保留I
var Chessboard55_I = [][]rune{
	{'A', 'B', 'C', 'D', 'E'},
	{'F', 'G', 'H', 'I', 'K'},
	{'L', 'M', 'N', 'O', 'P'},
	{'Q', 'R', 'S', 'T', 'U'},
	{'V', 'W', 'X', 'Y', 'Z'},
}

// Chessboard55_J 默认棋盘，保留J
var Chessboard55_J = [][]rune{
	{'A', 'B', 'C', 'D', 'E'},
	{'F', 'G', 'H', 'J', 'K'},
	{'L', 'M', 'N', 'O', 'P'},
	{'Q', 'R', 'S', 'T', 'U'},
	{'V', 'W', 'X', 'Y', 'Z'},
}

// NewRandomChessboard 基于给定的棋盘生成一个随机的棋盘，用于一次一密棋盘
func NewRandomChessboard(baseChessboard ...Chessboard) Chessboard {

	// 设置默认参数
	baseChessboard = variable_parameter.SetDefaultParam(baseChessboard, Chessboard55_I)

	// 把棋盘打平返回一维的格式，同时对一维的rune进行shuffle洗牌
	runeSlice := baseChessboard[0].Flat()
	slice.Shuffle(runeSlice)
	consumer := slice.NewSliceConsumer(runeSlice)

	// 按照洗牌后的一维rune组装为一个新的棋盘
	chessboard := make([][]rune, len(baseChessboard[0]))
	for indexX := range chessboard {
		chessboard[indexX] = make([]rune, len(baseChessboard[indexX]))
		for indexY := range chessboard[indexX] {
			chessboard[indexX][indexY] = consumer.Take()
		}
	}
	return chessboard
}

// Flat 返回当前棋盘展开为一维字符序列的样子返回
func (x Chessboard) Flat() []rune {
	runeSlice := make([]rune, 0)
	for _, line := range x {
		for _, column := range line {
			runeSlice = append(runeSlice, column)
		}
	}
	return runeSlice
}

// ToMapForEncrypt 返回加密用的字典表
func (x Chessboard) ToMapForEncrypt() map[rune]*Point {
	mapForEncrypt := make(map[rune]*Point)
	for lineIndex, lineValue := range x {
		for columnIndex, columnValue := range lineValue {
			mapForEncrypt[columnValue] = NewPoint(lineIndex, columnIndex)
		}
	}
	return mapForEncrypt
}

// 把棋盘转为字符串返回，用于观察棋盘长啥样
// 返回数据样例：
//
//	 [
//		[ I, C, L, O, M ]
//		[ P, H, D, R, Z ]
//		[ U, V, F, Y, B ]
//		[ G, X, T, Q, E ]
//		[ S, N, K, W, A ]
//	]
//
// .
func (x Chessboard) String() string {
	sb := strings.Builder{}
	sb.WriteString("[\n")
	for _, line := range x {
		sb.WriteString("\t[ ")
		for index, column := range line {
			sb.WriteRune(column)
			if index+1 != len(line) {
				sb.WriteString(",")
			}
			sb.WriteString(" ")
		}
		sb.WriteString("]\n")
	}
	sb.WriteString("]")
	return sb.String()
}

// ------------------------------------------------- --------------------------------------------------------------------

// Point 表示矩阵上的一个点
type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// ------------------------------------------------- --------------------------------------------------------------------
