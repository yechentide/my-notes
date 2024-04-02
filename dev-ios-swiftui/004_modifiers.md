# モディファイア

## モディファイアについて

SwiftUIでは、各ビューが持つメソッドを使って、ビューのカスタマイズを行う。  
これらのメソッドを**モディファイア**という。  
モディファイアは、カスタマイズ後のビューを生成して返すので、複数のモディファイアを並べて記述できる。(モディファイアチェーン)  
たくさんのモディファイアを指定すると、ビューがたくさん生成され非効率になることはない。  
SwiftUIが最適化をしているため、無駄を省いて高速動作ができるようになっている。

### モディファイアを追加

**モディファイアの適用順序が変わると、効果が違う時もある。**
追加は3つのやり方がある:
1. コードで直接編集する
2. ポップアップメニューで編集する  
    `command`押しながら編集したいビューをクリック  
    `Show SwiftUI Inspector`を選択肢て編集
3. 右側のAttributes Inspectorで編集する

## 標準モディファイア

### Viewに対して

- `frame()`: ビューのサイズを設定
- `fixedSize()`: フレームサイズのはみ出しを可能にする
    - `frame()`よりも**先に実行**しなければならない
    - 隣接ビューと重なって表示されることに注意!
- `border(_:width:)`: ビューに枠線を設定する
- `position(x:y:)`: ビューの表示位置を座標で指定する
    - SwiftUIの座標系は、右がxの正、下がyの正
    - `position()`を指定すると、フレーム設定がリセットされ、強制的に親ビューの大きさになる
- `offset(x:y:)`: ビューの表示をずらす
- `padding(_:)`: 余白を指定
    - 引数を指定しない場合、システムデフォルトの値になる
    - 複数の方向を指定する時、配列で記述する
- `foregroundColor(_:)`
- `backgroundColor(_:)`
- `overlay(_:alignment:)`: 別のビューを重なって表示する
    - ビューを引数として渡す
    - 重ねるビューのサイズは、元になるビューのサイズに限定される
- `stroke(_:lineWidth:)`
- `clipped(antialiased:)`: 矩形で切り抜く
- `clipShape(_:style:)`: 引数で指定した形状で切り抜く  
    例: `.clipShape(Circle())`
- `shadow(color:radius:x:y:)`: 影をつける
- `mask(_:)`: 他のビューの透明度情報で切り抜く
- `cornerRadius(_:antialiased:)`: ビューを角丸にする

### Textに対して

- `foregroundColor(_:)`
- `font(_:)`  
    例: `.font(.system(size: 20, weight: .bold, design: .rounded))`
- `fontWeight(_:)`
- `bold()`
- `italic()`
- `strikethrough(_:color:)`
- `underline(_:color:)`
- `kerning(_:)`: 2文字間の間隔を設定
- `tracking(_:)`: 文字列全体の文字間隔を調整 (`kerning(_:)`より優先される)
- `baselineOffset(_:)`: ベースラインの位置を指定
- `lineSpacing(_:)`: 文字列の行間を調整する
- `truncationMode(_:)`: 表示スペースが足りない時の`...`の表示位置
- `lineLimit(_:)`: 文字列の表示行数を指定する

### Imageに対して

- `resizable(capInsets:resizingMode:)`: 画像をリサイズ可能にする  
    **画像サイズを変更する前に必ずこれを実行する**
- `scaledToFit()`
- `scaledToFill()`
- `renderingMode(_:)`
- `clipped()`

### Buttonに対して

- `buttonStyle(_:)`

### TextField / SecureField に対して

- `textFieldStyle(_:)`

### Stackに対して

- `alignmentGuide(_:computeValue:)`
```swift
// カスタムalignment
extension VerticalAlignment {
    private enum MyCustomAlign: AlignmentID {
        static func defaultValue(in context: ViewDimensions) -> CGFloat {
            context.height / 2
        }
    }
    public static let myCustomAlign = VerticalAlignment(MyCustomAlign.self)
}
```

### Listに対して

- `listStyle(_:)`  
    例: `.listStyle(.inset)`
- `listRowInsets(_:)`  
    例: `.listRowInsets(EdgeInsets())`

### NavigationViewに対して

- `navigationTitle(_:)`
    - `navigationBarTitle(_:)`: iOS13まではこれ
- `navigationBarTitleDisplayMode(_:)`

### プレビュー関連

- `.environment(\.colorScheme, .dark)`
- `.environment(\.locale, Locale(identifier: "ja_JP"))`  
    zh_Hans_CN, ja_JP, en_US
- `.environmentObject(環境オブジェクト)`
- `Group`: ビューで複数のプレビューを表示できる
- `.previewDevice("iPhone SE")`
- `.previewLayout(.fixed(width: 400, height: 100))`
- `previewDisplayName(_:)`
- `constant()`: プレビューで定数を渡す(Bind<???>)

### その他

- `contextMenu(_:)`
- `ignoresSafeArea(_:edges:)`
- `edgesIgnoringSafeArea(_:)`: セーフエリア外にコンテンツを配置する
    - 引数は７つあり、配列表記で複数指定可能
- `aspectRatio(_:contentMode:)`

## 独自モディファイア

### モディファイア作成の流れ

1. `ViewModifier`プロトコルを採用した構造体を用意する
2. `body`メソッドを実装する
3. `body`メソッドでは必要な処理をした後にビューを返す

### 自作モディファイアの例

```swift
struct MyModifier: ViewModifier {
    func body(content: Content) -> some View {
        return content
            .font(.largeTitle)
            .padding()
            // その他の設定
    }
}

struct ContentView: View {
    var body: some View {
        VStack {
            // 適用の仕方は２通り。１つ目を使えば問題ないし見やすい
            Text("aaa")
                .modifier(MyModifier())
            ModifiedContent(content: Text("bbb"), modifier: MyModifier())
        }
    }
}
```
また、`View`を拡張すればもっと簡単に書ける
```swift
extension View {
    func myMod() -> some View {
        self.modifier(MyModifier())
    }
}
struct ContentView: View {
    var body: some View {
        Text("ccc")
            .myMod()
    }
}
```

### 自作モディファイアのテクニック

`ViewModifier`で、ビューの構造を変えたり、処理を追加したりできる
```swift
struct MyModifier: ViewModifier {
    func body(content: Content) -> some View {
        HStack {
            content
            Text("語尾に追加するテキスト")
        }
    }
}
```
