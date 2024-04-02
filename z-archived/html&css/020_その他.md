# その他

## ノーマライズ

normalize.cssは、ブラウザごとのCSS解釈の差をなくすためのもの。
HTML5の初期化にも対応
[normalize.css](http://necolas.github.io/normalize.css/)
reset.cssは同じ働きだが、ブラウザデフォルトの指定を全部消してからスタートするので、
基本的にnormalize.cssの方が好ましいだが、場合によってreset.cssを使う
[cssreset](http://www.cssreset.com/)

### 古いIEに認識させる

IEの古いバージョンは、HTML5の新要素を認識できないため、
Javascriptで要素名を指定する必要がある。
[html5shiv](http://code.google.com/p/html5shiv/)

```html
<head>
    <!—[if IT IE 9]>
    <script src=“./js/html5shiv.js”>
    <![endif]—>
</head>
```

[selectivizr](http://selectivizr.com/)
IEでCSS3のプロパティを使えるようにする [css3pie](http://css3pie.com/)
IEでCSS3のセレクタを使えるようにする
jQueryなどのライブラリを共存できる

### CSSプリプロセッサライブラリ

CompassやnibなどのCSSプリプロセッサライブラリを使うと、
ベンダープリフィックスを自動でつけてくれる
[compass-style](http://compass-style.org/)
[visionmedia](http://visionmedia.github.io/nib/)

```css
@import “./compass/css3”;
.botton { @include border-radius(2px); }
↓自動変換
.bottom{-webkit-border-radius: 2px;
        -moz-border-radius: 2px;
        -ms-border-radius: 2px;
        -o-border-radius: 2px;
        border-radius: 2px;
}
```

### ベンダープリフィックス

```css
{
    border: 6px solid #ccc;
    -o-border-radius: 8px;          /*Opera用の記述*/
    -moz-border-radius: 8px;        /*Mozilla(firefox)用の記述*/
    -ms-border-radius: 8px;         /*Microsoft(IE)用の記述*/
    -webkit-border-radius: 8px;     /*webkit(Chrome , Safari)用の記述*/
    border-radius: 8px;             /*W3C仕様に沿った記述も必須*/
}
```

### ブラウザ対応(プログレッシブ・エンハンスメント)

グラデーションが使えない時のために背景色を指定

```css
background: #06f;
背景にグラデーションを指定
background: -webkit-gradient( linear , left top , left bottom ,
                            color-stop(0,#06f) , color-stop(1,#03f) );
background: -webkit-linear-gradient( top , #06f 0 , #03f 100% );
background: -moz-linear-gradient( top , #06f 0 , #03f 100% );
background: -o-linear-gradient( top , #06f 0 , #03f 100% );
background: -ms-linear-gradient( top , #06f 0 , #03f 100% );
background: linear-gradient( top , #06f 0 , #03f 100% );
```
