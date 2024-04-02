# 列挙体

## 列挙体の説明

### What is 列挙体

列挙体(enum)は、複数の値を型として定義する
switch文を使って場合分け処理をできる
これを使えば、読みやすいコードを記述できる場合がある
[enumとenum class](https://qiita.com/sorahako0515/items/5e7cf724b8ce4f004aae)
c言語から引き継いだ`enum`と、c++で拡張された`enum class`両方ある
c言語でも使うソースコードでなければ、出来るだけ`enum class`を使おう

### 宣言

列挙型の各値はEnum Valueという値である。
Enum Valueに整数や文字列などのRaw Valueを代入できる
設定値を省略すると、自動的に0からカウントアップされる。
最初の値を1にすると、次は2, 3 ......
連番である必要がなく、途中から値を変更できる

```c++
enum class 列挙型名 {識別子, 識別子, 識別子, ...};
enum class 列挙型名 {
    識別子,         // 0
    識別子 = 100,   // 100
    識別子,         // 101
    ...
};

enum class Week {SUN, MON, TUE, WED, THU, FRI, SAT};
```

### 使い方

```c++
Week w;
w = Week::SUN;
switch(w){
    ......
}

cout << int(Week::TUE);     // 2
```

## enumとenum class

### enum

スコープを持たない列挙型
暗黙的にint型になる

```c++
enum SampleEnum : int{
    se1,
    se2
};

void Function(){
    auto val1=se1;
    auto val2=SampleEnum::se2;
    cout << "val1:" << val1 << ",val2:" << val2 << endl;
}
```

### enum class

明示的なキャストが必要になる
enum class 名を明示的なスコープとして指定する必要がある

```c++
enum class SampleEnumClass : int{
    sec1,
    sec2
};


void Function(){
    // auto val1=sec1;  // !コンパイルエラー
    auto val2=SampleEnumClass::sec2;
    std::cout << "val2:" << static_cast<int>(val2) << std::endl;
}
```

### 共通点

1. enum (class) `hoge` の後に規定型を指定できる
2. 前方宣言ができる

### 比較

|                            | 例                       | enum                                                 | enum class                                |
| -------------------------- | ------------------------ | ---------------------------------------------------- | ----------------------------------------- |
| 規定型の指定               | `enum SampleEnum : int`  | ◯                                                    | ◯                                         |
| 前方宣言                   | `enum SampleEnum : int;` | ◯                                                    | ◯                                         |
| 列挙への直接的なアクセス   | `auto = se1;`            | ◯                                                    | ✖️                                         |
| 列挙への直接的なアクセス例 |                          | SampleEnum v = se1;  SampleEnum v=SampleEnum:se1; | SampleEnumClass v = SampleEnumClass:sec1; |
| 値をint型としてアクセス    |                          | int i=v;                                             | int i = static_cast&lt;int&gt;(v)               |
