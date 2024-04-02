# 文字列

## Goの文字列

Goの文字列は不変なバイト列である  
テキスト文字列は、`rune`の列となるが、  
`len()`はバイト数を返すため、文字数と異なる時もある  
`s[i]`はi番目のバイトを返すため、i番目の文字ではない可能性もある

### 文字列変数への代入

変数へ文字列を代入するとき、古い文字列を修正するのではなく、  
新しい文字列を生成し、それを変数が保持する。  
また、文字列は不変なため、`s[1]=a`のような操作はエラーになる

### 文字列リテラル

一般的には`"`で囲む  
エスケープ文字を処理せずに、  
複数行でもそのままになる生文字列リテラルは、`` ` ``で囲む
```go
str := `line1
line 2
line3`
```

### Unicodeエスケープ

```go
"世界"
"\u4e16\u754c"                  // 16ビットのunicodeエスケープ
"\U00004e16\U0000754c"          // 32ビットのunicodeエスケープ(稀にしか使わない)
```

## 基本な使い方

### 部分文字列

- `i`番目から`j`番目まで: `s[i:j+1]`
- 最初から`j`番目まで: `s[:j+1]`
- `i`番目から最後まで: `s[i:]`

### runeとしての処理

```go
import (
    "fmt"
    "unicode/utf8"
)

func main() {
    for i := 0; i < len(str); {
        r, size := utf8.DecodeLastRuneInString(str[i:])
        fmt.Printf("%d\t%c\n", i, r)
        i += size
    }
}
```
上のようなループ処理が必要な時が多いが、他のやり方もある
```go
str := "Go言語をやり始めました！"
for i, c := range str {
    fmt.Printf("%d: \t%q\n", i, c)
}
// または
runeArr := []rune(str)
fmt.Println(string(runeArr))    // Go言語をやり始めました！
fmt.Println(string(runeArr[3])) // 語
```

## 文字列関連の標準パッケージ

### strings

文字列の検索、置換、比較、トリミング、分割、連結ための関数がある

### bytes

バイトスライスを操作するために、stringsパッケージと似たような関数を提供している  
不変な文字列を少しずつ増やして生成していく場合、`bytes.Buffer`型を使う方が効率的  
`Buffer`は空から始まり、string、byte、[]byteなどのデータ型を書き込まれると大きくなる  
書き込むのに使う関数は、`WriteByte()`、`WriteRune()`、`WriteString()`などがある  
文字列に変える時は、`.String()`の形でやれば良い

### strconv

様々な型と文字列との変換関数と、文字列にクォーとを付けたり取り除いたりする関数を提供している
- 整数 --> 文字列  
    1. `str := fmt.Sprintf("%d", 123)`
    2. `str := strconv.Itoa(123)`
- 整数 --> 別の進数の文字列
    1. `strconv.FormatInt(int64(123), 2)` = "1111011"
- 文字列 --> 整数
    1. `num, err := strconv.Atoi("123")`
    2. `num, err := strconv.ParseInt("123", 10, 64)`：10進数で、64ビットまで(0の時はint)

### unicode

`rune`を扱うための関数を提供している
- IsDigit
- IsLetter
- IsUpper
- IsLower
- ToUpper
- ToLower

### path

`path`パッケージと`path/filepath`パッケージは、階層的な名前を操作する時に役に立つ
