# NavigationView

## タイトル

displayModeの値は、`.large`、 `.inline`、 `.automatic`の３つ  
デフォルトは`.large`
```swift
struct ContentView: View {
    var body: some View {
        NavigationView {
            Text("aaa")
                .navigationBarTitle("Home")
                //.navigationBarTitle("Home", displayMode: .inline)
        }
    }
}
```

## ナビゲーションバーを非表示

```swift
NavigationView {
    View()
        .navigationBarTitle("Home")
        .navigationBarHidden(true)
}
```

## Navigationbaritems

ナビゲーションバーアイテムは、左右それぞれ指定できる。  
左側が`leading`、右側が`trailing`である。  
HStackで複数の表示もできる
```swift
struct MySubView: View {
    var body: some View {
        Text("SubView")
            .navigationBarItems(trailing: HStack {
                Button(action: {}) { Text("B01") }
                Button(action: {}) { Text("B02") }
            })
    }
}
```

## NavigationLink

```swift
struct MySubView: View {
    var prefix: String
    var index: Int
    var childCount: Int
    var displayText: String { return "\(prefix)(\(index))" }

    var body: some View {
        VStack {
            Text(displayText)
            if index < childCount - 1 {
                NavigationLink(
                    destination:
                        MySubView(prefix: prefix, index: index+1, childCount: childCount),
                    label: {Text("Go to child")}
                )
            }
        }
        .navigationBarTitle("\(displayText)", displayMode: .inline)
    }
}

struct ContentView: View {
    var body: some View {
        NavigationView {
            VStack {
                NavigationLink(
                    destination: MySubView(prefix: "A", index: 0, childCount: 3)) {
                    Text("A")
                }
                NavigationLink(
                    destination: MySubView(prefix: "A", index: 0, childCount: 1)) {
                    Text("B")
                }
                NavigationLink(
                    destination: MySubView(prefix: "A", index: 0, childCount: 5)) {
                    Text("C")
                }
            }
        }
    }
}
```

## 自分で前画面に戻る機能を用意する

遷移先のビューで、環境から`PresentationMode`構造体を取得し、`dismiss()`メソッドを実行すれば良い。  
`PresentationMode`は、現在別のビューで表示されているかどうかを示すモードを保持する構造体であり、  
この構造体のインスタンスは、`EnvironmentValues`構造体で管理されている。  
`@Environment`プロパティラッパを利用する必要がある。
```swift
// メインビューからサブビューに遷移して、メインビューに戻る
struct SubView: View {
    @Environment(\.presentationMode) var presentationMode
    var body: some View {
        VStack {
            Text("Hello")
            Button(action: {
                self.presentationMode.wrappedValue.dismiss()
            }) {
                Text("戻る")
            }
        }
    }
}
```
`self.presentationMode`は`Binding<PresentationMode>`型なので、  
`wrappedValue`という計算型プロパティ経由で、`presentationMode`のインスタンスにアクセスする  
自分で戻る機能を用意したので、ナビゲーションバーの戻るボタンを非表示にしたければ、`.navigationBarBackButtonHidden()`モディファイアを使う
