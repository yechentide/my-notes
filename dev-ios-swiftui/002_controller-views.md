# 基本的なビュー

## Text

`Text`ビューには、文字数、フォントサイズ、フレームの設定により、表示サイズが自動的に調整されるという特性がある。  
`+`を使うことで、`Text`ビューを連結して、１つのビューにできる。  
また、`Text`は自動的にローカライズに対応するため、引数をそのまま表示させたい場合、`Text(verbatim:)`を使う
```swift
Text("🍌🍌")
Text(Date(), style: .date)
```

## Image

画像を表示するには`Image`ビューを使う。  
リソースハンドル、UIImage、NSImage、CGImage、システムアイコンなどから画像を読み込める。  
iOS13以降では、`San Franciscoシンボルセット`から1500以上のアイコンを利用できる。  
これらのアイコンは、`SF Symbols`アプリで確認できる。

## Color

色を指定するには`Color`ビューを利用する。  
`Color`ビューは`View`プロトコルを適合しているため、色指定以外に、単純にその色のビューとして使うのも可能。  
また、RGB色や、アセットカタログで定義された色で、`Color`ビューを生成することもできる。

## Spacer

`Spacer`は伸縮可能な空白を作るビューである  
このビューは、広がれるだけ広がるため、２つの部品の間に入れれば、２つの部品は両サイドに配置される

## Divider

Stackビューなどの中で、各項目を視覚的に区別できるように、項目の間に仕切りとして使うビュー。  
`Divider`はビューなので、モディファイアによりカスタマイズができるが、あまり変わらない。

## Button

```swift
Button("文字列ラベル", action: { /*タップされた時の動作*/ })
Button(action: { /*タップされた時の動作*/ }) {
    // ボタンに表示する内容
}
```

## Toggle

```swift
@State private var buttonState = true
Toggle(isOn: $buttonState) {
    Text("Remember password")
}
```

## TextField

現状では、`.textFieldStyle()`モディファイアで線の太さや色を変えられないので、`.overlay()`などを使うと良い
```swift
@State private var user = ""
TextField("Enter the username", text: $user)
```

## SecureField

```swift
@State private var pass = ""
SecureField("Enter the pass", text: $pass)
```

## Slider

```swift
@State private var sliderValue: Float = 0.5
Slider(value: $sliderValue, in: -100...100, step: 0.1)
```
引数:
- value: 関連づける状態変数を先頭に`$`をつけて指定
- in: Sliderの指定可能な範囲。ClosedRange型。
- step: 増減の変化量。デフォルトは1。
- onEditingChanged: ユーザがSliderを操作している時に呼ばれる関数orクロージャ。デフォルトは`{_in}`で何もしない。
- minimumValueLabel: Sliderの左側に表示したいビュー
- maximumValueLabel: Sliderの右側に表示したいビュー
- label: ラベルとして表示したいビュー。現状、iOSでは表示されないが、省略できないため、`{ EmptyView() }`で指定するといいかも。

## Picker

選択肢として、画像などを使っても良い。  
また、選択された項目を特定するために`tag()`メソッドを使うが、数字じゃなくても良い。例: `tag("雪")`とか  
`ForEach`構造体で選択肢を定義する場合、`tag()`メソッドはいらない。  
ユーザが項目を選択すると、選択された項目のインデックスが入る。
```swift
@State private var selectedTag = 0
Picker(selection: $selectedTag, label: Text("A Picker")) {
    Text("*").tag(0)
    Text("A").tag(1)
    Text("B").tag(2)
    Text("C").tag(3)
}
```
UIKitでは、選択肢の中から１つ選ぶ`UISegmentedControl`がある。  
SwiftUIでは、Pickerビューのスタイルを変更するだけで実現できるが、ラベルが表示されなくなる。
```swift
@State var index = 0
Picker(selection: $index, label: Text("picker")) {
    ForEach(0...3, id: \.self) { index in
        Text("\(index)")
    }
}
.pickerStyle(SegmentedPickerStyle())
.pickerStyle(.segmented)
```

## DatePicker

DatePickerのイニシャライザには`displayedComponents`という引数がある。  
デフォルトは`[.hourAndMinute, .date]`で、日付と時刻の両方を指定できる。  
引数`in`で、選択可能範囲を指定できる
```swift
@State private var date = Date()
DatePicker(selection: $date, label: {Text("Date Label")})
```

