# タプル

## タプルの説明

### ２つの組：std::pair

```c++
#include <utility>

std::pair<char, int> p{'a', 1};
cout << p.first << endl;
cout << p.second << endl;
```

### 任意個数の組：std:tuple

```c++
#include <tuple>

std::tuple<int, char, char> t{1, 'a', 'A'};
cout << std::get<0>(t) << endl;
cout << std::get<1>(t) << endl;
cout << std::get<2>(t) << endl;
```

### 構造化束縛

C++17からの機能
`std::get()`は少し面倒なので、タプルを展開して個別の変数として扱うようにする。
タプル(pair&touple)だけじゃなく、配列や簡単な構造体でも展開できる
`auto`と型推論の右辺が必須

```c++
// 各要素をコピー
auto [変数, ...] = タプルなど;
// 各要素への参照
auto& [変数, ...] = タプルなど;
```
