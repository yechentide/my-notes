# 多言語対応

## アプリ名

アプリ名以外の、`Info.plist`内のものは全てこの方法で
- `Info.plist`で、`CFBundleDisplayName = “App”`(デフォルトの名前)
- 新しく`InfoPlist.strings`を作る
- ユーティリティエリア(右側)のFile inspector内の`Localize...`ボタンをクリック
- `CFBundleDisplayName = "アプリ";`を記述(ローカライズ)

## 日付や通貨

`Formatter`を利用した方が良い
- NSDateFormatter
- NSNumberFormatter

```swift
# String.localizedStringWithFormatと同様に三桁区切りなどを国際化する例
# あくまで例で、この場合はString.localizedStringWithFormatを使う方が良いです
var formatter = NSNumberFormatter();
formatter.formatterBehavior = .Behavior10_4
formatter.numberStyle = .DecimalStyle
formatter.stringFromNumber(12345)!
```

## 文字列

- プロジェクト設定で、対応言語を増やす
- `Strings File`を追加する。名前は`Localizable.strings`とする必要がある。
- ユーティリティエリア(右側)のFile inspector内の`Localize...`ボタンをクリック
- 文字列追加
    - `"キーとなる文字" = "表示したい文字（とフォーマット）";`
    - 最後に`;`が必要
    - `%@`, `%d`, `%f`などのフォーマットを使える
    - `String(format: _, arguments: _)`でフォーマットに引数を渡す
    - `NSLocalizedString("キーとなる文字", comment: "")`
    - キー文字列の先頭に目立つ絵文字を入れると、ローカライズもれ対策となる

### リソーステーブル

デフォルトのリソーステーブルは`Localizable.strings`だが、それと別に作っても大丈夫。  
以下のように利用する。
```swift
func NSLocalizedString(key: String, tableName: String? = default, bundle: NSBundle = default, value: String = default, #comment: String) -> String

// Home.stringsの中のRunButtonキーワード
NSLocalizedString("RunButton", tableName: "Home", comment: "")
```

### フォーマット

ローカライズの文字列には[フォーマット](https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Strings/Articles/formatSpecifiers.html#//apple_ref/doc/uid/TP40004265-SW1)を使える。
- `%@`: 文字列
- `%d`: 整数
- `%f`: 小数
- 引数の順番で指定できる
    - `%1$@`: １つ目の文字列
    - `%2$@`: ２つ目の文字列

引数を別の関数で渡す必要がある。
```swift
String(format: NSLocalizedString("HogeHoge", comment: ""), "aaa", "bbb")

static func localizedStringWithFormat(format: String, _ arguments: CVarArgType...) -> String
String.localizedStringWithFormat("output: %d", num)
```

### SwiftUIでのローカライズ

上は`NSLocalizedString`関数でリソーステーブル内キーワードを使うが、  [SwiftUIでは](https://developer.apple.com/documentation/swiftui/preparing-views-for-localization)少し違う。  

- `Text()`の場合、そのままキーワードを引数にすれば、自動的に`Localizable.strings`で探してくれる。
- 引数渡しは、`\()`を使っても良い: `Text("Copying \(copyOperation.numFiles) files")`
- 変数や定数に代入する場合、`LocalizedStringKey`を使う必要がある。
- キーワードとしてじゃなく、そのまま表示させたい場合: `Text(verbatim: "test")`

## 参考サイト

- [ローカリゼーション](https://developer.apple.com/jp/localization/)
- [iOSアプリの国際化対応の勘所とTips集(Swift版)](https://qiita.com/mono0926/items/c41c1ce18b90b765a8f2)
- [Xcode 6で導入されたstringsdictファイルを使って、イケてる複数形多言語対応](https://qiita.com/mono0926/items/647f6d741cd9965d9806)
- [Sample Code: Displaying Human-Friendly Content](https://developer.apple.com/documentation/foundation/formatter/displaying_human-friendly_content)
