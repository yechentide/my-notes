# 状態とデータフロー

## @State

### Stateの定義

`State`構造体はプロパティラッパとして定義されている。
```swift
@frozen @propertyWrapper struct State<Value>: DynamicProperty {
    init(wrappedValue: Value) {
        // SwiftUIが管理しているメモリ領域に値を格納する処理
    }
    var wrappedValue: Value {
        get {
            return /*SwiftUIが管理しているメモリ領域から値を取得する処理*/
        }
        nonmutating set {
            // SwiftUIが管理しているメモリ領域の値を更新する処理
        }
    }
    var projectedValue: Binding<Value> {
        get {
            return /*Bingingオブジェクトを返す処理*/
        }
    }
}
```

### Stateの仕組み

`State`を利用するには、プロパティラッパの`init()`を直接使うのではなく、  
`@State var count: Int = 0`のように、プロパティの前につける。  
これを「プロパティを`State`でラップする」という。  
このプロパティのことを、**状態変数**という。  
`@State`属性をつけたプロパティ`count`は、次のように展開される:
```swift
// @State var count: Int = 0 の裏側:
var count: Int {
    get { _count.wrappedValue }
    set { _count.wrappedValue = newValue }
}
var _count: State<Int> = State.init(wrappedValue: 0)
```
つまり、実際には`_count`がSwiftUI側で管理されている状態変数で、`count`はそれを介して値を取得・設定している。

### 射影プロパティ(projection property)

`@State`をつけたプロパティを、`Button`などの他のビューと関連づけるために、`$`をつけて渡すのは、射影プロパティ機能を利用している。  
射影プロパティは`projectedValue`という名前で定義されている。その値は`$状態変数名`で参照できるようになっている。  
`projectedValue`の戻り値は自由に設定できるが、`State`構造体の場合は`Binding<Value>`となる。  
このため、状態変数に`$`をつけるだけで、`State`が裏で管理している値にアクセスし、  
状態変数を定義したビューと他のビューとの間に双方向のデータの接続を可能にするためのBindingオブジェクトが得られる。

### ビューの更新

`@State`をつけた状態変数が変更されると、ビューも更新される。  
この機能を提供しているのが、`State`構造体が採用している`DynamicProperty`プロトコルである。  
`DynamicProperty`プロトコルを採用すると、`update()`メソッドを実装し、プロパティに変更があれば`update()`が呼ばれる。

### Stateの利用

- `State`はビューに紐づいて管理され、ビューの外(bodyプロパティの外)から変更してはいけない。  
    そのため、`State`変数を`private`にするのが望ましい。
- `State`の最大の役割はビューと状態変数を双方向で同期することである。  
    つまり状態変数が変わったらビューもかわり、ビューが変わったら状態変数も変わる。
- `@State`を先頭につけたプロパティは、swiftUIによって、そのメモリ領域が管理され、変更も監視される  
- ビューに渡す時は、`$`を先頭につけて渡す
```swift
struct ContentView: View {
    @State private var text: String = ""

    var body: some View {
        TextField("", $text)
        Text(self.text)
    }
}
```

## @Binding

### Bindingの定義

`Binding`プロパティラッパもSwiftUIのデータを管理する仕組みの一つであり、  
状態変数を定義したビューと他のビューとの間に双方向のデータの接続を作成する。  
実際の値は他のビューで定義され、共有されるという動作になる。  
また、`Binding`構造体にも射影プロパティ`projectedValue`が定義されているため、`$`をつけて他のビューに参照を渡すことができる。  
標準のコントロールも、`@Binding`を使って定義されている

### Bindingの利用

```swift
struct MyControllerView: View {
    @Binding var x: Double
    @Binding var y: Double
    var body: some View {
        VStack {
            Text( String(format: "%.0f, %.0f", $x.wrappedValue, $y.wrappedValue) )
            Slider(value: $x, in: 0.0...100.0)
            Slider(value: $y, in: 0.0...100.0)
        }
        .padding()
    }
}

struct ContentView: View {
    @State private var xValue: Double = 0.0
    @State private var yValue: Double = 0.0
    var body: some View {
        MyControllerView(x: $xValue, y:$yValue)
    }
}
```

### constantメソッド

引数を指定した値を持つ`Binding`を作れる
```swift
struct MyControllerView_Previews: PreviewProvider {
    static var previews: some View {
        MyControllerView(x: .constant(1.0), y: .constant(2.0))
    }
}
```

## @ObservedObject

### ObservedObjectの定義

`ObservedObject`構造体にも射影プロパティ`projectedValue`が定義されているため、`$`をつけて他のビューに参照を渡すことができる。  
ただし、`$オブジェクト名`だけでは、`Wrapper`を取得できるが`Binding`オブジェクトを取得できない。
- OK: `$オブジェクト名.Publishedプロパティ`   -> Bindingオブジェクト
- NG: `$オブジェクト名`                     -> Wrapperオブジェクト
- NG: `オブジェクト名.$Publishedプロパティ`   -> PublishedプロパティのprojectedValueメソッドの戻り値

### ObservedObjectの利用

