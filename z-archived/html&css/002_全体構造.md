# 全体構造

## 全体像

```html
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>タイトル</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="サイトの説明文">
    <meta name="keywords" content="キーワード,キーワード,キーワード,キーワード,キーワード,キーワード">
    <link rel="stylesheet" href="">
    <script src=""></script>
</head>
<body>
    内容
</body>
</html>
```

### `<!DOCTYPE html>`

これはタグではなく、DOCTYPE宣言(文書型宣言)。  
本来ならHTML5にはDOCTYPE宣言が不要だが、それをつけなければ、古いバージョンのページと認識されるため、必要最低限の簡単なDOCTYPE宣言をつけることとなる。  
DOCTYPE宣言があれば、ブラウザの表示モードが自動的に切り替える。

### html

DOCTYPE宣言の後にはhtml要素を配置する。  
それら以外の要素は全てhtml要素の中に書く。  
htmlに属性langを指定して、このページの言語を示すことができる。

### head

ページに関する情報を入れるための要素である。  
文字コード、ページタイトル、適用する外部ファイルのパス、などを示す要素を順不同で必要なだけ入れることができる。  
head要素内の内容は、タイトル以外、基本的にページに表示されない。

* `title`：ページのタイトル
* `meta`：ページに関する情報(メタデータ)を示す。  
    終了タグがない。  
    文字コードは、`charset`属性で指定する。  
    それ以外の情報を指定する場合、情報の種類によって、  
    `name`属性または`http-equiv`属性で種類を示し、  
    具体の情報を`content`属性の値として指定する。
* cssファイルの読み込み：  
    `link`要素で指定する。`href`属性にファイルのパスを入れる。
* javascriptファイルの読み込み：  
    `script`要素で指定する。`src`属性にファイルのパスを入れる。

### titleタグとmetaタグの順序

titleタグとmetaタグの順序は重要である。

1. 文字コードの指定（charset）
   * 文字コードの指定は、ドキュメントの最初の1,024bytes以内に含まれている必要がある。  
   * IEで文字化けを防ぐため、titleタグのような文字データの前に記述する。
2. 互換性機能（x-ua-compatible）
   * 「IE=edge」を指定して、IEの最新バージョンの標準モードで表示させる。
   * この指定は、titleタグと他のmetaタグ以外より前に記述する。

### link要素に指定できるもの

```css
<link rel="ファイルの種類" href="ファイルのURL" media="適用対象" type="MIMEタイプ">
```

* ファイルの種類
    |     値     |                 意味                 |
    | :--------: | :----------------------------------: |
    | stylesheet |         スタイルシート(CSS)          |
    | alternate  | 代替バージョン(異なる言語＆媒体向け) |
    |    icon    |           ページのアイコン           |
    |    prev    |   連続しているページ中の前のページ   |
    |    next    |  連続しているページ中の後ろのページ  |
* 適用対象(省略OK)
    |     値     |                  意味                  |
    | :--------: | :------------------------------------: |
    |    all     |               全ての機器               |
    |   screen   |   パソコン＆スマホ＆タブレットの画面   |
    |   print    |                プリンタ                |
    | projection |              プロジェクタ              |
    |     tv     |                 テレビ                 |
    |  handheld  |       携帯用機器(スマホではない)       |
    |    tty     |   文字幅が固定の端末(ターミナルなど)   |
    |   speech   | スピーチ・シンセサイザー(音声読み上げ) |
    |  braille   |            点字ディスプレイ            |
    |  embossed  |              点字プリンタ              |
* MIMEタイプ(省略OK)
    デフォルトの値は`text/css`

### body

bodyタグには、ブラウザに表示させたい内容を入れる。
