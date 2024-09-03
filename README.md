# 棋盘密码

# 一、这是什么？ 

棋盘密码是一种古典密码学中的加密方法，它使用一个预定义的矩阵（通常是一个5x5或6x6的棋盘）来加密信息。这种加密方式的基本原理是将字母表中的每个字母映射到矩阵中的特定行和列。明文信息通过查找每个字母在矩阵中的位置并记录下相应的行和列编号来转换为密文。解密过程则是加密的逆过程，通过密文提供的行和列编号在矩阵中找到对应的字母来恢复明文。

棋盘密码有几种不同的变体，包括：

1. **波利比奥斯棋盘（Polybius Checkerboard）**：这是一种5x5的棋盘，由希腊人Polybius在公元前2世纪发明。它将字母表中的26个字母（I和J共享同一个格子）映射到棋盘的行和列上。明文通过转换为对应的行和列编号来加密，例如“hello”可以加密为“2315313134”。解密时，需要将这些编号转换回字母。

2. **ADFGX棋盘密码**：这是第一次世界大战中德军使用的一种加密方法，是Polybius棋盘密码的改良版本。它使用一个5x5的棋盘，但只使用ADFGX这五个字母，因此每个字母的密文由两个ADFGX中的字母组成。ADFGX密码还可以使用一个密钥进行二次加密，增加安全性。

3. **ADFGVX棋盘密码**：这是ADFGX棋盘密码的加强版本，增加了一个行列，用V表示，形成一个6x6的棋盘。这样，除了字母，还可以加密数字0到9，总共可以表示36个字符。ADFGVX棋盘密码的特点是明文可以是字母或数字，而密文则由ADFGVX这六个字母组成。

棋盘密码的安全性主要依赖于密钥的保密性，一旦密钥泄露，加密就容易被破解。在实际使用中，密钥的管理和保护非常重要。尽管棋盘密码在现代加密算法面前显得较为简单，但它在密码学的历史和文化价值上仍具有重要意义，并且可以作为教学工具来教授密码学的基本原理。

# 二、安装

```bash
go get -u  github.com/cryptography-research-lab/go-chessboard-cipher
```

# 三、API Example

## 3.1 polybius

```golang
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
```

## 3.2 custom_chessboard

```golang
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
```

## 3.3 random_chessboard

```golang
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
```


# 四、相关题目
- https://www.bilibili.com/read/cv3145647

# 五、参考资料

- https://xz.aliyun.com/t/3603

# 六、TODO

- 兼容polybius的I/J
