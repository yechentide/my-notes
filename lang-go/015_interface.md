# インターフェース

`Interface(インターフェース)` とはメソッドの型だけを定義した型のこと

## 定義

インターフェースの定義の内容は、単なるメソッドリスト
```go
type 型名 interface {
    メソッド名1(引数の型, ...) (返り値の型, ...)
    .....
    メソッド名N(引数の型, ...) (返り値の型, ...)
}
```
全ての型と互換性を持っている `interface{}型（空インターフェース）` というものも定義できる
```go
interface{}

// 空インターフェースで宣言した変数にはどんな型の値でも代入可能
var obj interface{}
obj = 0123
obj = "String"
```

## 型アサーション

インターフェースの値 `<変数>` が具体的な型 `<型>` を保持し、基になる `<型>` の値を変数 `value` に代入することを主張  
２行目は、1番目の変数に型アサーション成功時に実際の値が格納され、2番目の変数には型アサーションの成功の有無（true/false）が格納される
```go
value := <変数>.(<型>)
value, status := <変数>.(<型>)
```
例：
```go
func main() {
    var intface interface{} = "hello"

    variable := intface.(string)
    fmt.Println(variable) //=> hello

    variable, ok := intface.(string)
    fmt.Println(variable, ok) //=> hello true

    float, ok := intface.(float64)
    fmt.Println(float, ok) //=> 0 false
    //格納失敗したが、成功したかの有無を確かめるokが存在するのでエラーにはならない。

    float = intface.(float64)
    fmt.Println(float) //=> panic: interface conversion: interface {} is string, not float64
    //成功したかの有無を確かめるokが存在しないのでエラーが発生する。
}
```

## 型Switch

データの型判定は `switch` 文でも行うことができる
```go
switch v := x.(type) {
case 型1: ...  // v は型1 の値になる
case 型2: ...  // v は型2 の値になる
    ...
default: ... 
}
```

## 構造体にインターフェースを実装

構造体にインターフェースを実装することのやり方
```go
func (引数 構造体名) 関数名(){
    関数の中身
}
```
例：
```go
type Person struct {} //Person構造体
type Person2 struct {} //Person2構造体

type People interface{
    intro()
}

func IntroForPerson(arg People) {
    arg.intro();
}

//Person構造体のメソッドintro()
func (p *Person) intro() { 
    fmt.Println("Hello World")
}

//Person2構造体のメソッドintro()
func (p *Person2) intro() {
    fmt.Println("Hello World")
}

func main(){
    bob := new(Person)
    mike := new(Person2) 

    IntroForPerson(bob) //=> Hello World
    IntroForPerson(mike) //=> Hello World
}
```
