# スマホ対応

## スマホへの対応

メディアクエリーを使っても、特別な指定をしない限り、スマホで見ると微妙に変わったりする

### webページが小さく表示される理由

昔では、スマホでページを閲覧する時、
最初にページ左上のごく一部しか見えなくて、
画面を縮小しなければならない。
それより、最初からページ全体を縮小して表示する方が良い
なので今では、ほとんどのスマホのブラウザは、
実際の表示領域の幅に合わせて表示するのではなく、
表示領域の幅が980pxあるものとして、webページを表示させる
結果として、webページはパソコンでみた時
と同様のものが縮小表示されることになる。

### 縮小しないで実サイズで表示させる

webページを最初から実寸で表示させられる。meta要素で設定する

```html
<meta name="viewport" content=" width=375, initial-scale=1.0 ">
<meta name="viewport" content=" width=device-width, initial-scale=1.0 ">
```

### 出力先に合わせて、異なるサイズの画像を表示させる

* **ピクセル密度で**

    100ppiの画面を基準にするとき、200ppi画面のサイズは縦横を２倍になって、「ピクセル密度が２倍」という
    img要素のsrcset属性で、
    ピクセル密度別に用意した大きさの異なる画像をカンマ区切りで列挙して指定できる
    srcset属性には、値として「画像のURL」だけでなく、
    半角スペースで区切って「その画像が何倍のピクセル密度向けであるのか」も示せる

    ```html
    <img src="１倍画像URL" srcset="画像URL ピクセル密度, 画像URL ピクセル密度, ........ " alt="">
    <img src="log200.png" srcset="log400.png 2x, log600.png 3x" width="200" height="100" alt="サンプルロゴ">
    ```

* **画像のサイズで**

    画像の実サイズの横幅を指定することで、
    複数用意された画像のうち、どれを使うのが最適なのかを判断できる
    この方法は主に、画像の表示サイズが可変な場合に使われる

    ```html
    <img srcset="画像URL 画像の幅w, 画像URL 画像の幅w, ...... " sizes="単位付き表示幅" src="画像URL" alt="">
    <img srcset="small.png 500x, medium.png 800w, large.png 1200w" sizes="100vw" src="small.png" alt="写真">
    ```

    注意： size属性は画像の表示幅を指定する属性 単位「vw」は、
    表示領域全体の幅の何パーセントかを示す
    size属性とwidth属性の違い(size属性のできること)：

    1. CSSの単位をつけて指定する
    2. CSSの calc() などの関数が使える

    ```css
    sizes="(max-width: 600px) 100vw, (max-width: 800px) 660px, calc(900px - 2vw)"
    ```

    `calc(900px - 2vw)` は、900pxから表示領域の幅２倍分を引いた値が幅となる

    メディアクエリーの () で示す条件と組み合わせて複数の幅が指定できる

    ```css
    sizes="(max-width: 600px) 100vw, (max-width: 800px) 660px"
    ```

* **条件に合致した時に使う画像を詳細に指定する**

    sizes属性は、条件に合う表示幅を指定できるが、
    使用する画像ファイルまで指定できない。
    メディアクエリーの条件ごとに使う画像やそれがどのピクセル向けか、
    といった情報を細かく指定したい場合、
    picture要素にsource要素とimg要素を入れて使う
    picture要素に、source要素がいくつあってもよくて、
    media属性で条件指定する。
    最初に合致したsource要素だけが有効になる。
    source要素には、srcsetとsizesが指定できるが、
    「x」と「vw」をつけてない画像は、「1x」とみなす
    picture要素に、必ずimg要素を１つだけ配置する

    ```html
    <picture>
        <source media="min-width: 1200px" srcset="pic1200.jpg, pic2400.jpg 2x">
        <source media="min-width: 800px" srcset="pic800.jpg, pic1600.jpg 2x">
        <img src="pic500.jpg" srcset="pic1000.jpg 2x" alt="写真">
    </picture>
    ```

    ```css
    img{ display: block; }
    @media screen and (min-width: 1200px){
        img{width: 1200px;}
    }
    @media screen and (min-width: 800px) and (max-width: 1199px) {
        img{width: 800px;}
    }
    @media screen and (max-width: 799px){
        img{width: 100vw;}
    }
    ```
