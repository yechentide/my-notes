# 変数＆定数

## 宣言

### まとめて宣言

定数も変数も、まとめて宣言できる
```go
const (
    a = 1
    b = 2
)
var (
    a = 1
    b = 2
)
```

### 定数

`const`キーワードを使用して宣言する  
定数は、`文字(character)`、`文字列(string)`、`boolean`、`数値(numeric)`のみで使える  
==定数は`:=`を使用して宣言することはできない==

#### 定数生成器iota

const宣言では、iotaの値はゼロから始まり、順番に個々の項目ごとに1増加する
```go
const (
    zero int = iota     // 0
    one                 // 1
    two                 // 2
)
const (
    a uint = 1 << iota  // 2進数の 1
    b                   // 2進数の 10
    c                   // 2進数の 100
    d                   // 2進数の 1000
    e                   // 2進数の 10000
)
```

#### 型付けなし(untyped)定数

多くの定数は、特定の型への結び付けを遅延させられている  
そのような定数は、内部ではより高い精度で表現され、多くの式で変換せずに使われる  
型付けなし定数は6種類ある：  
`型付けなし論理値`、`型付けなし整数`、`型付けなしルーン`、  
`型付けなし浮動小数点数`、`型付けなし複素数`、`型付けなし文字列`

### 変数

`var`キーワードで宣言する  
複数の変数の後に型を書くことで複数の変数を同時に定義できる  
初期値が与えられている場合、変数の型宣言は必要ない  
初期化されてない場合、数値型は0、真理型はfalse、文字列型は空文字列になる(ゼロ値)
```go
var num int
var var1, var2, var3 bool

var str = "Go Programming Language"
var num = 2

var str, num, bool = "Go language", 23, true
```
関数内では`:=`を使った変数宣言もできる  
`:=`で他変数宣言する時、すでに存在している変数にそのまま値を代入する  
ただし、少なくとも１つ以上の新しい変数を宣言する必要がある
```go
func main(){
    str := "Hello World"
    i, j := 0, 1
}
```

## 代入

- 普通の代入: `=` `++` `--` `*=`など
- タプル代入: `a, b, c = 1, 2, 3`

### `_`(ブランク識別子)

値が必要ない場合、これに代入する  
そうしないと、使われないローカル変数はエラーになる

## その他

### 名前に使える文字

- アルファベット（大文字小文字は区別）
- 数字
- Unicode文字（漢字など）
- `_`

### 名前に関するルール

- 大文字で始まる：パッケージ外から見える
- 長さ制限なし、短い傾向がある
- スコープが広ければ広いほど、名前が長い
- キャメルケース
