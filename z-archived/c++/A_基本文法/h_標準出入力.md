# 標準出入力

## 簡単な例

### 出力

```c++
#include <iostream>
using namespace std;

int main(){
    cout << "Hello, World!\n";
    cout << "Hello, World!" << endl;
    cout << "Hello" << ", " << "World" << "!" << endl;

    cout << "変数numの値は" << num << "です\n";
}
```

### 入力

```c++
#include <iostream>
using namespace std;

int main(){
    int num1, num2;
    cin >> num1;
    cin >> num1 >> num2;

    string s;
    cin >> s;
    // 入力された１行を丸ごと受け取る
    getline(cin, s);
}
```

## 出入力の基礎

### C++の入出力の基礎

1. 標準入力：`std::cin`
2. 標準出力：`std::cout`
3. 標準エラー出力：`std::cerr`
4. バッファー付き標準エラー出力：`std::clog`、特殊な用途でしか使わない
cinやcoutなどは、全てiostreamヘッダーの中のテンプレートクラスの実体である。

* `std::basic_istream`   -->.  `std::cin`
* `std::basic_ostream`   -->.  `std::cout` `std::cerr` `std::clog`
* `std::basic_iostream`  入力＆出力のどちらもできる

### 書式設定された出力

* フラグを設定して表示方法を変える

    ```c++
    cout.flags()    // セットされているフラグを取得
    std::cout.unsetf(  std::ios::フラグ名  )    // セットしてあるフラグを外す
    std::cout.setf(  std::ios::フラグ名  )    // フラグをセット
    ```

    |       フラグ名       |                             効果                             |
    | :------------------: | :----------------------------------------------------------: |
    | std::ios::boolalpha  |         bool型を出力するときにtrueとfalseで出力する          |
    |    std::ios::oct     |                    数値を8進数で出力する                     |
    |    std::ios::dec     |                    数値を10進数で出力する                    |
    |    std::ios::hex     |                    数値を16進数で出力する                    |
    |  std::ios::showbase  | 数値を出力するとき、8進数なら0を、16進数なら0xを先頭につける |
    |   std::ios::fixed    |         浮動小数点数を固定表記(123.456000)で出力する         |
    | std::ios::scientific |               浮動小数点数を科学表記で出力する               |
    |    std::ios::left    |                       左詰めで出力する                       |
    |  std::ios::internal  |     中央揃え（何桁の中央に揃えるかは別の方法で指定する）     |
    |   std::ios::right    |                       右詰めで出力する                       |

* flags()関数とsetf()関数のオーバーロード

    ```c++
    flags(フラグの組み合わせ)
    setf(フラグ, フラグの組み合わせ)
    ```

    `flags(フラグ)`は、全てのフラグを、引数のフラグの組み合わせ通りに変更する
    `setf(フラグ, マスクの組み合わせ)`は、第２引数で指定したものを全部クリアして、２つの引数の和集合でフラグ設定する

    | よく使うフラグの組み合わせ |               概要                |
    | :------------------------: | :-------------------------------: |
    |  `std::ios::adjustfield`   | left, internal, rightの組み合わせ |
    |   `std::ios::basefield`    |     dec, hex, octの組み合わせ     |
    |   `std::ios::floatfield`   |   scientific, fixedの組み合わせ   |

    ```c++
    auto defaultFlag = std::cout.flags();

    std::cout.setf(std::ios::scientific);    // 科学表記に
    std::cout.setf(std::ios::hex, std::ios::basefield);    // フラグクリア後、16進数フラグをセット

    std::cout.flags(defaultFlag);    // デフォルトにリセット
    ```

* 幅、精度、充填文字の指定

    ```c++
    std::streamsize   width() const;
    std::streamsize   width(std::streamsize wide);    // デフォルトは右詰め
    std::streamsize   precision() const;
    std::streamsize   precision(std::streamsize prec);    // デフォルトは有効桁数６桁、整数部分優先
    char   fill() const;
    char   fill(char ch);    // デフォルトは空白が使用される
    ```

### 他のテキスト入出力関数

```c++
std::istream&   get(char* buf, std::streamsize num);
std::istream&   get(char* buf, std::streamsize num, char delim);

std::istream&   getline(char* buf, std::streamsize num);
std::istream&   getline(char* buf, std::streamsize num, char delim);
```

* ２引数の`get()`：
    bufに最大num-1文字を保存、必ず最後にヌル文字を付け加える
    改行文字をストリームに残すため、連続に使うと次は空となる
* ３引数の`get()`：
    第３引数で区切り文字を指定できる以外、２引数のと同じ
* ２引数の`getline()`：
    ２引数の`get()`とほとんど同じ
    改行文字をストリームから取り除く
* ３引数の`getline()`：
    第３引数で区切り文字を指定できる以外、２引数のと同じ

### 入出力 マニピュレーターの使用

一部のマニピュレーターは、`#include<iomanip>`する必要がある

```c++
std::cout << std::hex << 1234;    // 16進数に
std::cout.setf(std::ios::boolalpha);
std::cout << std::noboolalpha << true;    // 1
```

|             マニピュレーター             |            効果            |                           補足                            |
| :--------------------------------------: | :------------------------: | :-------------------------------------------------------: |
|              std::boolalpha              | 対応するフラグをセットする |         フラグをクリアするstd::noboolalphaもある          |
|     std::oct,   std::dec,   std::hex     | 対応するフラグをセットする |     oct, dec, hexのうちセットしなかったものは外される     |
| std::left,   std::internal,   std::right | 対応するフラグをセットする | left, internal, rightのうちセットしなかったものは外される |
|      std::fixed,   std::scientific       | 対応するフラグをセットする |   fixed, scientificのうちセットしなかったものは外される   |
|              std::hexfloat               | 浮動小数点数を16進数で表記 |                             -                             |
|              std::showbase               | 対応するフラグをセットする |          フラグをクリアするstd::noshowbaseもある          |
|           std::setfill(char c)           |     充填文字をcにする      |                             -                             |
|         std::setprecision(int p)         |       精度をpにする        |                             -                             |
|             std::setw(int w)             |      幅(桁)をw似する       |                             -                             |
