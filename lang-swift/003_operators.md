# 演算子

## ルール(前後の空白)

`a+b` と `a + b`は大丈夫だが、`a+ b`や`a +b`はエラーになる

## 一覧

- 算術演算子：　`+` `-` `*` `/` `%`
- オーバーフロー演算子：　`&+` `&-` `&*`
- 論理演算子：　`&&` `||` `!`
- 比較演算子：　`>` `<` `>=` `<=` `==` `!=`
- 代入演算子：　`=`
- 複合代入演算子：　`+=` `-=` `*=` `/=` `%=` `&=` `^=` `|=` `<<=` `>>=` `&<<=` `&>>=`  
    Swiftでは`++`(インクリメント)と`--`(デクリメント)がない
- キャスト演算子：　`is` `as` `as?` `as!`
- クラスのインスタンスの実体が同一かどうか：　`===` `!==`
- 三項演算子：　` ? : `
    ```swift
    条件式 ? trueの場合の式 : falseの場合の式
    ```
- 範囲演算子：　`..<` `...`
    ```swift
    1..<3       // 1 <= x < 3  となる範囲
    1...3       // 1 <= x <= 3 となる範囲
    ```
- ビット演算子： 2進数は`0b`, 16進数は`0x`から始まる
  1. bitシフト：　`<<` `>>` `&<<` `&>>`
  2. bit積、bit和、排他的bit和、bit否定：　`&` `|` `^` `~`

## 演算子の定義

### 演算子の宣言

３つのパターンがある  
`パターン1`や`パターン2`の場合、トップレベルで演算子として使う記号(or記号列)と使い方(前置・後置・二項)を宣言する必要がある  
`パターン3`の場合は、改めて宣言する必要がない
1. 完全に新しい演算子を定義する
2. 既存の単項演算子に、二項演算子の機能を定義する
3. 既存の演算子で、違うデータ型(の組み合わせ)に対して適用可能にする
```swift
// 二項演算子
infix operator 演算子 : 優先度グループ名
// 前置演算子
prefix operator 演算子
// 後置演算子
postfix operator 演算子
```
「優先度グループ名」は==省略可能 ---> DefaultPrecedence==になる  
標準ライブラリにある演算子は、==優先度グループ名の再定義ができない==

### 定義を追加できない演算子

- 普通の代入演算子`=`
- キャスト演算子
- 関数宣言で利用される`->`
- 名前の修飾で利用される`.`
- コメントで利用される`//` `/*` `*/`
- 三項演算子、オプショナルチェーンなどの`?`
- 前置演算子として使えない。inoutパラメータの指定用の`&`
- 前置演算子として使えない。型パラメータの括弧`<`
- 後置演算子として使えない。型パラメータの括弧`>`
- 後置演算子として使えない。開始指示子`!`

### 演算子として使える文字列

各国のアルファベット、漢字、仮名、絵文字は不可  
`.`で始まる場合に限って、`.`や他のASCII文字を組み合わせて利用できる  
==例：==　`.&`や`...`はOK、`=.=`や`%.`はNG
- ASCII文字の範囲内：　`/ = - + ! * % < > & | ^ ~ ?`
- Unicodeの範囲内：矢印、数学演算子、句読点、CJK記号など

### 二項演算子の定義

引数ラベルは必要がない  
引数は２つだけ。左側のオペランドは第１引数、右側のオペランドは第２引数になる。
```swift
infix operator +- : RangeFormationPrecedence
func +- (number: Int, width: Int) -> String {
    // .....
    return ...
}
```

#### カスタムオペレータ

数値と文字列を`+`演算子で簡単に連結できるように拡張する
```swift
func + (num:Int, str:String) -> String {
    return String(num) + str
}
```

### 単項演算子の定義

パーセントを表す演算子
```swift
postfix operator %
postfix func % (num: Int) -> Double {
    return 0.01 * Double(num)
}
postfix func & (num: Double) -> Double {
    return 0.01 * num
}
```

### 引数自身の値を変える

上の例は新しい値を返すだけ  
C言語の`++`などは、呼び出す側の変数の値が変化するものもある  
関数のときと違って、呼び出す側は、`&`をつける必要がない
```swift
postfix operator ++
postfix func ++ (num: inout Int) {
    num += 1
}
```

### Swiftの優先度グループ

優先度高いから低い順
- BitwiseShiftPrecedence
- MultiplicationPrecedence
- AdditionPrecedence
- RangeFormationPrecedence
- CastingPrecedence
- NilCoalescingPrecedence
- ComparisonPrecedence
- LogicalConjunctionPrecedence
- LogicalDisjunctionPrecedence
- ==DefaultPrecedence==
- TernaryPrecedence
- AssignmentPrecedence

### 新しい優先度グループの定義

```swift
precedencegroup グループ名 {
    associativity: 指定
    higherThan: 他のグループ名
    lowerThan: 他のグループ名
    assignment: 値
}
```
- associativityは結合規則、`right` `left` `none`から１つを指定する。無指定だと`none`になる
- assignmentは代入演算子かどうかを表す、値は`true`か`false`  
    ここで指定しなくても`inout`で値を変更できるけど、オプショナルチェーンの時の代入ができなくなる
