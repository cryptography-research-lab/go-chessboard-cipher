package chessboard_cipher

// ToUpperCase 转换为大写字母
func ToUpperCase(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return 'A' + (c - 'a')
	} else {
		return c
	}
}

// ToLowerCase 转换为小写字母
func ToLowerCase(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return 'a' + (c - 'A')
	} else {
		return c
	}
}
