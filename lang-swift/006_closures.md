# クロージャ

## 概要

```swift
{
    [キャプチャリスト]           // 省略可能
    (仮引数: 型) -> 型 in       // 仮引数、戻り値は省略可能
    // 処理
}
```
クロージャは、内部に状態を保てる無名関数のようなものである。様々な応用が可能。  
引数について関数と似ている。デフォルト値の指定ができない  
return文１行のみの場合、returnを省略して式だけでも良い  
戻り値なしは、`-> ()`または`-> Void`で表す

## クロージャの`引数`と`->型`の省略

`引数`がなければ、省略可能  
`->型`は、クロージャの構造＆それが使われる状況から、戻り値の型を推論できる場合にのみ省略可能
1. ブロック内は１文のみ：クロージャ構造から推論できる
2. 代入される変数＆引数などの型からも推論できる
```swift
// 完全形
var clos: (Int,Int) -> String = { (a:Int, b:Int) -> String in return "\(a)/\(b)" }

// 省略形
var clos: (Int,Int) -> String = { a,b in "\(a)/\(b)" }
var clos: (Int,Int) -> String = { "\($0)/\($1)" }
```

## 関数とクロージャの型

関数とクロージャは型を持っていて、`(引数の型, ...) -> 戻り値の型`と表す

## 変数キャプチャ

先頭に変数名を`[ ]`で囲む部分はキャプチャリストという  
クロージャ生成時に変数のコピーが作成され、その後の変更は元のデータに影響しない
```swift
let a: () -> () ={
    [変数, 定数, ...] in
    /*処理*/
}
```

## @escaping属性

引数として渡されたクロージャが、関数の実行が終わった後に実行される可能性があるとき、`@escaping`属性をつける必要がある。[Swiftで理解する@escapingの9選サンプルコード](https://jp-seemore.com/app/15990/#toc2)
```swift
var todos: [() -> Void] = []

func afterAll(task: @escaping () -> Void) {
    print("aaa")
    todos.append(task)
}

afterAll {
    print("bbb")
}

print(todos.count)
if !todos.isEmpty {
    todos.first!()
}

/* Output:
aaa
1
bbb
*/
```

## @autoclosure属性

関数に渡された引数をクロージャにラップできる
```swift
var customerInLine = ["Mike", "Jon", "Bob"]

// @autoclosure属性なし
func serve(customer customerProvider: () -> String) {
    print("Now serving \(customerProvider())")
}
serve(customer: { customerInLine.remove(at: 0)　} )

// @autoclosure
func serve(customer customerProvider: @autoclosure () -> String) {
    print("Now serving \(customerProvider())")
}
serve(customer: customerInLine.remove(at: 0))
```
