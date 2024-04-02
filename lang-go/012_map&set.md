# Map

初期化されていない場合は`nil`になる  
マップのキーは重複できないため、`Set`の代わりに利用できる

## 宣言方法

```go
// 組み込み関数make()を利用して宣言
var aMap = make(map[キーの型]値の型, バッファの大きさ)
var aMap = make(map[キーの型]値の型)    // バッファの大きさは省略可能

// 初期値を指定して宣言
var 変数名 map[key]value = map[key]value{key1: value1, key2: value2, ..., keyN: valueN}
```

## 操作

マップ内に該当要素がなくても、ゼロ値が返されるため、全ての操作は安全である  
nilマップに対する操作はほとんど安全だが、唯一の例外は値の代入  
マップの要素は変数ではないので、`&`でそのアドレスを取得できない  
`Slice`と同様に、マップ同士の比較はできない
```go
// 代入、追加
map[key] = elem
// 値の取得
map[key]
// 削除
delete(mapの変数, key)
// 全要素アクセス
for k, v := range myMap { ... }
// 要素の存在チェック
if value, status := myMap["key"]; !status {
    fmt.Println("存在しない")
}
```

## Sliceをキーにする

マップのキーは比較可能でないといけない、  
そのために`Slice`をキーとして利用したければ、`Slice`を比較可能なキーに変換する関数を使う  
例えば文字列に変換する関数
