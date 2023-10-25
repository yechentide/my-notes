# 配列

## 配列の説明

### リスト

```javascript
["A", "B", "C", "D"];
```

### 定義方法

```javascript
// ① コンストラクタで生成
let arr = new Array();
let arr = new Array(1, 2, 3, 4);
// ② 変数を宣言し、リストを代入
let arr = ["B", "C", "D"];
```

### 使い方

* 要素のアクセス
    `arr02[0]`
* 要素の追加

    ```javascript
    // 要素を最初に追加:  A B C D
    arr02.unshift("A");
    // 要素を最後に追加:  A B C D E
    arr02.push("E");
    // 要素を最後に追加:  A B C D E F ?????????
    arr02[arr02.length]="F";
    ```

* 要素の削除
    削除したい要素をundefinedにする
    1. 途中の要素の削除

        ```javascript
        delete arr[2];
        // lengthは変わらない
        ```

    2. 連続する要素の複数削除

        ```javascript
        arr.splice(1,2)
        // lengthは変わる
        // arr[1]から2要素を消除して要素番号を詰める
        // 実際は途中を抜いて別の配列を作る
        ```

### 連想配列

別の言語で、マップや辞書型などとも呼ばれる
index番号ではなく、自分でつけた名前で値を管理する配列

* 定義方法

    ```javascript
    let Name = {   key1:value1   ,   key2:value2   /* , ..........*/   };
    let Name = {   "a b":value1   ,   "a c":value2   /* , ..........*/   };
    ```

* 使い方

    ```javascript
    Name[key2] = value;
    Name.key2 = value;     /*keyに空白がなければ*/
    ```

### 配列の格納するもの

配列に配列or関数を代入できる

* 代入

    ```javascript
    var arr = new Array();
    arr["total"] = function(){ /* */ };
    ```

* 使い方

    ```javascript
    arr["total"]();
    arr.total();     /*keyに空白がなければ*/
    ```

### 配列のメソッド

1. `unshift()`
2. `shift()`
3. `push()`
4. `pop()`
5. `concat()`
6. `join()`
7. `indexOf()`
8. `lastIndexOf()`
9. `splice()`
10. `slice()`
11. `sort()`
12. `reverse()`
13. `map()`

### その他

```javascript
var aaa = "月,火,水,木,金,土,日".split(",");
console.log(aaa);
console.log(aaa.toString());
console.log(aaa.join("-"));
```
