# セレクタ

## 一覧

### 効率的なセレクタ

ブラウザはセレクタを右から左へ解釈ため、セレクタが多いほど遅くなる
あまり要素をセレクタに使わない方が良い
セレクタは、IDとclassをうまく使って、できるだけ短くわかりやすく記述する

* 悪い例
    `div ul li.strong { ... }`
* 良い例
    `.text.strong { ... }`

### セレクタの優先度

!important　＞　style属性
＞　IDセレクタ　＞　Classセレクタ＆属性セレクタ
＞　タイプセレクタ　＞　ユニバーサルセレクタ

* 全く同じ場合は、後のものに上書きされる
* 強く適用したい場合、その要素の後にstyle属性を書く
* 固有性を無視して適用するので、 !important をできるだけ使わない

### セレクタの種類

* **ユニバーサル Selector**
    `*` のこと。全ての要素に適用。
* **タイプ Selector**
    `h1{...}`
    該当する全ての要素に適用。
* **クラス Selector**
    `*.class{...}`
    複数指定できる。「*」は省略可能
* **ID Selector**
    `*#ID{...}`
    1回のみ、単独or子孫セレクタとして使う時が多い。「*」は省略可能
* **子孫 Selector**
    `div p{...}`
    半角スペースでつなぐ
* **子 Selector**
    `div > h1{...}`
    `>` でつなぐ、ある要素の直下にある要素のみ指定
* **隣接 Selector**
    `h1 + p{...}`
    `+` でつなぐ、弟を対象に
* **間接 Selector**
    `h1 ~ p{...}`
    隣接は兄直後の弟のみが対象だが、間接は全ての弟を対象
* **属性 Selector**
    1. `要素 [属性名]`
    2. `要素 [属性名 = "属性値"]`
    3. `要素 [属性名 ~= "属性値"]`
    4. `要素 [属性名 |= "属性値"]`
    5. `要素 [属性名 ^= "属性値のはじめ"]`
    6. `要素 [属性名 $= "属性値終わり"]`
    7. `要素 [属性名 *= "属性値一部"]`

### 擬似クラスセレクタ

擬似クラスはHTMLの文書構造だけではわからない、「ある状況」の要素を指定

* **ターゲット擬似クラス**
    `:target {  background: #ccc;  }`
    アンカーリングで飛んだ先の要素に適用　( アンカーリングは # で始まるページ内遷移のこと )
* **nth-child()擬似クラス**
    | 擬似クラス           | 使い方                                               |
    | -------------------- | ---------------------------------------------------- |
    | :nth-child(n)        | 最初からn番目の子要素(n= odd奇数  even偶数  3n  ...) |
    | :nth-last-child(n)   | 最後からn番目の子要素                                |
    | :nth-of-type(n)      | 特定の要素を最初からのn番目の対象                    |
    | p :nth-of-type(1)    | p要素から一つ目の要素=p要素                          |
    | :nth-last-of-type(n) | 特定の要素を最後からのn番目の対象                    |
    | :first-child         | 最初の子要素、:nth-child(1) と同じ                   |
    | :last-child          | 最後の子要素、:nth-last-child(1) と同じ              |
    | :first-of-type       | 最初の指定の種類の子要素                             |
    | :last-of-type        | 最後の指定の種類の子要素                             |
    | :only-child          | 唯一になる子要素                                     |
    | :only-of-type        | 唯一になる指定の種類の子要素                         |
* その他の擬似クラス
    | 擬似クラス | 使い方                                           |
    | ---------- | ------------------------------------------------ |
    | a:link     | 未訪問のa要素                                    |
    | a:visited  | 訪問済みのa要素                                  |
    | a:active   | クリック中のa要素                                |
    | a:hover    | カーソルが乗っているa要素                        |
    | :focus     | 選択された状態                                   |
    | :enabled   | 有効になっている要素(ボタンなど)                 |
    | :disabled  | 無効になっている要素(ボタンなど)                 |
    | :checked   | チェックされている要素(ラジオ、チェックボックス) |
    | :root      | ルート要素(html要素)                             |
    | :empty     | 内容が空の要素                                   |
    | :not(s)    | sというセレクタに当てはまらない要素              |

### 擬似要素セレクタ

ある要素の一部をスタイリングする時に使う

* 要素::first-line
    要素の最初の行を装飾
* 要素::first-letter
    要素の最初の文字を装飾
* 要素::before
    要素の直前に要素を挿入する
* 要素::after
    要素の直後に要素を挿入する

### before / after 擬似要素について

1. 対象要素の前後にデザインを追加できる

    ```css
    .class::before { content: “\201C”; }   /*前に引用符をつける*/
    .class::after { content: “\201D”; }    /*後に引用符をつける*/
    ```

2. contentプロパティに指定の文字だけでなく、カウンタや指定要素も挿入できる
3. counter-increment で連番を指定して、contentでカウンタ「nnnnnn」を表示させる

    ```css
    .class { counter-increment: nnnnnn; }
    .class::before { content: counter(nnnnnn); }
    ```

4. contentでattr()式を指定して、要素の属性値を表示 (文字列は「”」で囲む)
5. 後に「xxx pt」と表示する

    ```html
    <li class=“class”  data-point=“xxx”></li>
    ```

    ```css
    .class::after {
        content: “(“ attr(data-point) ”pt)”;
    }
    ```

6. counterには、任意の文字、URLを指定して画像、counter関数を指定してカウンタ、ttr関数を指定してその要素の属性値、を表示できる
7. 要素を重ね合わせたような表現もできる

    ```css
    .class {
        box-sizing: border-box;
        position: relative;
    }
    .class::before {
        background: #ff0;
        -webkit-transform: rotate(10deg);
    }   /*10度傾ける*/
    .class::after {
        background: #0f3;
        -webkit-transform: rotate(-10deg);
    }   /*-10度傾ける*/
    ```

8. 重なり順は、下から、元の要素、before要素、after要素となる
