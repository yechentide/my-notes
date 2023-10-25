# オブジェクトの基本

## 簡単なオブジェクト

```javascript
var person = {
    name: ['Bob', 'Smith'],
    age: 32,
    greeting: function() {
        console.log('Hi! I\'m ' + this.name[0] + '.');
    }
};

person.name[0];
person['name'][0;]
```

## サブ名前空間

```javascript
var person = {
    name : {
        first: 'Bob',
        last: 'Smith'
    },
    age: 32,
    greeting: function() {
        console.log('Hi! I\'m ' + this.name[0] + '.');
    }
};

person.name.first;
person['name']['first'];
```

## オブジェクトのメンバを追加

存在しないキーで追加すれば良い

```javascript
person['eyes'] = 'hazel';
person.farewell = function() { alert("Bye everybody!"); }
```
