# Markdown

## 自分のルール

- リスト＆引用の下が普通の文の時、間に空行を入れる
- 表の上下に空行を入れる

## Markdownエディタ

自分がよく使うものは３つある。
- Vscode: プラグイン導入すれば、一番使いやすい。  
    おすすめプラグイン ==> `docs-markdown` `Markdown All in One` `markdownlint` `Markdown Preview Enhanced`
- Obsidian: 無料でUIが好き！使いやすい！
- Typora: Typoraは正式版から有料になったので使わなくなった。

|         |PC       |Phone    |
|:-------:|:-------:|:-------:|
|Edit     |VSCode   |Obsidian |
|View     |VSCode   |Obsidian |

## Markdown文法

### 基本的な文法

- 見出し: `#`から`######`まで。HTMLの`<h1>`から`<h6>`までと対応。
- リスト:
    - 順序なしリスト: `-` `+` `*`を使えるけど、`markdownlint`を使う場合、`-`以外は文句言われる。
    - 順序ありリスト: `1.` `2.`のように使う。
    - リストはネストできる。
    - `markdownlint`では、リストのインデントは指定されたが、無視するようにした。
    - インデントを固定の4つのspaceにしたいので、`Markdown All in One`の設定も変えた。
- チェックボックス: 選択されてない時は`- [ ]`、選択された時は`- [x]`。
- 段落: 改行を２回する。
- 改行: 行末に半角スペースを２つ入れる。
- 水平線: `*`または`-`を3つ以上書けば水平線になる。間に空白があっても大丈夫。`---`か`- - -`がおすすめ。
- フォント:
    - *斜体*: `*`で囲む。
    - **太字**: `**`で囲む。
    - ***斜体＆太字***: `***`で囲む。
    - ~~打ち消し線~~: `~~`で囲む。
    - 使ってるフォントにより効果がない場合もある。
- リンク: そのまま貼り付けてもよし。名前付きリンクは、`[名前](URL)`の形で記述する。
- 画像: `![画像タイトル](画像URL)`の形で記述する。
- ソースコード:
    - 一行だけ: `` ` ``で囲む
        - `` ` ``を表示させたい場合は、``` `` ` `` ```のように書く。
    - 複数行: ```` ``` ````で行を囲む。最初の```` ``` ````の後ろに言語を指定できる。

### ちょっとリッチな表現

- 表: 1段目がヘッダで、2段目は中央揃えなどを指定。ヘッダとそろえ方の指定は**必須**！
    ```markdown
    |Header    |Header    |  Header  |    Header|
    |----------|:---------|:--------:|---------:|
    |left      |left      |  center  |     right|
    ```
- 引用: 引用したい文の文頭に`>`をいれる。二重引用には`>>`を文頭におく。
- 注釈: 記号を置く場所は、`[^記号]`を入れる。注釈のところでは、`[^記号]:`の後ろに注釈内容を書く。
- 下付き文字: `~`で囲む。 f~x~(h)
- 上付き文字: `^`で囲む。 x^2^
- ハイライト: `==`で囲む。 ==ハイライト==  
    ただし、Githubのページでは効かない。VSCodeもプラグイン入れないと効かない。

### LaTeXで数式を表現

`$`でLaTexを囲む。複数行の場合は`$$`を使う
```markdown
$ \textcolor{red}{文字色を赤色にする方法} $
$$
a = \frac{1}{b}
\tag{1}
$$
```
$ \textcolor{red}{文字色を赤色にする方法} $
$$
a = \frac{1}{b}
\tag{1}
$$

### その他

- HTMLタグはそのまま使えるが、推奨されない
- 顔文字コードは`:`で囲めば使える
