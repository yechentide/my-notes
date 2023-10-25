# List

UIKitの`UITableView`と似ているが、非常に少ないコードで実装できる。

## 静的リスト

`List`の内部にそのままビューの定義を記述するだけで、リストを生成できる。  
カスタムビューで1行のデザインをして、そのカスタムビューを`List`内で書くのもよし。
```swift
List {
    Text("SwiftUI Programming")
    Image(systemName: "sun.max")
    HStack {
        Image(systemName: "moon.stars")
        Text("moon.stars")
    }
}
```

## 動的リスト

WebやDBから取得した情報を元にリストを表示する時、検索結果を表示する時など、  
リストに表示する項目数が動的に変化する時には`ForEach`と組み合わせる。  
表示項目が変化しない時は、`ForEach`を使わずに、`List`のイニシャライザにコレクションを渡す
```swift
// items: [Item]
List(landmarks, id: \.id) { landmark in
    LandmarkRow(landmark: landmark)
}
List {
    ForEach(items) { item in
        if ... {
            MyCellView(item: item)
        }
    }
}
```
リストだけなら、`List`を使えばよい  
別の行も入れる場合、`List`と`ForEach`を合わせて使う

## ForEach構造体

`id`引数には、コレクション内の要素を識別する情報への`キーパス`を指定する。  
`String`や`Int`など単純な値であれば、それ自体が識別子にできるため、`\.self`で指定する。  
自分で定義した構造体などでは、`\UserData.userID`のように指定する。  
さらに、構造体などが`Identifiable`プロトコルに適合していれば、`id`引数を省略できる。  
代わりに構造体内部で`id`プロパティを定義する必要があるが、`id=UUID()`で一意なInt型の値を用意してくれる。
```swift
var strArray: [String] = []
List {
    ForEach(strArray, id: \.self) { str in
        Text(str)
    }
}
```

## セクションを入れる

セクションは、Listビュー内の項目を区切るもの
```swift
List {
    Section(header: Text("Title01")) {
        Text("A")
        Text("B")
    }
    Section(header: Text("Title02"), footer: Text("END")) {
        Text("A")
        Text("B")
    }
    Section(header: HStack {
        Image(systemName: "person.crop.circle")
        Text("Account Data")
    }) {
        Text("aaa")
        Text("bbb")
        Text("ccc")
    }
    Section() {
        Text("ddd")
    }
}
```

## Listに行を追加

`ForEach`でリストを生成する場合、データとなるオブジェクトに`@State`をつけて、  
それにデータを挿入するだけで、リストに行を追加できる

## スワイプして行を削除

`ForEach`の`onDelete()`メソッドを実装すれば、スワイプして削除する機能を実現できる。
```swift
List {
    ForEach(array, id: \.self) { item in
        Text(item)
    }
    .onDelete { offsets in
        self.array.remove(atOffsets: offsets)
    }
}
```

## 編集モード

`List`には通常モードと編集モードがある。  
編集モードでは編集用の機能を使えるようになる。  
モードの切り替えは`EditButton`ビューを使う

### セルの並べ替え

**編集モードのとき**にセルの並べ替えを可能にするには、  
`ForEach`の`onMove`モディファイアを使う。  
`onMove`モディファイアでセルが移動された時の処理を指定する  
`onMove`モディファイアに指定する処理は、クロージャーorメソッドである
```swift
struct ContentView: View {
    @State var languages: [String] = []
    var body: some View {
        NavigationView {
            List {
                ForEach(languages, id: \.self) { lang in
                    Text(lang)
                }
                .onMove(perform: move)
            }
            .navigationBarTitle("Languages")
            .navigationBarItems(trailing: EditButton())
        }
    }
}

func move(srcIndexes: IndexSet, dstIndex: Int) {
    // 移動された時の処理
    if let i = srcIndexes.first {
        let s = languages[i]
        languages.remove(at: i)
        if i < dstIndex {
            languages.insert(s, at: dstIndex-1)
        } else {
            languages.insert(s, at: dstIndex)
        }
    }
}
```

### セルの削除

編集モードのときにセルの並べ替えを可能にするには、  
`ForEach`の`onDelete`モディファイアを使う。  
`onDelete`モディファイアでセルが削除された時の処理を指定する  
`onDelete`モディファイアに指定する処理は、クロージャーorメソッドである
```swift
struct ContentView: View {
    @State var languages: [String] = []
    var body: some View {
        NavigationView {
            List {
                ForEach(languages, id: \.self) { lang in
                    Text(lang)
                }
                .onDelete(perform: deleteItem)
            }
            .navigationBarTitle("Languages")
            .navigationBarItems(trailing: EditButton())
        }
    }
}

func deleteItem(indexes: IndexSet) {
    // 移動された時の処理
    for index in indexes.reversed() {
        languages.remove(at: index)
    }
}
```

### 編集モードの取得

`List`だけでなく、どちらのモードになっているを知りたい時、環境情報から調べる
```swift
@Environment(\.editMode) var mode

if mode?.wrappedValue == .active {
    // 編集モードである
} else {
    // 通常モードである
}
```
