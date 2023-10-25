# メモ

## jsを読み込む

* HTMLファイル内に書く場合
  
  ```javascript
    document.addEventListener("DOMContentLoaded", function() {
          ...
    });
    // onloadでも良い
  ```

* 別ファイルに書く場合
  
  ```javascript
    <script src="script.js" defer></script>
  ```

## async と defer

[async / defer](https://qiita.com/phanect/items/82c85ea4b8f9c373d684)

スクリプトのブロッキングの問題を回避するためのモダンな機能である

* `async`: 

## アラート

```javascript
// 入力欄つき
let name = prompt('あなたの名前は？');
// 単純なアラート
alert('こんにちは、' + name + 'さん。初めまして！');
```
