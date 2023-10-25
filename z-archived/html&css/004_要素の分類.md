# 要素の分類

## 歴史

### HTML5以前のバージョンで使われた分類方法

HTML5では使われなくなったが、CSSにおいて重要な概念となる

* ブロックレベル要素：ひと固まりの部分
* インライン要素：画像や文などブロックレベル要素の一部v
* その他の要素

### HTML5の分類方法

HTML5のコンテンツモデルでは、七つのカテゴリーがある。

* フローコンテンツ   (Flow content)
* セクショニングコンテンツ   (Sectioning content)
* 見出しコンテンツ   (Heading content)
* フレージングコンテンツ   (Phrasing content)
* 組み込みコンテンツ   (Embedded content)
* インタラクティブコンテンツ   (Interactive content)
* メタデータコンテンツ   (Metadata content)

### 参考

[HTML5における要素の分類](http://www.htmq.com/html5/007.shtml)

[HTML5のコンテンツ・モデルの概念を理解する](https://qiita.com/yu310fu/items/06dd90901490ba1b619d)

## 分類方法

### フローコンテンツ

ある特定の要素の内部にしか配置できないといった制限がなく、body要素の中であればどこでも良い
「body要素の子要素として直接配置できるのはフローコンテンツのみ」という仕様になっている。
ほとんどの要素はまずこのカテゴリーに含まれている。
そして、それらのほとんどはFlow contentでありつつ、他のカテゴリーにも含まれる
大半の要素はフローコンテンツに属するので、属さない要素を知ればいい。

#### Flow contentに属しない要素

|        関連性        |                          要素                           |
| :------------------: | :-----------------------------------------------------: |
|      ルート要素      |                          html                           |
|    メタデータ関連    |              head, meta, title, link, base              |
|    セクション関連    |                          body                           |
|  テキストのグループ  |                 li, dd, dt, figcaption                  |
|    テキストの一部    |                         rt, rp                          |
|     組み込み関連     |                  param, source, track                   |
|        表関連        | caption, th, tr, td, tbody, thead, tfoot, col, colgroup |
|     フォーム関連     |                legend, optgroup, option                 |
| インタラクティブ関連 |                         summary                         |

#### Flow contentの主な要素

* `header` `main` `footer`
* `p` `hr` `address` `div`
* `ul` `ol` `dl`
* `blockquote`   <--- ブロックレベルの引用文。インラインで引用したい場合は「q」要素を使う

### セクショニングコンテンツ

セクショニングコンテンツに属する要素は４つだけ

* `section`：章・節・項のよな構成になる時
* `article`：そのセクションだけで独立している内容となる時
* `aside`：ページのメインコンテンツと関係ない時
* `nav`：ナビゲーション

### 見出しコンテンツ

見出しコンテンツに属する要素は６つだけ
`h1` `h2` `h3` `h4` `h5` `h6`

### フレージングコンテンツ

まとまった文や段落ではなく、文の要素(一部)のこと。
フレージングコンテンツの内容として入れられる要素は、フレージングコンテンツのみ。

#### Phrasing contentの主な要素

* `em`：特定な部分を強調する
* `strong`：特定な部分が重要だと示す
* `i`：生物の学名、船の名前、専門用語、慣用句、翻訳された文書などを斜体で表示する
* `b`：注目してほしい所を太字にする。製品名、キーワードなど
* `q`：引用文。あえて使わなくて、自分で引用符をつけても良い。
* `small`：小さな文字で表示されるような部分を示す
* `br`：改行
* `code`：ソースコードを示す。preタグと一緒に使うことが多い`<pre><code></code></pre>`
* `span`：特に意味がなく、範囲を示す
* `ruby`：ふりがな、など
* `a`：リンクを作成

### 組み込みコンテンツ

画像、動画、音声、他のHTMLページを組み込むために使用する
組み込みコンテンツに属する要素は７つだけ

* `<img>`：画像を挿入。

    ```html
    <img src="a.jpg" alt="代替テキスト" width="100px" height="100px">
    ```

* `<video>`：動画を挿入。

    ```html
    <video src="a.mp4" width="100px" height="100px" poster="b.jpg" controls autoplay loop muted></video>
    ```

​   `src`は１つのデータしか指定できない
​   `poster`属性で、再生できるまでの間に、表示させておく画像を指定。
​   `controls`属性で、再生ボタンや音量調整スライダーを表示させる。
​   この属性がなければ、JSなどで再生することになる。
​   `autoplay`属性で自動再生。
​   `loop`属性でループ再生。
​   `muted`で音量０で再生。

* `<audio>`：音声データを挿入。
    属性は`video`より少ないが、説明は同じ

    ```html
    <audio src="a.mp3" controls autoplay loop muted></audio>
    ```

* `<source>`：動画や音声のデータを複数指定できる。その時videoとaudioでsrcを使わない

    ```html
    <video controls autoplay loop muted>
        <source src=""><source src="">
        <p>例</p>
    </video>
    ```

* `<iframe>`：Webページ内で別のWebページを表示するときに使う。

    ```html
    <iframe src="https://google.com" name="sample" width="200" height="150">
    この部分はインラインフレームを使用しています。
    </iframe>
    ```

* その他
    `canvas` `object` `embed`

### インタラクティブコンテンツ

ユーザーによる操作や入力が可能な要素。
| グループ             | 要素                                           |
| -------------------- | ---------------------------------------------- |
| テキスト関連         | a                                              |
| 組み込み関連         | img, iframe, video, audio, embed, object       |
| フォーム関連         | label, input, textarea, select, button, keygen |
| インタラクティブ関連 | details, menu                                  |

* `details`：
    一般的なアプリケーションでよく見られる、
    ▶︎マークをクリックすると情報の表示/非表示を切り替える部分をマークアップするための要素。
    ▶︎マークの横に常に表示する内容は、summary要素として、detailsの最初の子要素にする。
* `menu`：
    基本的にはコマンドのリストをマークアップするための要素。
    type属性を指定しない時はリストの一種となが、
    type属性の値がcontextの時はコンテキストメニューとなり、
    toolbarのときはツールバーとなる。
    ツールバーになるときにのみ、インタラクティブコンテンツになる。
* `input`：
    テキストの入力欄や送信ボタン、ラジオボタン、チェックボックスなどの、
    データの入力・選択・操作するための空要素。
    typeにはtext, submitなど20種類くらい指定できる

### メタデータコンテンツ

Webページには表示されないが、ページ全体に関する情報を設定するための要素。
基本的にhead要素に配置する。

* `meta`
    他のメタデータコンテンツでは示せないよなメタデータを示す
    文字コードを示す場合は専用のcharset属性を使うが、
    それ以外は、name & contentのペア、または、http-equiv & contentのペアを使う
* `title`
    Webページのタイトルまたは文書名
* `link`
* `style`
* `script`
* `noscript`
* `command`
* `base`
    Webページで使われる相対URLの基準位置とする絶対URLを指定する
