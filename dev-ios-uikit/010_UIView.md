# UIView

## UIViewクラス

[UIView](https://developer.apple.com/documentation/uikit/UIView)

### 継承関係

```swift
class UIView : UIResponder      // : NSObject
```

### コンストラクタ

```swift
init()
init(frame: CGRect)
init?(coder: NSCoder)
```

## アプリ起動時に最初に表示する画面

View ControllerのAttributesインスペクタで、  
`Is Intitial View Controller`にチェックを入れる

## ビューの作成＆表示

ルートビューは`self.view`で参照し（`self`は省略可）、`addSubview()`で追加する
```swift
let label = UILabel()
label.text = "ラベルだ"
// ラベル領域
label.frame = CGRect(x:80, y:150, width:110, height:21)
// 背景色＆文字色
label.backgroundColor = UIColor.orange
label.textColor = UIColor.white
```
サブビューの作成は上と同じ。`UIView()`で作って、`frame`と`background`で座標＆サイズ＆背景色を指定する

## ビューの座標＆領域

ビューのプロパティ：
- `center`：　x座標とy座標を持つ`CGPoint型`
- `frame`：　ビューの領域を示す`CGRect型`
    1. `frame.origin`：変更可能。左上の座標を示す`CGPoint型`
    2. `frame.size`：　変更可能。幅と高さを示す`CGSize型`
    3. `frame.mixX`：　read-only。左上の座標x。yについても同様
    4. `frame.midX`：　read-only。中心の座標x。yについても同様
    5. `frame.maxX`：　read-only。右下の座標x。yについても同様
    6. `frame.width`：  read-only。幅
    7. `frame.height`：read-only。高さ
- `bounds`：　ビューが占める矩形の縦横サイズ。`CGRect型`だが位置は常に(0,0)とする

### ローカル座標の変換

上の`center`や`frame`はローカル座標なので、  
スーパービューが違う、２つのビューの間で座標を調整する時、`convert()`を使う
```swift
// view1の座標ptをview2の座標系に変える
view1.convert(pt, to:view2)
// view2の座標ptをview1の座標系に変える
view1.convert(pt, from:view2)
```
