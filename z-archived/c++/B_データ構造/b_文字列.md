# 文字列

## 基本の使い方

### 要素の参照

```c++
str[i]
str.at(i)
```

### 挿入・削除・置換

* 文字列の末尾への結合：`+=` `append()`
* 指定位置への挿入：`insert()`
* 指定範囲の削除：`erase()`
* 指定位置の置換：`replace()`
コンストラクタと同様、様々な形式で引数を与えられる

### 主なメンバ関数

* `clear()`：空文字列に
* `empty()`：boolを返す
* `size()`：charの数を返す（漢字の文字数と違う）
* `length()`：charの数を返す（漢字の文字数と違う）
* `substr()`：部分文字列を返す。第１引数は開始位置、第２引数は文字数(省略すると最後まで)

## コンストラクタ

### 文字列の生成(コンストラクタ)

```c++
/*1*/  std::string str;
/*2*/  std::string s1 = "hello";        // hello
/*3*/  std::string str{"abcdefg", 4};   // abcd
/*4*/  std::string str = s1;            // コピー
/*5*/  std::string str(3, 'a');         // aaa
/*6*/  std::string str = {'a', 'b', 'c'};   // abc
/*7*/  std::string str{s1.begin()+1, s1.end()-1};   // ell
/*8*/  std::string str{s1, 1, 3};       // ell
```

インスタンス生成後に、文字列をセット(アサイン：丸ごと更新)する場合、`assign()`を使う
使い方は上とほぼ一緒

## 文字列と配列

### 文字列と配列の関係

文字列を「char型の配列」で扱うことができる

```c++
// 配列で初期化
char str01[] = {'H', 'e', 'l', 'l', 'o', '\0'};
char str01[6] = {'H', 'e', 'l', 'l', 'o', '\0'};

// ""で初期化
char str02[] = "World";
char str02[6] = "World";
```

注意：　char配列において、`" "`は初期化の時にのみ使える。後から`" "`で代入できない

```c++
cout << str01 << "\n";
```

## 文字列とポインタ

### 文字列とポインタの関係

文字列を「char型のポインタ」で扱うことができる

```c++
char *str03 = "Hello";
str03 = "Goodbye";

cout << str03 << "\n";
```

上のと違って、この場合は`" "`で初期化できるし、再代入もできる
（ポインタの指すアドレスにある文字列が変わる）

```c++
char *p;

p = "hello";
cout << p << "\n";      // hello
cout << *p << "\n";     // h
cout << &p << "\n";     // 0x7ffeea816940

p = "bye";
cout << p << "\n";      // bye
cout << *p << "\n";     // b
cout << &p << "\n";     // 0x7ffeea816940
```

## 数値・文字列間変換

### 数値・文字列間変換の仕方

整数へ変換する際に、第３引数に2~36の基数を指定できる(2進数〜36進数)
また、第２引数を指定する場合、変換できなかった文字のindexをそこに格納する
`std::stoul` `std::stol` などもある

```c++
std::to_string(数値)

std::stoi("-10")    // 文字列 --> int
std::stod("1.23")   // 文字列 --> double
```

* 高速な変換関数(C++17から)
    `#include <charconv>`

    ```c++
    std::to_chars()
    std::from_chars()
    ```

    引数や使い方など、詳しい内容を省略

## 文字列の検索

### 検索

```c++
str.find(文字列)
str.find(文字列, 開始位置)
// 後ろから前へ検索
str.rfind(文字列)
str.rfind(文字列, 開始位置)
// 特定の１文字を検索
str.find_first_of('a')  // 最初のa
str.find_last_of('a')   // 最後のa
str.find_first_not_of('a')  // 最初のaじゃない文字
str.find_last_not_of('a')   // 最後のaじゃない文字
```
