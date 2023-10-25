# 辞書型

## 辞書

`Dictionary<Key, Value>`型  
マップ型、辞書型などと呼ばれる  
**同じキーを持った**辞書を複数個作るとき、辞書ではなく構造体を使うべき

### 辞書の作成

定数＆変数があり、型推論＆型アノテーションがある
```swift
let 辞書名:[keyの型:valueの型] = [キー:値, ...]
let 辞書名:Dictionary<keyの型, valueの型> = [キー:値, ...]

// 空の辞書
let emptyDictionary = [keyの型, valueの型]()
let emptyDictionary = Dictionary<keyの型, valueの型>()
let emptyDictionary:[keyの型:valueの型] = [:]
```

### KeyとValueにできる型

- Keyの型に制限あり  
    Keyの型の値をもとに、ハッシュ値を計算できるような型でなければならない
- Valueの型に制限なし、タプルを値としても使える

### プロパティ

- `isEmpty`：　辞書が空かどうか (true/false)
- `count`：　辞書の要素の個数
- `keys`：　キーのコレクションを返す。`辞書.keys.contains("key")`で使える
- `values`：　値のコレクションを返す。

### 値のアクセス

キーが存在しないときnilと返ってくるので、オプショナル型である
```swift
辞書名[キー]
```
オプショナル値を計算式などで使うとき、アンラップする必要がある。  
末尾に `!` をつければ、強制的にアンラップした値を使って計算できる
```swift
var members = ["東京":15, "大阪":12, "札幌":9]
let tokyoValue = members["東京"]
let osakaValue = members["大阪"]
print(tokyoValue! + osakaValue!)            // 27
```
あるキーに対応する値が存在しないとき、既定値を指定できる
```swift
let okinawa = members["沖縄", default:100]
```
全ての値を取り出したい場合、`for-in`文を使う
```swift
let dic:[String:Int] = [.......]
for pair in dic {
    print(pair.key + pair.value)
}

for (key,value) in dic {
    // ...
}
```

### 値の追加＆変更

```swift
// 変更
辞書名["キー"] = 値

// 追加
辞書名["存在しないキー"] = 値
```

### 要素の削除

- 一つの要素を削除
    1. `removeValue(forKey:キー)`メソッドを使う  
        指定したキーが辞書に存在しない場合、`nil`を返すため、オプショナル型となる  
        キーが存在した場合、該当要素を削除し、そのキーの値を返す
    2. 値を`nil`にする
- 全ての要素を削除: `removeAll()`

### 辞書の複製

辞書を別の変数に代入したとき、参照ではなく複製となる（レイジーコピー）
