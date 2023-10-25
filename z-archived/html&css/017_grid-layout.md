# Grid Layout

## グリッドレイアウト

グリッドレイアウトは、ボックスを格子状に縦横に区切って、
そこに子要素を自由に当てはめていくという手法である。
フレキシブルボックスレイアウトより複雑だが、縦か横かの一方向に限定されない

```html
<div id="page">
    <header></header>
    <main></main>
    <div id="sub1"></div>
    <div id="sub2"></div>
    <footer></footer>
</div>
```

```css
#page{
    /* 容器となる親要素にgridを指定する */
    display: grid;
    /* このプロパティは列の幅を左から順に指定する */
    grid-template-columns: 100px 200px 100px;
    /* このプロパティは行の高さを上から順に指定する */
    grid-template-rows: auto auto auto;
}
```

子要素をグリッドセルに当てはめていく。指定しない場合、順に当てはめていく
(これ以外にも他のやり方がある、下を参照)

```css
header{
    /* 縦の１番目の線はら４番目の線までの領域 */
    grid-column: 1 / 4;
    /* 横の１番目の線から１セル分の領域 */
    grid-row: 1;
}
main{
    grid-column: 2;
    grid-row: 2;
}
#sub1{
    grid-column: 1;
    grid-row: 2;
}
#sub2{
    grid-column: 3;
    grid-row: 2;
}
footer{
    grid-column: 1 / 4;
    grid-row: 3;
}
```

### グリッドをわかりやすく定義する別の方法

```css
#page{
    display: grid;
    grid-template-columns: 100px 200px 100px;
    grid-template-rows: auto auto auto;      /* ここまでは同じ */
    grid-template-areas: "head head head"    /* グリッドのセルごとに名前をつける */
                         "sub1 main sub2"
                         "foot foot foot";
}
header{
    grid-area: head;
}
main{
    grid-area: main;
}
#sub1{
    grid-area: sub1;
}
#sub2{
    grid-area: sub2;
}
footer{
    grid-area: foot;
}
```

### その他

```css
/* 幅を1:2:3の比率に分割。fr=fraction(比) */
grid-template-columns: 1fr 2fr 3fr;
/* 幅から(200+200)px引いた残りの幅を、真ん中のセルの幅にする */
grid-template-columns: 200px 1fr 200px;
```

grid-column = grid-column-start + grid-column-end
grid-row = grid-row-start + grid-row-end
grid-area = grid-column + grid-row
