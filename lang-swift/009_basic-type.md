# 基本型

## 代入時の違い

### 値型

値型のデータを渡す時、まず参照を渡す  
書き込みが発生してから、データを複製し始める（==Copy-On-Write==）  
渡した後の操作は元データに影響しない  
値型のものを定数に代入すると、その値を変更できなくなる  
==Swiftでは大体のものが値型==
- String型
- 配列
- タプル
- 構造体

### 参照型

データ自体に対する参照を代入される。  
参照型のものを定数に代入すると、参照先を変えられないが、その値を変えられる
- クラスのインスタンス
- クロージャのインスタンス
- NS系列

### 値型データの共有

シークエンス(配列など)の一部を部分列として取り出すとき、==Copy-On-Write==の手法が使われるけど、  
複製を作るのではなく、部分列の情報(元のどの部分など)を効率よく保存できる別の型を作って共有した方がいい  
例えばSequenceプロトコル内では、先頭要素を取り除いた部分列を返す関数がある：
```swift
@inlinable public func dropFirst(_ k: Int = 1) -> DropFirstSequence<Self>
```
`DropFirstSequence<Self>`もシークエンスだけど、  
元のシークエンスはどれか、部分列はどこからどこまでかという情報しか持たない(つまりサイズが小さい)  
要素を取り出す必要があれば、計算を行って元のシークエンスから値を取り出す  
実際に`DropFirstSequence<Self>`を使って変数やメソッドを定義する必要がない  
配列(Array)や文字列(String)にも同じ仕組みがある  
部分列を取り出すと、`ArraySlice`や`String.SubSequence`という型になる  
ただし、このような共有のためのデータが使われている間、==元のデータはメモリから解放されない==

## リテラル

- 数値リテラル：  
    `123456` `123_456` `123_4__56`  
    `0b01` `0o123_45670` `0x1234abcd`  
    `1.23e1` `1.23e-2`
- 文字列リテラル：  
    `"abc"`など`"`で囲むもの
- 配列リテラル：  
    `[1, 2, 3]` `["a", "b", "c"]`
- 辞書リテラル：  
    `["a":1, "b":2]`

## 特殊文字(エスケープシーケンス)

- `\n`: ラインフード
- `\r`: キャリッジリターン
- `\"`: ダブルクオート
- `\'`: シングルクオート
- `\\`: バックスラッシュ
- `\0`: null文字
- `\t`: タブ

## 主な基本型

### 数値型

- 整数  `Int` `UInt`
    1. `UInt`は0以上の数を表す
    2. 桁数に応じてInt8, Uint8, Int16, Uint16, Int32, Uint32, Int64, Uint64
    3. それぞれ`min`と`max`プロパティを持つ
- 浮動小数点数`Float` `Double`
    1. 精度が違う
    2. `infinity`プロパティを持つ

### 文字型

`Character`型：１文字(英文字？)を扱うデータ型  
文字列から１文字ずつ取り出す時はこの型になる  
文字列リテラルはデフォルトでString型になる
```swift
let str = "a"               // String型
let cha: Character = "a"    // Charater型
```

### 論理型

`Bool`型：trueかfalseのどっちかの値を必ずとる

## 型の別名

```swift
typealias Num = Int32
public typealias Void = ()
```

## 型の確認

1. 変数・定数・式の型を返す
    ```swift
    type(of: 変数名または定数名)
    ```
2. 式が型T、クラスT、Tのサブクラス、プロトコルTに準拠する時にtrue
    ```swift
    式 is 型T
    ```

## 型の変換

Swiftでは、==黙然な型変換が行われない==！
```swift
型名(数値)

let i = 123             // Int型の123
let s = String(i)       // String型の"123"

let s1 = "123"
let i1 = Int(s1)        // 123

let s2 = "abc"
let i2 = Int(s2)        // nil
```

### キャスト演算子

- `as`  
    同等な型に変換。  
    (スーパークラスに)アップキャストできる
    ```swift
    式 as 型
    ```
- `as?`  
    条件付きキャスト。戻り値はオプショナルで、型変換が失敗した場合はnilを返す  
    (サブクラスに)ダウンキャストできる
    ```swift
    式 as? 型
    ```
- `as!`  
    強制的キャスト。型変換が失敗した場合はエラーになる  
    キャストの効果は`as?`と同じ
    ```swift
    式 as! 型
    ```
