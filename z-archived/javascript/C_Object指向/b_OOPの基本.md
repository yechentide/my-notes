# OOPの基本

## コンストラクター関数①

```javascript
function createNewPerson(name) {
    const obj = {};
    obj.name = name;
    obj.greeting = function() {
        alert('Hi! I\'m ' + obj.name + '.');
    };
    return obj;
}
```

この関数でインスタンスを生成できる

```javascript
const salva = createNewPerson('Salva');
salva.name;
salva.greeting();
```

## コンストラクター関数②

```javascript
function Person(name) {
    this.name = name;
    this.greeting = function() {
        alert('Hi! I\'m ' + this.name + '.');
    };
}
```

これはJavaなどのコンストラクターと似ている

```javascript
let person = new Person('Bob');
person.name
person.greeting()
```

## Object() コンストラクター

```javascript
let person1 = new Object();
person1.name = 'Chris';
person1['age'] = 38;
person1.greeting = function() {
    alert('Hi! I\'m ' + this.name + '.');
};
```

または

```javascript
let person1 = new Object({
    name: 'Chris',
    age: 38,
    greeting: function() {
        alert('Hi! I\'m ' + this.name + '.');
    }
});
```

## オブジェクトのコピー

少数のインスタンスのみを生成するとき、

わざわざコンストラクターを作るのがめんどくさいので、

`Object.create(インスタンス)`でコピーすれば良い

```javascript
let person1 = Object.create(person);
person1.name;
person1.greeting();
```