## Stepper

Stepperのイニシャライザの`onIncrement`, `onDecrement`, `onEditingChanged`引数にクロージャを渡せば、  
それぞれのタイミングで処理を行ってくれる。
```swift
@State private var stepperValue = 3
Stepper(value: $stepperValue, in: 1...9) {
    Text("Value: \(self.stepperValue)")
}
```

## Alert

`Alert`はOS標準のアラートを作るためのコンテナである。  
ボタンの個数や種類によって、使用するイニシャライザやボタンの作り方が少しずつ異なる。

### `Close`だけ

最も単純なパターンで、アラートを閉じるためのボタンだけを持つ  
`dismissButton`は、`default`、`cancel`、`destructive`を指定できる
```swift
struct ContentView: View {
    @State private var isShowingAlert: Bool = false
    var body: some View {
        Button(action: {
            self.isShowingAlert = true
        }) {
            Text("Show Alert")
        }

        .alert(isPresented: $isShowingSheet) {
            Alert(
                title: Text("Sample"),
                message: Text("SwiftUI Alert Message"),
                dismissButton: .default(Text("Close")))
        }
    }
}
```

### `Ok`と`Cancel`

ボタンを２つ持つアラート
```swift
struct ContentView: View {
    @State private var isShowingAlert: Bool = false
    @State private var actionName = ""
    var body: some View {
        VStack {
            Button(action: {
                self.isShowingAlert = true
            }) {
                Text("Show Alert")
            }

            .alert(isPresented: $isShowingSheet) {
                Alert(
                    title: Text("Sample"),
                    message: Text("Are you sure to execute?"),
                    primaryButton: .default( Text("OK"), action: {self.actionName="OK"} ),
                    secondaryButton: .cancel( Text("Cancel"), action: {self.actionName="Cancel"} ))
            }

            Text(actionName)
        }
    }
}
```

### `Delete`と`Cancel`

何かを削除したり、破壊したりする時に、専用のボタンを表示する。  
そのようなアラートのボタンは、`destructive`で指定する
```swift
struct ContentView: View {
    @State private var isShowingAlert: Bool = false
    @State private var actionName = ""
    var body: some View {
        VStack {
            Button(action: {
                self.isShowingAlert = true
            }) {
                Text("Show Alert")
            }

            .alert(isPresented: $isShowingSheet) {
                Alert(
                    title: Text("Sample"),
                    message: Text("Are you sure to delete?"),
                    primaryButton: .destructive( Text("Delete"), action: {self.actionName="Delete"} ),
                    secondaryButton: .cancel( Text("Cancel"), action: {self.actionName="Cancel"} ))
            }

            Text(actionName)
        }
    }
}
```

### オブジェクトをバインドしてアラートを表示

`isPresented`引数には、`Bool`以外のオブジェクトをバインドしても良い。  
その場合、nilのときはアラート非表示で、値が代入されnilでなくなったら、アラートが表示される。  
ただし、バインドされるオブジェクトは、`Identifiable`プロトコルに適合する必要がある。
> SwiftUI全体に言えることだが、SwiftUIではどのオブジェクトがどのデータを使っているかを監視して、  
> データが更新されたら、そのデータを使っているビューを更新しようとしている。  
> その際にSwiftUIが要求するデータが一意であることを保証する必要がある。

## ActionSheet

`ActionSheet`はOS標準のアクションシートを作るためのコンテナである。  
アクションシートには複数のボタンが表示され、ユーザは表示された選択肢の中から１つ選択する。  
ボタンの種類は`Alert`と同じだが、3つ以上のボタンを表示できる

### ActionSheetを表示する

```swift
struct ContentView: View {
    @State private var isShowing: Bool = false
    @State private var actionName = ""
    var body: some View {
        VStack {
            Button(action: {
                self.isShowing = true
            }) {
                Text("Show ActionSheet")
            }

            .actionSheet(isPresented: $isShowing) {
                ActionSheet(
                    title: Text("Sample"),
                    message: Text("Are you sure to delete?"),
                    buttons: [
                        .default( Text("Option"), action: {self.actionName="Option"} ))
                        .destructive( Text("Delete"), action: {self.actionName="Delete"} ),
                        .cancel( Text("Cancel"), action: {self.actionName="Cancel"} ))
                    ]
                )
            }

            Text(actionName)
        }
    }
}
```
