# UIテスト

## ツール

- 公式(Xcode標準)：`XCTest Framework`
- サードパティ：
    1. `EarlGrey`
    2. `Appium`

## XCTest Framework

### 実装の流れ

1. `XCUIApplication`クラスを用いてアプリを起動させる
2. 起動した対象アプリ(`XCUIApplication`)に対して、`XCUIElementQuery`を使用して、UI要素`XCUIElement`を見つける
3. 見つけたUI要素`XCUIElement`に対して行いたい操作(tapメソッドなど)を行う
4. 操作した結果に対してXCTestアサーションを使用し、UI要素などの状態を期待値と実際の値で比較する
