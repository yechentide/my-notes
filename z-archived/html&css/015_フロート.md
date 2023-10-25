# フロートレイアウト

## 基本説明

floatプロパティで指定する。値は：left, right, none
フロートの状態になると、インライン要素でも、
そのボックスはブロックレベルのボックスになる
解除にはclearプロパティを使う。値は：left, right, both, none
clearプロパティは、ブロックレベル要素にしか指定できない

## フロートによる２段組みレイアウト

```html
<header></header>
<main></main>
<div id="sub"></div>
<footer></footer>
```

```css
main{float:right;} #sub{float:left;} footer{clear:both;}
```

## フロートによる３段組みレイアウト

1. 方法１：
    （問題点はmainを中央に配置できない）

    ```html
    <header>
    </header>
        <main>
        </main>
        <div id="sub1"></div>
        <div id="sub2"></div>
    <footer>
    </footer>
    ```

    ```css
    main{
        float:right;
    }
    #sub1,#sub2{
        float:left;
    }
    footer{
        clear:both;
    }
    ```

2. 方法２：
    （問題点は段の高さが揃っていない。背景をいじれば、揃っているように見せられる。）

    ```html
    <header>
    </header>
        <div id="contents">
            <main></main>
            <div id="sub1"></div>
        </div>
        <div id="sub2">
        </div>
    <footer>
    </footer>
    ```

    ```css
    main{
        float:right;
    }
    #sub1{
        float:left;
    }
    contents{
        float:left;
    }
    sub2{
        float:right;
    }
    footer{
        clear:both;
    }
    ```

## clearfixで不具合解消

floatプロパティとclearプロパティは、
元々画像などの横に文字を回り込ませる目的で用意されたもので、
それを段組のようにボックスを横に並べるように使うと不具合が生じる。
(フロートを指定した要素のボックスは、親要素ボックスの高さと無関係になって、親要素からはみ出す)
単純にコンテンツを２段組にするだけであれば、clearを使っても特に問題ない。
しかし、段組部分を囲う親ボックスにボーダー表示させたり、
背景を表示させたりすると、はみ出さないようにしたい時、ちょっと手間がかかる。

### やり方は４通り

1. (勧めない)親要素の内部に置いて、floatプロパティが指定されている要素よりも後ろに、
   ブロックレベル要素を追加して、その要素にclearを指定
2. (勧めない)親要素自体もフロートさせる
3. overflowプロパティで「visible」以外の値を指定する 例えば、overflow: hidden　で指定すると、はみ出した部分は表示されなくなる。場合によって使えない
4. 副作用がなく、いつでも安全に利用できる ーー＞ clearfix

### clearfixの原型

contentプロパティで、floatを含む親要素の最後に空の文字列を挿入し、
displayでブロック化して、それにclearを指定する。
CSSへの対応が不十分の時代だったので、
裏技的手法(CSSハック)も組み合わせて使うことで、
どのブラウザでもうまく機能するようになった。

```css
.clearfix:after{
    content: ".";
    display: block;
    height: 0;
    clear: both;
    visibility: hidden;
}
.clearfix{
    display: inline-block;
}
/*Hides from IE-mac*/
* html .clearfix{ height: 1%; }
.clearfix{ display: block; }
/*End hide from IE-mac*/
```

使用方法は、親要素に「class="clearfix"」を指定するか、
.clearfix の部分をすでに指定されているクラス名にする

### 現在のclearfix

```css
.clearfix:after{
    content: "";
    display: block;
    clear: both;
}
```
