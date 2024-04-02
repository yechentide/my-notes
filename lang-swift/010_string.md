# 文字列

## String型

### 書記素クラスタとCharacter型

書記素クラスタとは、合成された文字のこと  
Character型は１つ以上の文字コードを組み合わせた書記素クラスタを格納できる
```swift
let gu: Character = "\u{30B0}"              // グ
let ku: Character = "\u{30AF}\u{3099}"      // ク + ゛
```

### 文字列型

`String`型：文字列を扱う。通常はこれを使う  
`+`演算子または`+=`演算子で文字列どうしを連結できる  
**String型とCharacter型は、`append(_:)`メソッドを使う**
```swift
let str01 = "Hello"
let bird: String
bird = "アマツバメ"

let str02 = "Hello" + String(2019+1)

// 空の文字列
let str03 = String()
let str04 = ""

// 同じ文字の繰り返し
let str05 = String(repeating:"★", count:10)    // ★★★★★★★★★★
```

### String型のプロパティ

- `isEmpty`：　空の文字列かどうか (true/false)
- `count`：　文字列の文字数

### String型のメソッド

- `lowercased()`：　全ての文字を小文字に
- `uppercased()`：　全ての文字を大文字に
- `hasPrefix(str)`：　引数と前方一致 (true/false)
- `hasSuffix(str)`：　引数と後方一致 (true/false)
- `contains(str)`：　引数を文字列に含まれているかどうか (true/false)
- `range(target)`：  
    targetの範囲を確定して、  
    `upperBound`でtargetより後ろのものを取り出す。`lowerBound`もある  
    `removeSubrange()`でtargetを削除する
    ```swift
    var str = "東京都千代田区神南1-2-3"
    let result = str.range(of: "東京都")    // String?型

    if let range = result {
        let afterStr = str[range.upperBound...]
        print(afterStr)         // 千代田区神南1-2-3
    }

    if let range = result {
        str.removeSubrange(range)
        print(str)              // 千代田区神南1-2-3
    }
    ```

### 文字列に変数を埋め込む

文字列リテラル内で値を展開したい時、`\()`というエスケープシーケンスを使う  
`"結果：\(result)"`

### 複数行文字列

- 複数行文字リテラル  
    `"""`で囲む  
    `"""`の終了のインデントと一致する限り、引用された各行の先頭のインデントは削除される
    ```swift
    let apples = 3
    let oranges = 5
    let quotation = """
    I said "I have \(apples) apples."
    And then I said "I have \(apples + oranges) pieces of fruit."
    """
    ```

### エスケープ文字を使わない文字列の表現

1. `#"  "#`で文字列を囲む  
    文字列の中に`#`がある場合、``##"  "##``で囲む
    ```swift
    let raw1 = #" \Solution="ε" "#
    let raw2 = ##" \Color="#ff0" "##
    ```
2. `#"内容"#`という形式で、文字列埋め込みや特殊文字を指定できる。`\#`で記述する
    ```swift
    let rgb = "ff0"
    let raw3 = #" \Solution="\#u{3b5}" "#
    let raw4 = ##" \Color="#\##(rgb)" "##
    ```
3. `#"""内容"""#`は複数行にわたる文字列を扱う

### String型の比較

- `==`演算子を使う  
    暗黙な型変換がないため、**型を揃えてから比較**しよう〜
- `>` `<` `>=` `<=`演算子を使う  
    アルファベット順、かな順の比較

### 数値⇄文字列

- 文字列→数値
    ```swift
    Int("123")
    Double("111.23")
    ```
- 数値→文字列
    ```swift
    String(123)
    (11.2 + 22.3).description
    "\(123.321)"
    ```

## 便利な拡張機能

```swift
/// String extension for convenient of substring
extension String {
    /// Index with using position of Int type
    func index(at position: Int) -> String.Index {
        return index((position.signum() >= 0 ? startIndex : endIndex), offsetBy: position)
    }

    /// Subscript for using like a "string[i]"
    subscript (position: Int) -> String {
        let i = index(at: position)
        return String(self[i])
    }

    /// Subscript for using like a "string[start..<end]"
    subscript (bounds: CountableRange<Int>) -> String {
        let start = index(at: bounds.lowerBound)
        let end = index(at: bounds.upperBound)
        return String(self[start..<end])
    }

    /// Subscript for using like a "string[start...end]"
    subscript (bounds: CountableClosedRange<Int>) -> String {
        let start = index(at: bounds.lowerBound)
        let end = index(at: bounds.upperBound)
        return String(self[start...end])
    }

    /// Subscript for using like a "string[..<end]"
    subscript (bounds: PartialRangeUpTo<Int>) -> String {
        let i = index(at: bounds.upperBound)
        return String(prefix(upTo: i))
    }

    /// Subscript for using like a "string[...end]"
    subscript (bounds: PartialRangeThrough<Int>) -> String {
        let i = index(at: bounds.upperBound)
        return String(prefix(through: i))
    }

    /// Subscript for using like a "string[start...]"
    subscript (bounds: PartialRangeFrom<Int>) -> String {
        let i = index(at: bounds.lowerBound)
        return String(suffix(from: i))
    }
}
```
