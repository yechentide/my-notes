# メディアクエリー

## メディアクエリーとは

メディアクエリーを使うと、出力する媒体や状態ごとに、適用するCSSを変えられる
CSSの組み込み方の中で、link要素やstyle要素にはmedia属性を指定できると説明した。
CSS3からこの機能は拡張されて、出力媒体の種類だけでなく、
その媒体の特性や状態を示す式も書き込めるようになっている。
これによって、windowの幅によって適用したいCSSを指定できるようになり、
この機能をメディアクエリーと呼んでいる

### メディアクエリーの書き方

* CSS3以前の方法

    ```css
    media="screen"
    media="screen, print"
    ```

* CSS3からの方法
    「and (メディア特性:値)」を必要な数だけ追加できる

    ```css
    media="screen and (min-width: 600px)"
    media="screen and (min-width: 600px) and (max-width: 800px)"
    media="screen and (min-width: 600px) and (max-width: 800px), print"
    ```

### メディア特性

width, min-width, max-width
height, min-height, max-height
device-width, min-device-width, maxdevice-width
orientation(値はportrait, landscape)
aspect-ratio,device-aspect-ratio
など

### @media について

@media を使った書式を使うと、CSSのソースコードの中に、
出力媒体とメディアクエリーの指定を書き込める。
こうすると、特定の出力媒体が特定の状態時にのみ適用することになる

```css
@media screen and (min-width: 600px) and (max-width: 800px) {
    main{ background: white; }
}
```