`State`のように一時的な値ではなく、アプリが管理する独自のオブジェクトをバインディングすることもできる  
そのようなオブジェクトの先頭には、`@ObservedObject`をつける  
また、このオブジェクトを、`ObservableObject`プロトコルに適合させる必要がある  
オブジェクト内の、変更を通知するプロパティに、`@Published`アノテーションをつける  
`ObservableObject`プロトコルは、`Combine`フレームワークで定義されている。
- データを保存するクラスで`ObservableObject`プロトコルを採用する
- ビューの中で`@ObservedObject`属性を付加したプロパティを定義し、`ObservableObject`インスタンスを保存する
- データを保存するクラスでは、監視対象となるプロパティに`@Published`をつける  
    これでこのプロパティが変更された時に、SwiftUIのビューに対して更新通知が送られる
```swift
import Combine
class UserAccount: ObservableObject {
    @Published var userName = ""
    @Published var email = ""
    @Published var password = ""
}

struct ContentView: View {
    @ObservedObject var user: UserAccount

    var body: some View {
        VStack {
            TextField("name", $text)
            TextField("email", $text)
            TextField("pass", $text)
            Text(user.userName)
            Text(user.email)
            Text(user.password)
        }
    }
}
```

### 状態更新を手動で送る

単純にプロパティの値が更新されたことをSwiftUIに送信するだけなら、`@Published`でいいだが、  
特別な制御が必要な場合、`Combine`フレームワークの`ObservableObjectPublisher`クラスを使うことで、状態の更新を手動で通知できる。
```swift
import SwiftUI
import Combine

class MyData: ObservableObject {
    // @Published var score = 0 と同じ動作
    let objectWillChange = ObservableObjectPublisher()
    var score = 0 {
        willSet {
            objectWillChange.send()
        }
    }
}
```

## @EnvironmentObject

`ObservableObject`プロトコルに適合するオブジェクトを、  
ビュー階層の最上位のビューの`.environmentObject()`メソッドで設定すると、  
このオブジェクトはビュー全体で共有される  
このようなオブジェクトの先頭に、`@EnvironmentObject`をつける
- データを保存するクラスで`ObservableObject`プロトコルを採用する
- データを使うビューのうち、出発点となるビューで保存したいデータを`.environmentObject()`モディファイアに渡す
- ビューの中で`@EnvironmentObject`属性を使用してビューのプロパティを定義し、  
    このプロパティを経由して`.environmentObject()`モディファイアに渡したオブジェクトのインスタンスにアクセスする
```swift
import Combine
class AppData: ObservableObject {
    @Published var favoriteColor: Color = Color.black
}

struct TextMessageView: View {
    @EnvironmentObject var appData: AppData     // 上層のビューから引き継ぐ
    var body: some View {
        Text("Hello")
            .foregroundColor(self.appData.favoriteColor)
    }
}
struct TextMessageView_Previews: PreviewProvider {
    static var previews: some View {
        TextMessageView()
    }
}

struct ContentView: View {
    @EnvironmentObject var appData: AppData     // .environmentObject()メソッドで設定される
    var body: some View {
        TextMessageView()
    }
}
struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
            .environmentObject(AppData())
    }
}

// SceneDelegate.swift
class SceneDelegate: UIResponder, UIWindowSceneDelegate {
    // ....
    let contentView = ContentView().environmentObject(AppData())
    // ....
}
```

### @StateObject

```swift
@main
struct LandmarksApp: App {
    @StateObject private var modelData = ModelData()

    var body: some Scene {
        WindowGroup {
            ContentView()
                .environmentObject(modelData)
        }
    }
}
```

## 使い分け

### Stateを使う時

SwiftUIで管理され、他のビューにアクセスされない、一時的な値の代入
- ビューが編集モードか閲覧モードか。  
    切り替えるボタンもビューの中にあり、他のビューには編集状態を影響させない時
- リスト表示のフィルタリングの設定

### ObservedObjectを使う時

アプリ側で管理、型は独自定義
- オブジェクトをアプリが管理する時
- 複数のビューで共有して使用するオブジェクトをバインディングさせる時

### EnvironmentObjectを使う時

`EnvironmentObject`は`ObservedObject`の特殊ケース  
アプリ側で管理、型は独自定義
- アプリ全体でオブジェクトを使用する時
- ビューの階層をまたがって、オブジェクトを渡したい時

## Environment

SwiftUIでは、全体に関わる値(環境値)を、`.environment(キーパス, 値)`で指定できる  
環境値は`EnvironmentValues`構造体のプロパティに格納されていて、そのプロパティ名をキーパスで指定する  
例: `.environment(\.locale, Locale(identifier: ja_JP))`

### EnvironmentValuesの主なプロパティ

`environment()`メソッドで値を設定できるものもあれば、取得のみのプロパティもある
- calendar
- locale
- colorScheme
- editMode

### 使い方

ビューで環境値を参照したい場合、`Environment`プロパティラッパを使う。  
`@EnvironmentObject`と似ているが、`@EnvironmentObject`が任意のオブジェクトに渡せるに対して、  
`@Environment`は主にOSで事前に定義された値を取得・操作するためのものである。
```swift
struct ContentView: View {
    @Environment(\.accessibilityEnabled) var accessibilityEnabled
    var body: some View {
        Text("accessibilityEnabled: \(accessibilityEnabled ? "True" : "False")")
    }
}
struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
            .environment(\.accessibilityEnabled, true)
    }
}
```
