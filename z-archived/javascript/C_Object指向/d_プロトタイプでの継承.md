# プロトタイプでの継承

## コンストラクタ関数の定義

下のコードはPersonコンストラクタである

```javascript
function Person(first, last, age, gender, interests) {
	this.name = {
		first,
		last
	};
	this.age = age;
	this.gender = gender;
	this.interests = interests;
};
Person.prototype.greeting = function() {
	alert('Hi! I\'m ' + this.name.first + '.');
};
```

Personコンストラクタを継承して、teacherコンストラクタを定義する

```javascript
function Teacher(first, last, age, gender, interests, subject) {
	Person.call(this, first, last, age, gender, interests);

	this.subject = subject;
}
```

引数なしのコンストラクタを継承する場合、`call()`の引数は`this`のみとなる

## prototypeとconstructorの設定

上の段階では、`Teacher.__proto__ === Function.prototype`がtrueであり、

`Person.prototype`に定義されている`greeting`メソッドを使えない。

ここで、`Person.prototype`を参照するようなオブジェクトを新しく作って、それを`Teacher.prototype`にすれば良い

```javascript
Teacher.prototype = Object.create(Person.prototype);
```

しかしこうすると、`Teacher.prototype.constructor`の値は`Person.prototype.constructor`の値と同じくなる

`Teacher.prototype.constructor`を再設定する必要がある

```javascript
Teacher.prototype.constructor = Teacher;
```

