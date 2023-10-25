# 関数

## 関数の基本

関数は `０〜複数` の引数を取ることができる  
引数の後, 戻り値の値に型名を書く必要がある
```go
func <関数名>([引数]) [戻り値の型] {
    [関数の本体]
}

func margeString(s1 string, s2 string) string {
    return s1+s2
}
```
連続する引数が同じ型の場合は、引数の型の省略ができる
```go
func add(x, y int) int {
    return x + y
}
```
複数の戻り値もあり得る
```go
func multipleArgs(arg1, arg2 string)(string, string){
    return arg2, arg1
}
```
可変長引数
```go
func funcA(a int, b ... int) {
    fmt.Printf("a=%d\n", a)
    for i, num := range b {
        fmt.Printf("b[%d]=%d\n", i, num)
    }
}
```
