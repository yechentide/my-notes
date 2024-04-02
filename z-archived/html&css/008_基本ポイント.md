# 基本ポイント

## 一覧

### 専門用語

```css
h1{ color: white; }
```

* セレクタ：h1
* プロパティ：color
* プロパティ：white

```css
h1{
    color: rgb(233, 233, 233);
    font-size: 24px;
}
```

* 宣言ブロック： { と } で囲んだ範囲
* 宣言：宣言ブロックの中の;で区切った部分

### 記述場所

1. HTMLの`link`要素で読み込む

    ```html
    <head>
    <link rel="stylesheet" href="">
    </head>
    ```

2. `style`要素内に記述する
    しかし、この方法では他のHTMLやCSSの表示指定を共有できなくなる。  
    そのため、何らかの理由がある場合にのみstyle要素を使う。
    style要素に`media`属性と`type`属性を指定できる

    ```html
    <head>
    <style>
        h1{ color: white; }
    </style>
    </head>
    ```

3. style属性で指定する
    メンテナンス性が低下するため、おすすめできない

    ```html
    <h1 style="color:white;">
    ```

### 書き方のルール

* `セレクタ`、`プロパティ`、`プロパティ値`、`{` 、`}` 、`:` 、`;` の各記号の前後に、半角スペース・改行・タブを自由にいれられるが、見やすくなるように、いつも一定のパターンで統一するのが良い
* セレクタは、`,`で区切って同時に複数指定できる

### 記述順番

機能で並べる

1. 表示系(display, list-style, overflow ..)
2. 配置系(position, float, clear, z-index ..)
3. ボックス系(width, height, margin, padding ..)
4. 背景系(background)
5. テキスト系(color, font, text-align ..)

### コメント

```css
/*コメント*/
/*
    コメント
*/
```

### CSSファイルの文字コード

必ずソースコードの先頭に書く！

```css
@charset "UTF-8";
```

### CSSファイルを読み込む

CSSファイルの中から別のCSSファイルを読み込むことができる
処理速度が落ちるため、しないのが良い

```css
@import "./style.css";
@import url(../aa/main.css);
```

また、htmlの`media`属性の値と同じものを指定できる

```css
@import url(../aa/main.css) print;
@import url(../aa/main.css) tv, projection;
```
