# 図形の描画

## シェイプの描画

### シェイプの共通事項

- 大きさは`frame`モディファイアで指定できる  
    指定しない場合、親ビューにフィット
- 色は`foregroundColor`モディファイアで指定する

### 様々なシェイプ

- 長方形: `Rectangle()`
- 角丸長方形: `RoundedRectangle(cornerRadius: 50)`  
    (角の丸め具合は半径で指定)
- 角丸長方形: `Capsule()`  
    (角の丸め半径=幅と高さの短い方の半分)
- 円: `Circle()`
- 楕円: `Ellipse()`

## Path

`Path`もシェイプの一種で、形状はコードで定義

### Pathの基本形

```swift
Path { path in
    // ここでパスの内容を定義
}
// 線や色の指定
```

### Pathの描画

- 点の移動  
    `path.move(to: CGPoint)`  
    `Path`は点と点ウィ結ぶ、直線や曲線で構成され、  
    これらの線を追加するまえに、線の始点まで移動する必要がある。
- 直線を追加  
    `path.addLine(to: 終点)`: 現在の点から線を引く  
    `path.addLines(points)`: 渡された配列内の点を結ぶ
- 円弧  
    `path.addArc(center:radius:startAngle:endAngle:clockwise:)`  
    clockwiseはtrueなら左回り、falseなら右回り  
    角度は数学とは逆(下の方が90°)
- ベジェ曲線  
    `path.addCurve(to:control1:control2:)`  
    `path.addQuadCurve(to:control)`  
    `move()`で始点を指定する必要がある
- 楕円  
    `addEllipse(in:)`
- 四角形  
    `addRect()`  
    `addRects()`
- 角丸四角形  
    `path.addRoundedRect(in:cornerSize:)`

### Pathの色など

- `.stroke(lineWidth: 線の太さ)`
- `.fill(色)`

## ビューとの組み合わせ

### backgroundモディファイア + シェイプ

```swift
HStack {
    Button(action: {}) {
        Text("Button01")
    }
    .padding()
    .background(
        Capsule()
            .foregroundColor(.blue)
    )
}
```

### クリッピングに使う

シェイプのクリッピングを行うと、ビューがシェイプの形で切り抜かれる  
クリッピングを行うには、ビューに`clipShape`モディファイアを使う
```swift
Image("Sample")
    .resizable()
    .aspectRatio(contentMode: .fill)
    .frame(width: 300, height: 300)
    .clipShape(Circle())
```

### マスクを適用する

`mask`モディファイアを使うと、ビューにマスクを適用させ、  
マスクのアルファチャンネルによって、透明になる場所とそうでない場所をコントロールできる
```swift
View()
    .mask(MaskView())
```

## グラデーション

SwiftUIでは、グラデーションを描画するビューを使う

### Gradient

```swift
Gradient(colors: [.blue, .white])
Gradient(stops: [
    .init(color: .red, location: 0.0),
    .init(color: .white, location: 0.3),
    .init(color: .green, location: 1.0)
])
```

### LinearGradient

`UnitPoint`の引数は、LinearGradientビューの大きさを1として、0.0~1.0までの値を指定。
```swift
LinearGradient(
    gradient: Gradientビュー,
    startPoint: UnitPoint(x:y:),
    endPoint: UnitPoint(x:y:))
```

### RadialGradient

```swift
RadialGradient(
    gradient: Gradientビュー,
    center: UnitPoint(x:y:),
    startRadius: 開始半径,
    endRadius: 終了半径)
```

### AngularGradient

```swift
AngularGradient(
    gradient: Gradientビュー,
    center: UnitPoint(x:y:),
    startAngle: 開始角度,
    endAngle: 終了角度)
```

## 画像処理を行う(エフェクト)モディファイア

- blur: ガウスぼかしを行う
- shadow: ビューに影をつける
- opacity: 不透明度を変更する
- brightness: 明るさを変更する
- contrast: コントラストを変更する
- colorInvert: 色を反転する
- colorMultiply: 色の値に指定した色を掛ける
- blendMode: ビューが重なった時の色の描画モードを設定する
- saturation: 彩度を変更する
- grayscale: グレースケール化する
- hueRotation: 色相を変更する
- luminanceToAlpha: 輝度からマスク画像を作る
