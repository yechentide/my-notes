# プロトタイプオブジェクト

## プロトタイプ

JavaScriptでは、全てのものがオブジェクト(またはそのインスタンス)である

そして、全てのオブジェクトは`prototypeオブジェクト`および`__proto__プロパティ`を持っている

* `__proto__`: これはポインタみたいなプロパティである。どこかの`prototype`を指している

* `prototype`: 子オブジェクトに継承させたいものをこの中に記述する。(インスタンスはこれを持たない)

    `prototype`もオブジェクトなので、その中にも`__proto__`プロパティがある

## プロトタイプチェーン

あるインスタンスから、メソッドを呼び出す時の話。

もしインスタンス自身にメソッドが見つからない場合、

そのインスタンスの`__proto__`プロパティから、どこかの`prototype`に辿り、その中で探す。

それでもない場合、さらにその`prototype`内の`__proto__`より、新たな`prototype`に辿り着く。

こうやって繰り返して、最終的に`__proto__ === null`となり、`undefined`が返される。

### つまり

インスタンスを生成するときに、オブジェクト自身の`prototype`に対する参照を、インスタンスの`__proto__`プロパティに格納する

このオブジェクトの`prototype`を更新すれば、全てのインスタンスに影響を与える

自作コンストラクターにメソッドを追加する際に、`prototype`に追加すれば、生成済みのインスタンスでも使えるようになる

## 具体的な関係

```javascript
Object.prototype.__proto__ === null

Object.__proto__ === Function.prototype

// 生成したインスタンスの__proto__
(new Object()).__proto__ === Object.prototype



Function.prototype.__proto__ === Object.prototype

Function.__proto__ === Function.prototype

// 生成したインスタンスの__proto__
(new Function()).__proto__ === Function.prototype



Array.prototype.__proto__ === Object.prototype

Array.__proto__ === Function.prototype

// 生成したインスタンスの__proto__
(new Array()).__proto__ === Array.prototype
```

### 自作コンストラクタで生成したインスタンス

例えば下のようなコンストラクタでインスタンスを生成したとする

```javascript
function Person(name){
    this.name = name;
}
const person = new Person("yuki");
```

関係性は以下のようになる

```javascript
person.__proto__ === Person.prototype
Person.prototype.constructor.__proto__ === Function.prototype
Person.prototype.constructor.prototype === Person.prototype
```

### `Object.create()`で生成したインスタンス

`Object.create()`を使ってもインスタンスを生成できる

```javascript
const anotherPerson = Object.create(person);
```

このとき、`anotherPerson.__proto__`が参照しているのは、`person`インスタンスとなる

つまり、`person`に何かを追加すれば、`anotherPerson`もそれを使える(`anotherPerson.__proto__`で探す)

`Person.prototype`に追加して、`anotherPerson`でも使える(`anotherPerson.__proto__.__proto__`で探す)

### `=`で生成したインスタンス

```javascript
const person01 = person
```

`Object.create()`と違って、`person01.__proto__ === person.__proto__ === Person.prototype`

## コンストラクタ

コンストラクタ関数の値は、`prototype`内の`constructor`プロパティに保存されている

インスタンスからは、`インスタンス.constructor`でアクセスできる(実際は`インスタンス.__proto__.constructor`)

こういうことを利用して、インスタンスからコンストラクタを入手し、それで新しいインスタンスを生成できる

```javascript
const person02 = new anotherPerson.constructor("abc");
person02.__proto__ === Person.prototype    // true
```

`anotherPerson.__proto__`が参照しているのは、`person`インスタンス

しかし、`constructor`というプロパティは`anotherPerson`にも`person`にもなく、

結局`Person.prototype`まで探しに行ったので、`person02.__proto__ === Person.prototype`がtrueとなる

### constructor.name

`インスタンス.constructor.name`で、コンストラクタ関数の関数名を取得できる

## プロトタイプの変更

コンストラクタ関数の`prototype`プロパティを変更すれば、

コンストラクタから作成されたすべてのオブジェクトインスタンスで使用可能になる

ただし、`prototype`に追加した関数内部で、メンバのプロパティを利用したいなら、

`this`キーワードを使う必要がある