# イベント＆ジェスチャ

## ジェスチャ

SwiftUIでジェスチャによるインタラクションを実装するには、ビューにジェスチャ用のモディファイアを蒸し、実装する処理を指定する

### シングルタップ

```swift
View()
    .onTapGesture {
        // タップと判定された時の処理
    }
```

### 複数回タップ

```swift
View()
    .onTapGesture(count: 対応したい回数) {
        // 複数回タップと判定された時の処理
    }
```

### ロングタップ

`.onLongPressGesture()`の引数の中で、判定に使う最小期間、最大移動距離などがある
```swift
View()
    .onLongPressGesture(pressing: {press in /*押された時だけがtrue*/}) {
        // ロングタップと判定された時の処理
    }
```

### ジェスチャのカスタマイズ

`.gesture()`モディファイアを使って、ビューに対するカスタムジェスチャを定義できる。  
`.gesture()`の第１引数には、検知したいジェスチャとして、`Gesture`プロトコルを採用した構造体のインスタンスを渡す。
```swift
View()
    .gesture(
        TapGesture()
            .onChanged { _ in
                // 処理
            }
            .onEnded { _ in
                // 処理
            }
    )
```

### ドラッグ

```swift
View() {
    .gesture(
        DragGesture()
            .onChanged { value in
                // valueはDragGesture.Value型。time, location, startLocation, translationなどのプロパティがある。
                // 処理
            }
    )
}
```

## ライフイベント

### ビューの表示・非表示

```swift
View()
    .onAppear {
        // ビューが表示された時の処理
    }
    .onDisappear {
        // ビューが非表示になった時の処理
    }
```

### キーボード

`TextField`の位置によって、キーボードで隠れてしまう時がある。  
そのため、キーボードの高さだけビュー全体を上にずらすことがよく行われる。  
UIKitと同じく、SwiftUIも`NotificationCenter`を利用するが、  
処理は`ObservableObject`プロトコルを採用したクラスで行い、最終的には`@ObservedObject`を介してビューに伝える。
```swift
class KeyboardObserver: ObservableObject {
    @Published var keyboardHeight: CGFloat = 0.0

    func startObserve() {
        NotificationCenter.default.addObserver(
            self,
            selector: #selector(keyboardWillChangeFrame(_:)),
            name: UIResponder.keyboardWillChangeFrameNotification,
            object: nil)
    }

    func stopServe() {
        NotificationCenter.default.removeObserver(self, name: UIResponder.keyboardWillChangeFrameNotification, object: nil)
    }

    @objc func keyboardWillChangeFrame(_ notification: Notification) {
        if let keyboardEndFrame = notification.userInfo?[UIResponder.keyboardFrameEndUserInfoKey] as? NSValue,
           let keyboardBeginFrame = notification.userInfo?[UIResponder.keyboardFrameBeginUserInfoKey] as? NSValue {
            let endMinY = keyboardEndFrame.cgRectValue.minY
            let beginMinY = keyboardBeginFrame.cgRectValue.minY
            keyboardHeight = beginMinY - endMinY
            if keyboardHeight < 0 {
                keyboardHeight = 0
            }
        }
    }
}

struct ContentView: View {
    @ObservedObject var keyboard = KeyboardObserver()
    @State var inputText = ""

    var body: some View {
        VStack {
            Spacer()
            TextField("入力エリア", text: $inputText)
        }
        .onAppear {
            keyboard.startObserve()
        }
        .onDisappear {
            keyboard.stopServe()
        }
        .padding(.bottom, keyboard.keyboardHeight)
    }
}
```
