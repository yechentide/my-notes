# TypeScript

## 基本

### 型

`boolean` `number` `string` `any`

### 型のエイリアス

```typescript
type 別名 = 型;
```

### 変数＆定数

```typescript
let 変数名: データ型 = 初期値;
let 変数名: データ型;
let 変数名 = 初期値;

let 変数名;
let 変数名: any;

const 定数名 = 値;
```

### 列挙型

```typescript
enum Week {
   SUN,
   MON,
   THU
}
let m: Week = Week.MON;
```

### 配列

```typescript
let 変数名: Array<型> = [初期値, ...];
let 変数名: Array<型> = [];
let 変数名: 型[] = [初期値, ...];
let 変数名: 型[] = [];
let 変数名 = [初期値, ...];

let 変数名: any[] = [];
let 変数名 = [];
```

* 要素の追加：
  1. 末尾に追加：`push()`、複数の引数可
  2. 先頭に追加：`unshift()`、複数の引数可
* 要素の削除：
  1. 末尾を削除：`pop()`、削除された値を返す
  2. 先頭を削除：`shift()`、削除された値を返す
  3. 任意の順番で削除：`splice(開始位置,削除する数)`、第３引数から、入れ替わる値を指定できる
* 要素の参照：`配列名[index]`
* 要素の個数：`配列名.length`
* 配列の列挙：
  1. 要素＆index：`forEach()`
  2. index：`for in`
  3. 要素：`for of`
* ソート：`sort()` `reverse()`

### マップ(連想配列)

```typescript
let 変数名 = new Map<キーの型, 値の型>();
let 変数名 = new Map<キーの型, 値の型>(  [key,value],[key,value]...  );
let 変数名 = new Map(  [key,value],[key,value]...  );
let 変数名: Map<キーの型, 値の型> = new Map();

let 変数名: Map<any,any> = new Map();
let 変数名 = new Map();
```

* 要素の追加：`変数名.set(key, value)`
* 要素の削除：
  1. １つ：`変数名.delete(key)`
  2. 全て：`変数名.clear()`
* 要素の参照：
  1. `変数名.get(キー)`
  2. `変数名.has(キー)`：boolean
* 要素の個数：`size()`
* 要素の列挙：`forEach()` `for of`。`for in`は使えない

### 関数

オーバーロードが可能

```typescript
// 普通の関数
function 関数名(引数:型) : 戻り値の型 {
   // ...
}
// 匿名関数
let 関数の変数名 = function (引数:型) : 戻り値の型 {  };
// ラムダ式
let ラムダ式の変数名 = (引数:型) : 戻り値の型 => {  };      // return必要
let ラムダ式の変数名 = (引数:型) : 戻り値の型 => １行の式;   // return不要
```

### Generics

C++のテンプレート関数と似ている

```typescript
function 関数名<T>(引数:T) :戻り値の型 {
   switch( typeof 引数){   }
   // ...
}
```

### Tuple

```typescript
let 変数名: [型, 型, ...] = [値, 値, ...];
let 変数名: [型, 型, ...];
変数名[index];
```

### Null許容型(Nullable型)

```typescript
変数名 || 値がnullの時に使う値
let 変数名: 型 | null;
let 変数名: 型 | undefined;
function 関数名(引数名?:型){  }     // オプショナル引数を省略すると、undefinedとなる
```

## クラス

### クラス定義

```typescript
class クラス名{
   プロパティ:型;
   メソッド(引数:型) : 戻り値の型{}
   static メソッド(引数:型) : 戻り値の型{}
}

let 変数名 = new クラス名();
クラス名.静的メソッド(引数);
```

### オブジェクトの初期化

```typescript
// プロパティの初期値指定
class クラス名{
   プロパティ:型 = 初期値;
}
```

```typescript
// コンストラクタ
class クラス名{
   constructor(引数:型){
      // ...
   }
}
```

```typescript
// オブジェクトリテラルによる初期化
let 変数名 = {
    プロパティ: 値;
    関数名: function(){}
};
```

### アクセス修飾子

`public` `private` `protected`

### getterとsetter

```typescript
class Person{
   private age:number;
   get age():number {return this.age;}
   set age(y:number) {this.age = y<0 ? 0 : y;}
}
```

### 継承

オーバーライドが可能

```typescript
class クラス名 extends クラス {}
```

## インターフェース

### 定義

```typescript
interface 名前{
   プロパティ:型;
   メソッド(引数:型) : 戻り値の型;
}
```

### 実装

```typescript
class クラス名 implements インターフェース名 {
    // ...
}
```
