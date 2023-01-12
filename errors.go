package chessboard_cipher

import "errors"

// ErrEncryptText 传入的要加密的文本不符合要求，传入的要加密的文本必须全是英文字母
var ErrEncryptText = errors.New("text for encrypt must all letter")

// ErrDecryptText 传入的要解密的文本不符合要求，传入的要解密的文本必须全是数字，并且是偶数个
var ErrDecryptText = errors.New("text for decrypt must all number")
