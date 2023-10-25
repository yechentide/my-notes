# 制御構文

## For

初期化ステートメントでは、短い変数宣言はよく利用される  
`()`は必要ないが、`{}`は必要  
真ん中は省略可能(`;`も省略可能)
```go
for i := 0; i < 10; i++ {
    sum += i
    fmt.Println(sum)  //=> 0 1 3 6 10
}

var i = 10
for i > 3 {
    fmt.Println(i)
    i--
}
```

## If

`for`と同様に`()`は必要ないが、`{}`は必要  
条件の前に、評価するための簡単なステートメントを書ける  
（ここで宣言された変数は、`if`のスコープ内だけで有効）
```go
func condition(arg string)string{
    if v := "GO"; arg == v {
        return "This is Golang"
    }else{
        return "This is not Golang"
    }
}
```

## Switch

選択されたcaseだけを実行して、それに続く全てのcaseは実行されない  
そうしたい場合は`fallthrough`を使う
```go
lang := "Go"
switch lang {
case "Ruby":
    fmt.Println("This is Ruby")
case "Go":
    fmt.Println("This is Go")
default:
    fmt.Println("This is a programming language")
}
// => This is Go
```
`switch`の右に何も条件を書かない場合は`switch true`と書くのと同じ  
`if`文をよりシンプルに書ける
```go
lang := "Go"
switch {
case lang == "Ruby":
    fmt.Println("This is Ruby")
case lang == "Go":
    fmt.Println("This is Go")
default:
    fmt.Println("This is a programming language")
}
// => This is Go
```

## Defer（遅延実行）

defer文を定義すれば、関数を抜ける前に必ず実行される  
中断処理やエラーハンドリング処理などで有効  
一つの関数内で複数のdefer文を定義できる  
後から定義したものから順に実行される
```go
func main(){
    defer fmt.Println("Golang") //defer1
    defer fmt.Println("Ruby") //defer2
    fmt.Println("JS") 
    //=> JS
    //=> Ruby
    //=> Golang
}
```
