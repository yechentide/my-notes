# Class構文

ECMAScript 2015から、JavaやC++と似たような構文でクラスを記述できる

## クラス構文

```javascript
class Person {
	constructor(first, last, age, gender, interests) {
		this.name = {
			first,
			last
		};
		this.age = age;
		this.gender = gender;
        this.interests = interests;
	}
	
    // クラスメソッド
	greeting() {
		console.log(`Hi! I'm ${this.name.first}`);
	};
}
```

```javascript
let han = new Person('Han', 'Solo', 25, 'male', ['Smuggling']);
han.greeting();			// Hi! I'm Han
```

## クラス構文による継承

```javascript
class Teacher extends Person {
	constructor(first, last, age, gender, interests, subject, grade) {
		super(first, last, age, gender, interests);

		// 科目と学年は教師によって決まっている
		this.subject = subject;
		this.grade = grade;
	}
}
```

## getter と setter

`subject`のような識別子は、getterとsetterの識別子として使いたいので、

クラスの内部では、別の名前で値を保存する必要がある

ここでは`_subject`としている

```javascript
class Teacher extends Person {
	constructor(first, last, age, gender, interests, subject, grade) {
		super(first, last, age, gender, interests);
		// 科目と学年は教師によって決まっている
		this._subject = subject;
		this.grade = grade;
	}

	get subject() {
		return this._subject;
	}

	set subject(newSubject) {
		this._subject = newSubject;
	}
}
```

使い方：

```javascript
let snape = new Teacher('Severus', 'Snape', 58, 'male', ['Potions'], 'Dark arts', 5);

console.log(snape.subject)			// Returns "Dark arts"
snape.subject="Balloon animals" 	// Sets _subject to "Balloon animals"
console.log(snape.subject) 			// Returns "Balloon animals"
```

