# 列挙型

## 列挙型とは

列挙型(enum)は複数の値を型として定義する  
switch文を使って場合分け処理をできる

### 定義

```swift
enum 名前 {
    case 値1
    case 値2
    case 値3
}
enum 名前 {
    case 値1, 値2, 値3
}
```
定義場所：
- クラス定義の外：グローバルな列挙型となる
- クラス定義の中：そのクラス内でのみ有効
- メソッド定義の中：そのメソッド内でのみ有効

### Enum ValueとRaw Value

列挙型の各値はEnum Valueという値である。  
値の型を指定すれば、Enum Valueに整数や文字列などのRaw Valueを代入できる  
Int型に指定した場合、設定値を省略すると、自動的に0からカウントアップされる。  
最初の値を1にすると、次は2, 3 ......  
連番である必要がなく、途中から値を変更できる  
==指定できるのは、整数、実数、論理値、文字列リテラルのみ==
```swift
enum 名前:型 {
    case 値1 = 値
    case 値2
    case 値3 = 値
}
```
Raw Valueを取り出す：
```swift
let ev = 列挙型名.Enum Value
print(ev.rawValue)
```
また、Raw ValueからEnum Valueを割り当てることもできる  
この時、オプショナル型となる
```swift
let a = 列挙型名(rawValue: 値)
if let b = a { print(b) }
```

### 列挙型の値を代入

```swift
enum MensSize {
    case S, M, L, XL
}

var mySize = MensSize.L
// 変数mySizeは型推論で、MensSize型となる
// 型を確定したら、次の代入は列挙型名を省略できる
mySize = .XL
```

### 全てのケースを取り出す

列挙型の定義で、`CaseIterable`プロトコルの採用を指定するだけでOK
```swift
enum Direction: Int, CaseIterable { ... }
for elem in Direction.allCases { ... }
```

### enumで型を列挙する

enumの中でenumを定義できる
```swift
enum Pattern {
    case monotone(_:PColor)
    case border(color1:PColor, color2:PColor)
    case dots(base:PColor, dot1:PColor, dot2:PColor)

    enum PColor:String {
        case red = "赤"
        case green = "緑"
        case yellow = "黄"
        case white = "白"
    }
}
```
Pattern型の値を作成：
```swift
// 赤シャツ
let shirt1 = Pattern.monotone(.red)
// 白赤ボーダーシャツ
let shirt2 = Pattern.border(color1:.white, color2:.red)
// 黄色地に白緑ドットのシャツ
let shirt3 = Pattern.dots(base:.yellow, dot1:.white, dot:.green)
```
switch文で値を取り出す
```swift
switch shirt1 {
    case .monotone(let color):                print(color)
    case .border(let color1, let color2):     print(color1, color2)
    case .dots(let base, let dot1, let dot2): print(base, dot1, dot2)
}
```

### 列挙型にプロパティ＆メソッドを実装

列挙型の中で定義できるプロパティは、計算型のみ
```swift
enum 名前 {
    case 値1, 値2, 値3

    static let 定数名:型 = 初期値
    static var 変数名:型 = 初期値
    var 変数名:型 {   get{return 値}   set(引数){}   }

    static func メソッド名(引数名:型) -> 戻り値の型 {
        // ...
        return 戻り値
    }
    func メソッド名(引数名:型) -> 戻り値の型 {
        // ...
        return 戻り値
    }

}
```
プロパティの定義例：
```swift
enum Ticket {
    case Gold, A, B
    static var name = "入場券"

    var area:String {
        get {
            switch self {
                case .Gold:
                    return "ゴールド席"
                case .A:
                    return "A席"
                case .B:
                    return "B席"
            }
        }
    }

}

Ticket.name = "ライブ入場券"
let ticket = Ticket.A
print(Ticket.name, ticket.area)
```
メソッドの定義例：
```swift
enum Signal:String {
    case Green = "緑色"
    case Red = "赤色"

    var color:String { return self.rawValue }

    static func description() -> String {
        return "GreenまたはRedのシグナルです"
    }

    // 列挙型の自分の値を変更するには、mutatingをつける
    mutating func turn(){
        if self==.Green{
            self = .Red
        } else {
            self = .Green
        }
    }

}
```
