# リスト＆テーブル

## リスト

### 順序つきリスト

```html
<ol>
    <li>...</li>
    <li>...</li>
</ol>
```

### 順序なしリスト

```html
<ul>
    <li>...</li>
    <li>...</li>
</ul>
```

### 定義リスト

```html
<dl>
    <dt>Aさん</dt>
        <dd>ーー＞優しい人</dd>
    <dt>Bさん</dt>
        <dd>ーー＞イケメン</dd>
    <dt>cさん</dt>
        <dd>ーー＞頭が良い</dd>
</dl>
```

## テーブル

### テーブル構造

HTML5では、「border="1"」はレイアウトのために使っていないことを明示している。
1を入れるか、空にする
thは見出しで、tdはセルとなる

```html
<table border="1">
    <caption>表のタイトル</caption>
    <tr>
        <th>1行目1列目</th>
        <th>1行目2列目</th>
    </tr>
    <tr>
        <td>2行目1列目</td>
        <td>2行目2列目</td>
    </tr>
    <tr>
        <td>3行目1列目</td>
        <td>3行目2列目</td>
    </tr>
</table>
```

### th要素に指定できる属性

* `colspan`：右方向にセルを結合、値は整数
* `rowspan`：下方向にセルを結合、値は整数
* `headers`：このthセルの見出しとなるセルのidを入れる。音声ブラウザなどのため
* `scope`：この見出しセルの対象となるセルの範囲を示す。
    キーワードは、row, col, rowgroup, colgroup

### td要素に指定できる属性

* `colspan`
* `rowspan`
* `scope`

### セルの結合

h要素やtd要素に対して、colspan属性やrowspan属性を使って結合する

* colspan属性は横に隣り合うセルを結合する
    `<td colspan="3"></td>`

* rowspan属性は縦に隣り合うセルを結合する
    `<td rowspan="2"></td>`

### 表が長くなる場合

thead tbody tfoot で表の行を区切って、区別しやすくできる

```html
<table>
    <thead>
        <tr>
            <th></th><th></th><th></th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td></td><td></td><td></td>
        </tr>
    </tbody>
    <tfoot>
        <tr>
            <td></td><td></td><td></td>
        </tr>
    </tfoot>
</table>
```
