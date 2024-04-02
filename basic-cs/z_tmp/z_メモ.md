# アルゴリズム・メモ

## 2 Sum

### 問題

- 入力: 整数の配列、ターゲットの整数
- 出力: 和がターゲットになる２つの整数のindex

### 解法1   O(n)

配列内の整数をキー、そのindexを値として辞書に保存  
各整数をそれぞれ見ていく

### 解法2   O(n), ソートするならO(nlogn)

配列を先にソートする必要がある  
両端にそれぞれポインタを用意し、その和の大きさによってポインタを動かす

### 関連問題

[3 Sum & 4 Sum](https://github.com/raywenderlich/swift-algorithm-club/tree/master/3Sum%20and%204Sum)

## A-Star

A*探索アルゴリズムは、**グラフ探索アルゴリズム**の１つ。グラフ上でスタートからゴールまでの道を見つける。  
A*アルゴリズムは、[ダイクストラ法](https://ja.wikipedia.org/wiki/ダイクストラ法)を推定値付きの場合に一般化したもので、h が常に0の場合は、もとのダイクストラ法と同じ。  

### 参考文献

[A* - Wikipedia](https://ja.wikipedia.org/wiki/A*)  
[よくわかるA*(A-star)アルゴリズム (Unity2Dのサンプルコードつき) - Qiita](https://qiita.com/2dgames_jp/items/f29e915357c1decbc4b7)  
[ダイクストラ法とA*(A-star)探索を3Dメッシュ上で行いThree.jsで可視化をしてみた - Qiita](https://qiita.com/Raysphere24/items/5892cd8e623d20fcb308)

## AVL Tree

「どのノードの左右部分木の高さの差も1以下」という条件を満たす二分探索木のこと。平衡二分探索木の1つである。  
木に対する操作によって条件を満たさないノードが発生しても、[回転](https://ja.wikipedia.org/wiki/木の回転)と呼ばれる操作を行うだけで木をAVL木に再構成でき、平衡を維持できる。
