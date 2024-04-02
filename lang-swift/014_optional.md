# オプショナル型

## オプショナル

`値がある`と`値が空`のいずれを表す型  
Swift の型は nil を許容しないので、許容するために`Optional<Wrapped>`型を使う場面がある  
`Array<Element>`型のElementのように、Wrappedも**プレースホルダ型**である

### 糖衣構文

`Optional<Int>`は、`Int?`と表記できる

### 宣言

型の後ろに `?` をつける  
初期化しない場合、`nil`となる
```swift
let 定数名:<型>?
```

### ラップ＆アンラップ

`Optional(5)`のように、値がラップされている（包まれている）
```swift
let num:Int? = 5
print(num)          // Optional(5)
```
値を取り出すための作業をアンラップという
- 強制アンラップ  
    これは一番簡単な方法である。Optional Valueに `!` をつけるだけ  
    しかし、値がないつまり`nil`のとき、エラーになる
    ```swift
    let num:Int? = 5
    print(num!)         // 5
    ```
- `??` 演算子(nil合体演算子)  
    値が`nil`の時に使う値を用意する
    ```swift
    var count:Int?
    var price:Int = 250*(count??2)
    // countがnilで、priceが500
    count = 3
    price:Int = 250*(count??2)
    // countが3で、priceが750

    x ?? y ?? z ?? 0
    ```
- 安全なアンラップ(Optional Binding)
    1. `nil`じゃなければアンラップして一時変数に代入する  
        普通のif文と同様、条件が複数あっても良い。`,`は`&&`のような役割  
        定数でなくても良い
        ```swift
        if let 定数 = オプショナルバリュー {
            // ...
        } else {
         // ...
        }

        if let 定数1=値1, let 定数2=値2 {}

        // 同じ名前を使った省略形
        let nickname: String? = "mike"
        if let nickname {
            print("Hey, \(nickname)")
        }
        ```
    2. `nil`じゃなければアンラップして一時変数に代入し、ループに入る
        ```swift
        while let 定数 = オプショナルバリュー {
            // ...
        }
        ```
    3. `nil`じゃない場合アンラップした値を定数に代入し、**guard文以降で利用できる**  
        `nil`の場合処理を中断する。
        ```swift
        guard let 定数 = オプショナルバリュー else {
            return
        }
        ```

### オプショナルチェーン

配列などのオブジェクトの値に `.` でアクセスするとき、  
対象に `?`をつけることで、値が`nil`の場合の実行時エラーを回避できる
```swift
var balls:[(size:Int, color:String)] = []
var size01 = balls.first.size               // エラー
var size02 = balls.first?.size              // size02 = nil
```
3行目のように、`balls.first`が`nil`なので、処理はそこで終わり、`nil`と返る  
また、オプショナルチェーンで取り出した値はオプショナル型なので、アンラップする必要がある
```swift
if let size = balls.first?.size {}
```

### 有値オプショナル型

アンラップしなくても使える  
変数の初期値がnilだが、その後nilになることがない時にだけ使うべき  
有値オプショナル型を==データ構造の要素として使えない==  
戻り値として使える（多くはC言語などからポインタに相当する値が返される場合）が、  
nilとなる場合は、==受け取る変数側はオプショナル型==である必要がある
```swift
var 変数名: 型!

(Int!, Int!)    // エラー
[String!]       // エラー
[String]!       // OK
```

### オプショナルでない引数に渡す場合

`!`でアンラップしてから渡す必要がある
```swift
func makeZero(_ p: inout Int) { p = 0 }
var num: Int? = 200
makeZero(&num!)
print(num!)     // 0と表示される
```

### 失敗のあるイニシャライザ

インスタンスまたはnilを返すイニシャライザ
```swift
init?(引数){
    if(条件){
        // ...
        return nil
    }
    // ...
}
```
