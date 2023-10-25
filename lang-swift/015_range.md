# 範囲型＆ストライド型

## 範囲型

### 範囲を表す型

|            範囲型             |   区間   | カウント可能 |
| :---------------------------: | :------: | :----------: |
|        `Range<Bound>`         | 半開区間 |      X       |
|    `CountableRange<Bound>`    | 半開区間 |      ◯       |
|     `ClosedRange<Bound>`      |  閉区間  |      X       |
| `CountableClosedRange<Bound>` |  閉区間  |      ◯       |

`Array<Element>`型のElementのように、Boundも**プレースホルダ型**である  
カウント可能なものは、for-in文でその要素に順次アクセスできる

### 範囲型の種類

範囲を表現するため、`<T>`は比較可能な型、つまりプロトコルComparableに準拠する必要がある  
上限値or下限値のどっちかしかないものは片側範囲という  
for-inで使うために、プロトコルSequenceに準拠しなければならない  
プロトコルCollectionはプロトコルSequenceを継承したものなので、これに準拠したものでもfor-inで使える

|   例    | 型                         | プロトコル             |
| :-----: | :------------------------- | :--------------------- |
| `A..<B` | `Range<Int>`               | Comparable, Collection |
| `A..<B` | `Range<Double>`            | Comparable             |
| `A...B` | `ClosedRange<Int>`         | Comparable, Collection |
| `A...B` | `ClosedRange<Double>`      | Comparable             |
| `A...`  | `PartialRangeFrom<Int>`    | Comparable, Sequence   |
| `A...`  | `PartialRangeFrom<Double>` | Comparable             |
| `..<B`  | `PartialRangeUpTo<T>`      | Comparable             |
| `...B`  | `PartialRangeThrough<T>`   | Comparable             |

### 範囲演算子

- `..<`  
    半開区間を作る演算子  
    `1.0..<3.5`の場合、1.0<= x < 3.5と意味する  
    Intの場合はカウント可能だが、型アノテーションでカウント不可能の型の値を作れる  
    一方、Doubleは型アノテーションを使っても作れない
- `...`  
    閉区間を作る演算子  
    `1...3`の場合、1<= x <=3と意味する  
    Intの場合はカウント可能だが、型アノテーションでカウント不可能の型の値を作れる  
    一方、Doubleは型アノテーションを使っても作れない

### 範囲型のプロパティ

- `lowerBound`：下限値
- `upperBound`：上限値

### 範囲型のメソッド

- `constains()`：引数の値が範囲内の場合true
- `overlaps()`：引数の範囲と重複する場合true。片側範囲には使えない
- `shuffled()`：含まれる要素をランダムに並び替える。プロトコルSequenceに準拠するデータ型に使える。

### 乱数

- `Int.random(in: 範囲型)`
- `Double.random(in: 範囲型)`
- `Bool.random()`

## ストライド型

### 概要

開始点＆終了点＆刻み幅で構成される構造体である。  
`StrideThrough`と`StrideTo`の２つがある

### 例

```swift
for x in stride(from:0, to:10, by:2){
    print(x, terminator:" ")            // 0, 2, 4, 6, 8
}
for x in stride(from:0, through:10, by:2){
    print(x, terminator:" ")            // 0, 2, 4, 6, 8, 10
}
```

### ストライド型とプロトコルStrideable

ストライド型の正式な宣言はジェネリック関数
```swift
func stride<T>(from:T, to:T, by:T.Stride) -> StrideTo<T> where T:Strideable
```
`T`はパラメーターである。`T`はプロトコルStrideableに準拠する必要がある。  
`T.Stride`は型`T`の付属型である。  
プロトコルStrideableはプロトコルComparableを継承している
```swift
public protocol Strideable : Comparable {
    associatedtype Stride : Comparable, SignedNumeric
    func distance(to other: Self) -> Self.Stride
    func advanced(by n: Self.Stride) -> Self
}
```
