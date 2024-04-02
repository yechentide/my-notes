# APIデザインガイドライン

[API Design Guidelines](https://www.swift.org/documentation/api-design-guidelines/)

## 命名

### 曖昧さを減らすために必要な言葉を入れる

```swift
// ⭕️
employees.remove(at: x)
// ❌ 「x」を消したいのか、「x番目」の要素を消したいのか
employees.remove(x)
```

### 必要のない言葉を省略する

```swift
// ⭕️
allViews.remove(cancelButton)
// ❌
allViews.removeElement(cancelButton)
```

### 型ではなく役割で変数などの名前を決める

```swift
// ⭕️
var greeting = "Hello"
protocol ViewController {
  associatedtype ContentView : View
}
class ProductionLine {
  func restock(from supplier: WidgetFactory)
}
// ❌
var string = "Hello"
protocol ViewController {
  associatedtype ViewType : View
}
class ProductionLine {
  func restock(from widgetFactory: WidgetFactory)
}
```

### `NSObject`,`Any`,`AnyObject`や`String`,`Int`など型から役割を特定できない場合、補足情報を与える

```swift
// ⭕️
func addObserver(_ observer: NSObject, forKeyPath path: String)
grid.addObserver(self, forKeyPath: graphics)
// ❌
func add(_ observer: NSObject, for keyPath: String)
grid.add(self, for: graphics)
```

### 英語のフレーズを優先的に使う

```swift
// ⭕️
x.insert(y, at: z)          “x, insert y at z”
x.subViews(havingColor: y)  “x's subviews having color y”
x.capitalizingNouns()       “x, capitalizing nouns”
// ❌
x.insert(y, position: z)
x.subViews(color: y)
x.nounCapitalize()
```

### factory系のメソッド名はmakeから始まる

例: `x.makeIterator()`

### 副作用の有無でメソッド名を決める

- 副作用がないメソッドの名前は名詞にすべき: `x.distance(to: y)`, `i.successor()`
- 副作用があるメソッドの名前は動詞にすべき: `print(x)`, `x.sort()`, `x.append(y)`
- mutating/non-mutatingメソッドはペアで定義する
    - 動作は動詞で表現される場合
        - `x.sort()`, `z = x.sorted()`
        - `x.append(y)`, `z = x.appending(y)`
    - 動作は名詞で表現される場合
        - `y.formUnion(z)`, `x = y.union(z)`
        - `c.formSuccessor(&i)`, `j = c.successor(i)`

### アサーションとして真理値のデータを命名する

例: `x.isEmpty`, `line1.intersects(line2)`

### プロトコルの命名

- 「どんなものか」を表す際は名詞を使う
    - Collection
- 「何ができるか」を表す際は接尾辞`able`, `ible`, `ing`を使う
    - Equatable, ProgressReporting
- 他の型、プロパティ、定数、変数は全て名詞で表す

### 慣用語を使おう

- わかりにくい言葉を避ける
    - ⭕️ skin
    - ❌ epidermis
- 専門用語を使う場合、その意味を正しく理解していることが重要
- 略称を避ける
- 勝手に名前の最適化せず、既存名を尊重して受け入れる
    - `List`は初心者にはわかりやすいが、`Array`が良い
    - 完全な単語の`sine`よりも、`sin()`の方が良い

## 慣習

### 計算量がO(1)以上の場合に注意書きをする

そうしないと、重い計算だと知らずに多用されてしまうリスクがある

### トップレベル関数よりもメソッドやプロパティを使う

トップレベル関数は限られたケース内でしか利用されない
1. 明確な`self`がない: `min(x, y, z)`
2. 制約がないような汎用関数: `print(x)`
3. すでに確立された表示の一部: `sin(x)`

### CamelCaseを使う

- 型やプロトコルは`UpperCamelCase`
- それ以外は`lowerCamelCase`
- 大文字の略語は、そのまま使うか、全部小文字に直すべき

### 同じ役割を持つメソッドや違うドメインのメソッドは同じ名前にしても良い

```swift
// ⭕️
extension Shape {
  func contains(_ other: Point) -> Bool { ... }
  func contains(_ other: Shape) -> Bool { ... }
  func contains(_ other: LineSegment) -> Bool { ... }
}
// ShapeとCollectionは違うドメイン
extension Collection where Element : Equatable {
  func contains(_ sought: Element) -> Bool { ... }
}
// ❌
extension Database {
  func index() { ... }
  func index(_ n: Int, inTable: TableID) -> TableRow { ... }
}
extension Box {
  func value() -> Int? { ... }
  func value() -> String? { ... }
}
```

### 関数の引数はわかりやすいものにする

```swift
// ⭕️
func filter(_ predicate: (Element) -> Bool) -> [Generator.Element]
mutating func replaceRange(_ subRange: Range, with newElements: [E])
// ❌
func filter(_ includedInResult: (Element) -> Bool) -> [Generator.Element]
mutating func replaceRange(_ r: Range, with: [E])
```

### よく使われる値があるならデフォルト引数にする

デフォルト引数は後ろに回すべき
```swift
// ⭕️
let order = lastName.compare(royalFamilyName)
extension String {
  public func compare(
     _ other: String, options: CompareOptions = [],
     range: Range? = nil, locale: Locale? = nil
  ) -> Ordering
}
// ❌
let order = lastName.compare(royalFamilyName, options: [], range: nil, locale: nil)
extension String {
  public func compare(_ other: String) -> Ordering
  public func compare(_ other: String, options: CompareOptions) -> Ordering
  public func compare(
     _ other: String, options: CompareOptions, range: Range) -> Ordering
  public func compare(
     _ other: String, options: StringCompareOptions,
     range: Range, locale: Locale) -> Ordering
}
```

### 区別しにくい場合、引数ラベルを全て省略する

```swift
min(number1, number2)
zip(sequence1, sequence2)
```

### データを変換するためのinitializerの第１引数ラベルを省略する

そして第１引数は常に変換元である

### 第１引数が前置詞句の一部となる場合、ラベルを指定する

```swift
// ⭕️
a.moveTo(x: b, y: c)
a.fadeFrom(red: b, green: c, blue: d)
// ❌
a.move(toX: b, y: c)
a.fade(fromRed: b, green: c, blue: d)
```

### 第１引数が文法的な語句の一部となる場合、ラベルから関数名に移動する

```swift
// ⭕️
view.dismiss(animated: false)
let text = words.split(maxSplits: 12)
let studentsByName = students.sorted(isOrderedBefore: Student.namePrecedes)
// ❌
view.dismiss(false)   // Don't dismiss? Dismiss a Bool?
words.split(12)       // Split the number 12?
```

### タプル内の値やクロージャの引数にラベルを指定する

```swift
mutating func ensureUniqueStorage(
  minimumCapacity requestedCapacity: Int,
  allocate: (_ byteCount: Int) -> UnsafePointer<Void>
) -> (reallocated: Bool, capacityChanged: Bool)
```
