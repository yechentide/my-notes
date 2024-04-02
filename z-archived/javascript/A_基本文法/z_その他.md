# その他

## 未分類

### Webページで使う

1. HTMLファイルに書く
    `<script> ...... </script>`
2. 外部ファイルを読み込む
    `<script type="text/javascript" src="......">   </script>`

### js無効のブラウザへの対処

```html
<script>
// <! --
    .....
// -- >
</script>
<noscript>無効です</noscript>
```

### strictモード

より厳格なエラーチェックを行うモードで、プログラム全体or関数単位で設定できる。
`"use strict";` をプログラムor関数の先頭に書くとこのモードになる

## 分割代入

```javascript
let data = [12, 34, 45, 56, 89];
let [data1, data2, data3, data4, data5] = data;

let food = {name: "コンソメパンチ", price: 240, size: "やや大きめ"};
let {price, size, memo = "パーティー用"} = food;
console.log(price); // 240
console.log(size);  // やや大きめ
console.log(memo);  //パーティー用
```

## スプレッド演算子

```javascript
let [f, ...r] = [1,2,3]     // f=1, r=[2,3]
function func(first, second, ...rest){} // 可変長引数

const a1 = [1,2,3]
const a2 = [...a1,4,5,6]        // [1,2,3,4,5,6]
```

## タイマー

```javascript
// 一定時間ごとに指定した関数を実行させる。第２引数の単位はミリ秒
setInterval(func01, 1000);
// setInterval()で開始したタイマーを停止する
let timer = setInterval(func01, 1000);
clearInterval(timer);
```

## リアルタイム処理

```javascript
// 10msec間隔で関数を実行
setInterval(関数名,10)

// 上の記述方法と同じだが、こっちは関数内に書く
setTimeout(関数名,10)

/* 違い */
// setIntervalは、呼出先の関数の処理が終わってなくても、時間がくると呼び出す。
// よほど思い処理でないと変わらないのだが、setTimeoutの方が安全
```

## モジュール

ES6では、モジュールがファイル単位となる

* 公開するファイル (shop.js)

    ```javascript
    export class Table1 {}
    export class Table2 {}
    class Table3 {}
    class Table4 {}
    ```

* 引用するファイル
    モジュールがデフォルトで公開するクラス/関数を指定できる

    ```javascript
    import {Table1,Table2} from 'shop'
    ```

* 公開するファイル (shop2.js)

    ```javascript
    export default class Table {}
    ```

* 引用するファイル

    ```javascript
    import Table from 'shop2'
    ```
