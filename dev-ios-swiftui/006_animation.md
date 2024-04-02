# アニメーション

## ビューの変形

### ビューの位置を移動する

```swift
View()
    .position(x:y:)     // ビューの中心を設定
View()
    .offset(x:y:)       // x,y方向にずらす量を設定
```

### ビューを回転させる

```swift
// ビュー中心を原点にして回転
View()
    .rotationEffect(Angle(degrees:))
// anchor引数で、回転中心を変更
View()
    .rotationEffect(Angle(degrees:), anchor: .bottomTrailing)
```

### ビューを拡大・縮小させる

```swift
View()
    .scaleEffect(倍率)
View()
    .scaleEffect(x: 1.0, y: 1.0, anchor: .center)
```

## アニメーションを行う

### animation

デフォルトでは、モディファイアに渡される値が変わると、新しい値でレンダリングされ、それが完了したら、画面が書き換わる。  
`animation`モディファイアを使うと、新しい値までの中間値でもレンダリングされ、結果を描画される
```swift
View()
    .アニメーションを適用するモディファイア
    .animation(アニメーションの種類)
```
`animation`モディファイアで使う引数は:  
`default`, `easeIn`, `easeOut`, `easeInOut`, `linear`  
これらには全て時間を指定できる  
例: `.animation(.easeIn(duration: 3.0))`  
この5つ以外に、`spring()`というアニメーション関数がある

### withAnimation

```swift
withAnimation(アニメーションの種類) {
    // アニメーションの対象にする変更処理
}
```

### アニメーションの例

```swift
struct ContentView: View {
    @State private var ratio: CGFloat = 1
    @State private var degrees: Double = 0
    @State private var offsetY: CGFloat = 0

    var body: some View {
        VStack {
            Text("Hello SwiftUI!")
                .font(.largeTitle)
                .scaleEffect(ratio)
                .rotationEffect(Angle(degrees: degrees))
                .offset(x: 0, y: offsetY)
                .animation(.default)

            Button(action: {
                ratio = (ratio==1.0) ? 2 : 1
                degrees = (degrees==0) ? 360 : 0
                offsetY = (offsetY==0) ? -100 : 0
            }) {
                Text("Animation")
                    .foregroundColor(.white)
                    .padding()
                    .background(Capsule())
            }

        }
    }
}
```

## トランジションを適用する

トランジションは、ビューが追加or削除される時に実行されるアニメーションである。  
適用するには`transition`モディファイアを使うが、アニメーションの１つなので、  
`animation`モディファイアか、`withAnimation`関数と組み合わせて使う必要がある
```swift
// transition + animation
if 表示状態を入れたプロパティ {
    View()
        .transition(トランジションの種類)
        .animation(アニメーションの種類)
}

// transition + withAnimation
if 表示状態を入れたプロパティ {
    View()
        .transition(トランジションの種類)
}
Button(action: {
    withAnimation(アニメーションの種類) {
        表示状態を入れたプロパティの値変更
    }
}) {
    Text("...")
}
```

### 移動の方向を指定する

移動しながら消えていくトランジションを表示したい時、`slide`トランジションを使える。  
カスタムで実装することもでき、移動であれば、`move`メソッドを使えば手軽にできる
```swift
View()
    .transition(.move(edge: 移動する方向))
```

### トランジションを組み合わせる

トランジションを組み合わせるには、`AnyTransition`のエクステンションを作って、タイププロパティを追加する
```swift
extension AnyTransition {
    static var トランジション名 : AnyTransition {
        let insertion = 追加される時のトランジション
        let removal = 削除される時のトランジション
        return asymmetric(insertion: insertion, removal: removal)
    }
}
```
