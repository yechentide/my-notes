# 定数＆変数&演算子

## 定数

### 定数宣言

ES6から定数を宣言できるようになった
**宣言と同時に初期化する**

```javascript
const Name = value;
```

## 変数

### 変数宣言

```javascript
var x=2;
var pi=3.14;
var x=true;
let y=false;
let nname="Bill Gates";
let answer='Yes I am!';
```

var = variable(変えられる)
初期化されていない変数の値は、undefinedとなる
宣言の前に変数を使うと、その変数がグローバル変数となる
varで宣言した変数は、宣言場所によって、グローバル変数orローカル変数になる
なるべくグローバル変数を使わないこと！
ES6から`let`で変数を宣言できるようになった

### letとvarの違い

```javascript
let a = 10;      var b = 10;
```

どちらも変数を宣言しているが、変数のスコープが違う。
letはブロックスコープで、宣言をした{ }内のみ有効
varは関数スコープで、宣言をした関数内のみ有効

## 定数＆変数のルール

### 変数名の命名規則

* アルファベット
* 数字
* Unicode文字（漢字など）
* `_` `$`

先頭文字は数字がNG 予約語を使用禁止

### 変数名の記法

* キャメル記法
    userName
* パスカル記法
    UserName
* アンダースコア記法
    user_name

## 演算子

### 条件演算子

```javascript
変数  =  条件  ?  値1  :  値2  ;
// 条件がtrueの時は値1を返す、falseの時は値2を返す
```

### 指数

5の２乗 = `5**2`

### `==`と`===`の違い

```javascript
// 両方とも値が等しいかどうかを判断する比較演算子だが、「===」の方が厳密
"123" == 123        // true
"123" === 123       // false

//  !=と!==についても同じ
"123" != 123        // false
"123" !== 123       // true

// プリミティブ型とオブジェクト型の比較
// プリミティブ型：値を比較
// オブジェクト型：同じインスタンス稼動かを比較
var o1 = new String("JavaScript");
var o2 = new String("JavaScript");
var o3 = o2;
o1 == o2        // false
o1 === o2       // false
o3 == o2        // true
o3 === o2       // true
// プリミティブ型とオブジェクト型を比較
var s1 = "JavaScript";
var o1 = new String("JavaScript");
o1 == s1        // true
o1 === s1       // false
```

### 型チェック

```javascript
typeof(引数)
// 引数がオブジェクト型の時は"object"を、文字列型の時は"string"を、
// 数値型の時は"number"を、真偽値型の時は"boolean"を戻す
var Name = typeof value;
```
