# コンテナビュー

## Group

他のビューを内包して、グループ化できるコンテナビューの１つで、画面には影響しない。  
Stackビューと同じように、`Group`の`{}`の中でビューの定義を並べて書ける。

## VStack / HStack / ZStack

Stackビューのイニシャライザの最後の引数`content`には、`@ViewBuilder`属性が付加されている。  
これ(関数ビルダ機能)により、Stackビューの中にビューの定義を並べて記述できるようになる。  
ただし、10個のビューを受け取るメソッドまでしか定義されてないので、**11個以上はエラーになる**
- `HStack`：水平方向に並べる
- `VStack`：垂直方向に並べる
- `ZStack`：重ねて表示
```swift
struct ContentView: View {
    var body: some View {
        VStack {
            Text("A")
            Text("B")
            Text("C")
            HStack(alignment: .top, spacing: 10) { /*......*/ }
            VStack(alignment: .leading, spacing: 10) { /*......*/ }
        }
    }
}
```
２つの引数、`alignment`と`spacing`がある。どちらも省略可能。  
`spacing`で子ビュー同士の間隔を設定できる。`ZStack`ではこの引数を使えない。  
`alignment`で配置位置を指定できる。
- HStack
    `.top`, `.center`, `.bottom`, `.firstTextBaseline`, `.lastTextBaseline`
- VStack
    `.leading`, `.center`, `.trailing`
- ZStack
    `.topLeading`, `.top`, `.topTrailing`,  
    `.leading`, `.center`, `.trailing`,  
    `.bottom`, `.bottom`, `.bottomTrailing`

## ScrollView

引数でスクロール可能な方向を指定する  
配列表示で指定する必要がある
```swift
ScrollView([.horizontal, .vertical]) {
    // 表示するコンテンツ
}
```

## Form

入力フォームを作るためのコンテナビューである  
NavigationViewと組み合わせて使うことを想定して設計されている  
SwiftUIのいくつかのビューは、`Form`内に配置されると自動的に適応し、  
入力領域のサイズ調整や画面遷移を行って、快適に入力できるように動作が変わる  
また、`List`と同じく、`Form`内でも`Section`ビューを使える
```swift
struct ContentView: View {

    @State private var textField = ""
    @State private var secureField = ""
    @State private var toggleState = false
    @State private var pickerSelected = 0

    var body: some View {

        NavigationView {
            Form {
                Text("This is a text")
                TextField("Text Field", text: $textField)
                SecureField("Secure Field", text: $secureField)
                Toggle(isOn: $toggleState, label: {
                    Text("Toggle")
                })
                Picker(selection: $pickerSelected, label: Text("Picker"), content: {
                    Text("Item 1").tag(0)
                    Text("Item 2").tag(1)
                    Text("Item 3").tag(2)
                })
            }
        }

    }
}
```

## Sheet

シートはビューの上に重なって表示されるビューのこと。  
SwiftUIではシートも他のビューと同じビューである。  
`sheet`モディファイアを使って、あるビューをシートとして表示させる。  
ある機能が現在表示中のビューとは独立に動作し、そのビューの中で完結したり、  
その結果を、シートを表示する前に表示していたビューに伝える時などに使う
```swift
struct ContentView: View {
    @State private var isShowingSheet: Bool = false
    var body: some View {
        Button(action: {
            self.isShowingSheet = true
        }) {
            Text("Open the Sheet")
        }
        .sheet(isPresented: $isShowingSheet) {
            MySheetView()
        }
    }
}
```

### sheetを閉じる

シートを閉じるには、`.sheet()`モディファイアの`isPresented`引数として渡されたものの値をfalseにすれば良い。  
つまり、表示するシート内で、その値を変更できるボタンなどを配置すれば良い  
あるいは、`NavigationView`のように`self.presentationMode.wrappedValue.dismiss()`を使ってもよい  
また、iOS13から、シートを上から下へスワイプしても閉じることができる  
シートを閉じた時の処理は、`.sheet()`モディファイアの`onDismiss`引数で記述する

## TabView

`TabView`は複数のビューをタブを使って切り替えるコンテナビューである。

### TabViewの基本形

タブバーに表示するものを、`.tabItem()`モディファイアで指定する
```swift
TabView {
    MyView01()
        .tabItem {
            Image()
            Text("Label1")
        }
    MyView02()
        .tabItem {
            Image()
            Text("Label2")
        }
}
```

### タブバーの表示を切り替える

`TabView`は通常、一番上に記述したタブが選択された状態で表示されるが、任意のタブをデフォルトにすることもできる
1. それぞれのタブで表示するビューに`.tag()`で番号を振る
2. 選択中のタグの番号を保持する状態変数を用意する
3. 状態変数と`TabView`をバインドする
```swift
struct ContentView: View {
    @State private var selection = 1
    var body: some View {
        TabView(selection: $selection) {
            View()
                .tabItem{
                    // ...
                }.tag(0)
            View()
                .tabItem{
                    // ...
                }.tag(1)
        }
    }
}
```

## カスタムビュー

デベロッパーが独自に作成するSwiftUIのビューは全てカスタムビューであり、  
プロジェクトを作成する時からある`ContentView`もカスタムビューである。  
`command`キーを押しながら部品をクリックすると、カスタムビュー(SubView)として抽出できる。  
これでコードの階層を浅くできる。
```swift
struct ContentView: View {
    var body: some View {
        VStack {
            subView01()
            subView02()
        }
    }
}
struct subView01: View {
    var body: some View {
        HStack {
            Image(systemName: "trash")
            Image(systemName: "trash")
        }
    }
}

struct subView02: View {
    var body: some View {
        HStack {
            Text("aaa")
            Text("bbb")
        }
    }
}
```

## GeometryReader

`GeometryReader`は、`Path`などを使ったシェイプのコンテナビューである  
親ビューに合わせて柔軟に大きさが変化して描画できる
```swift
struct ContentView: View {
    var body: some View {
        GeometryReader { geometry in
            Path { path in
                // Pathの定義
            }
            .fill(Color.blue)

            Path { path in
                // Pathの定義
            }
            .stroke(lineWidth: 20)
            .foregroundColor(.red)
        }
    }
}
```
