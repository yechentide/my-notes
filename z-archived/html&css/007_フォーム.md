# フォーム

## form要素

### コード例

```html
<form action="送信先のURL" method="送信方法" enctype="MIMEタイプ" name="フォームの名前">
    <input type="text" name="...">
    ...
</form>
```

### 使い方

Webサーバにデータを送るための入力域を表示
通常、input、select、textareaなどの入力要素を、form要素で囲むことで実現する
内容となる入力域には、input要素や、ドロップダウンリスト、テキストエリアなどが使える

* `action`：入力されたデータを受け取るプログラムの送り先URLの指定
    サーバ側のプログラムを指定するのが普通だが、メールで送信するのもできる
* `method`：「送信方法」の指定。get(初期値)かpostが指定できる。
    getはURLの後ろにデータを付け加えて送信。postはURLとは別にデータを送信
* `enctype`：postの時に使うコード化の方法。
    (初期値)`application` `x-www-form-ur` `multipart` `form-data` `text/plain`
* `name`：フォームを参照するための名前の指定

## inout要素

### 一般形

```html
<input type="text" name="部品の名前" value="初期値/ラベル/送信値"
       size="文字数" maxlength="最大文字数" checked readonly disabled
       src="画像URL" width="" height="" alt="代替テキスト">
```

* `type`：
    1. text：１行のテキスト入力欄(一般テキスト用)
    2. password：１行のテキスト入力欄(パスワード用)
    3. checkbox：チェックボックス
    4. radio：ラジオボタン
    5. file：ファイル送信用部品
    6. hidden：画面上には表示させずに送信するテキスト
    7. submit：送信ボタン
    8. reset：リセットボタン
    9. button：汎用ボタン
    10. image：画像の送信ボタン
    11. 実装されてない部品：search, tel, url, email, color, range, number, time, .......
* `name`：この部品の名前を指定する。データはこの名前とペアで送信される。
    同じ選択項目内のチェックボックス＆ラジオボタンには、同じ名前をつける必要がある
* `value`：テキスト入力欄の場合は初期文字に、
ボタンの場合はそのボタンのラベルに、
チェックボックス＆ラジオボタンの場合は、選択される時にサーバーに送る値
* `size`：テキスト入力欄の文字数を指定。入力欄の幅が変わる。初期値は20
* `maxlength`：入力欄に入力できる最大文字数
* `checked`：チェックボックス＆ラジオボタンを選択状態にする
* `readonly`：変更不可(選択可能)の状態にする
* `disable`：変更不可・選択不可の状態にする
* `src`：画像の送信ボタンの関連設定
* `width`：画像の送信ボタンの関連設定
* `weight`：画像の送信ボタンの関連設定
* `alt`：画像の送信ボタンの関連設定

### テキストフィールド

これは１行だけの文字入力域となる
nameは受信プログラムが項目を判断できるためのもの
sizeは入力可能な文字数、requiredは必須項目であることをブラウザに伝える
`<input type="text" name="...">`
`<input type="text" name="" size="" required>`

### パスワード

１行だけの文字入力域だが、入力したものは表示されない
`<input type="password" name="...">`

### チェックボックス

いくつでも選択できる
valueは、そのボタンが選択された時に、サーバ側に伝える値
checkedは、選択ずみを表す。属性値は指定しない
`<input type="checkbox" name="..." value="...">aaa`
`<input type="checkbox" name="..." value="..." checked>bbb`

### ラジオボックス

一つしか選べない
valueは選択されたときにプログラムに送る値
`<input type="radio" name="..." value="...">aaa`
`<input type="radio" name="..." value="..." checked>bbb`

### ラベル

文字列をフォームのinput要素に対応付ける
`<label for="...">...選択肢の文字列...</label>`
checkboxやradioだけなら、ボタンをクリックしなければならないが、labelを使えば、文字列をクリックしても選択できる
label要素の内容はインライン要素だけ。フォーム部品もインライン要素に分類される
forは、関連づける部品のidを入れる、省略時は要素内容中の部品となる

```html
<input type="radio" name="gender" id="otoko" value="male">
    <label for="otoko">男性</label>
</input>
```

### サブミットボタン

フォームに入力されたデータを送信するボタンを表示する
valueは、ボタン上に表示するテキスト
`<input type="submit" name="aaa" value="送信する">`

## その他

### 汎用のボタン

typeは、「submit」は送信ボタン(初期値)、
「reset」はリセットボタン、
「button」は汎用ボタンである
valueは、ボタン上に表示するテキスト
onclickは、クリックされた時に呼び出すJS関数
`<input type="button" value="ボタンa" onclick="aaa()">`

### ボタン要素

input要素のボタンとほぼ同じだが、内容の部分はボタン上に表示される、画像もOK
typeには、button,submit,resetが指定できる。省略するとsubmitになる。
`<button type="button" onclick="aaa()">...内容...</button>`

### テキストエリア

行数や文字数に制限がないテキスト入力域
colsは１行の文字数、rowsは行数
内容の部分は最初に表示される文字列で、ユーザが書き換え可能。省略可
**幅によりレイアウト崩れる可能性がある！！！**
`<textarea rows="4" cols="30" name="...">初期内容......</textarea>`

### ドロップダウンリスト

三角のボタンを押すと展開される選択肢から選ぶリスト
valueは、その選択肢が選択された時に、サーバ側に伝える値
optgroupで選択肢をグループ化
selectedで最初から特定の項目を選択状態にする

```html
<select name="...">
    <optgroup label=”関東”>
        <option value="..." selected>...内容...</option>
        <option value="...">...内容...</option>
        <option value="...">...内容...</option>
    </optgroup>
    <optgroup label=”関西”>
        <option value="...">...内容...</option>
        <option value="...">...内容...</option>
        <option value="...">...内容...</option>
    </optgroup>
    ...
</select>
```

メニューを選択できない状態にする場合
`<select name="..." disabled></select>`
これ以外に、
sizeで、メニューの縦幅を行数で指定できる。
multipleはで、複数選択可能にする
この２つの属性を１個指定すると、リストボックスになる

```html
<select name="..." size="5" multiple>
    <option value="..." selected>option A</option>
    <option value="...">option B</option>
    <option value="...">option C</option>
    <option value="...">option D</option>
    <option value="...">option E</option>
    <option value="...">option F</option>
    <option value="...">option G</option>
</select>
```

### 部品をグループ化

グループにタイトルをつけたい時、legendを使う。省略可

```html
<fieldset name="" disable>
    <legend>グループのタイトル</legend>
    <p>
        <label> 例：<button>abc</button> </label>
    </p>
</fieldset>
```
