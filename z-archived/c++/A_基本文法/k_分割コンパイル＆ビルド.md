# 分割コンパイル＆ビルド

## ファイルの分割

* myfunc.h

    ```c++
    // 関数のプロトタイプ宣言
    戻り値の型 関数名(型 仮引数名, 型 仮引数名, ...);
    ```

* myfunc.cpp

    ```c++
    #include "myfunc.h"
    // 関数の定義
    戻り値の型 関数名(型 仮引数名, 型 仮引数名, ...){
        // ...
        return 戻り値;
    }
    ```

* Sample.cpp

    ```c++
    #include "myfunc.h"
    using namespace std;

    int main(){
        // 関数を呼び出す
        return 0;
    }
    ```

変数のスコープ：

* ローカル変数：ブロック内
* グローバル変数＆関数：全てのファイル
* `static`付きグローバル変数＆関数：ファイル内

## extern変数

変数宣言が暗黙のうちに変数定義まで含んでいる。
コンパイラは変数定義を見つける度に、変数の実体を作ってしまう
そのため、複数のファイルで同じ変数名のグローバル変数を宣言すると、
リンクの際に複数の同じ名前の実体ができてしまい、区別できなくなる。
**グローバル変数宣言**の際に、`extern`をつければ、
このグローバル変数の実体はどこか別のところで定義される、とコンパイラに知らせる。

* Main.cpp

    ```c++
    #include <iostream>

    extern int value;
    void showValue();    // 他のファイルにある関数のプロトタイプ宣言

    int main(){
        cout << &value << endl;    // 0x123
        cout << value << endl;    // 42
        value = 0;
        showValue();    // 0x123, 0
    }
  ```

* Variable.cpp

    ```c++
    #include <iostream>

    int value = 42;

    void showValue(){
        cout << &value << endl;
        cout << value << endl;
    }
    ```

## inline関数

* inline展開
    プロトタイプ宣言のところに定義を記述すればinline関数になる
* inline指定
    ヘッダファイルで宣言と定義の両方がされた関数は、
    このヘッダファイルをincludeされた複数のソースファイルにも関数の実体ができてしまい、ODR違反となる。
    インライン指定をすればこの問題を解決できる

    ```c++
    inline void func1();
    void func1(){ /**/ }

    inline void func2(){ /**/ }
    ```

* 自動インライン化
    １行に収まるような小さな関数は、クラスの宣言に定義を書いても良い

## 名前空間

名前空間の指定がないものは、暗黙的に全てがグローバル名前空間に属する

```c++
namespace 名前空間名{
    // ...
}
```

ある名前空間にあるものを使いたい場合、スコープ解決演算子`::`で所属する名前空間も一緒に指定すべき

```c++
名前空間名::関数
名前空間名::クラス 変数名;
```

名前空間のネストもできる。ライブラリなどで機能管理するときに役に立つ

```c++
namespace A{
    namespace B{
        void func();
    }
}
A::B::func();
```

```c++
// 省略記法
namespace A{
    // ...
}
namespace A::B{
    void func();
}
```

名前空間の省略

```c++
using namespace 名前空間名;
using namespace 名前空間名::ネスト名前空間名;
```

特定の名前を導入

```c++
using 名前空間名::関数;
using 名前空間名::クラス;
```

名前空間の別名

```c++
namespace 別名 = 名前空間名;
```

グローバル名前空間からの絶対パスで指定

```c++
::名前空間名::関数;
```

他のソースファイルから呼び出されたくない場合、無名名前空間を使う

```c++
namespace{
    // ...
}
```

## リンケージ

C++からC言語の関数を呼ぶ場合と、C言語からC++の関数を呼ぶ場合に、リンケージを使う
指定がない場合、デフォルトの`extern "C++"`となる
Cリンケージを指定した関数はCからも使えないとダメなので、C++の機能が一部制限される。特にオーバーロードが使えなくなる

```c++
extern "C" 戻り値の型 関数名(引数);
```

ヘッダファイルでリンケージを一括指定できる

```c++
extern "C" {
    // ...
}
```

## プリプロセッサ

* プリプロセッサ命令
    プリプロセッサ命令はハッシュ`#`で始まる命令である
    １行において、ハッシュの前は空白以外許されない。ハッシュと後ろの命令の間に空白を入れて良い
    `#include`命令：ヘッダーファイルを読み込むプリプロセッサ命令
* マクロ
    識別子を別の識別子列に置き換える命令
    マクロが置き換える文字列を省略すると、マクロより下のその識別子を取り除くことになる

    ```c++
    #define マクロ名
    #define マクロ名 置き換えられる識別子
    ```

    マクロはプリプロセッサで処理されるため、この段階ではまだ名前空間など認識されていない。
    したがってマクロ同士の衝突が起こりやすくて、他のマクロと重複しないように、長めの名前にするのが一般的
    マクロの定義を消すには

    ```c++
    #undef マクロ名
    ```

    引数を与えて、置き換える文字列を変更できる関数形式マクロもある

    ```c++
    #define マクロ名(引数)
    #define マクロ名(引数) 置き換えられる識別子
    ```

    ```c++
    void hello();

    #define id(name) name   // 受け取った引数で置き換える

    int main(){
        id(hello)();    // hello()を呼び出す
    }
    ```

* マクロの結合と展開
    マクロ定義の中で他のマクロも使える。その際、引数より先に演算子が展開される

    ```c++
    #define make(a,b) a b
        make(left,right)()    // left right()
    #define make(a,b) ab
        make(left,right)()    // ab()
    #define make(a,b) a ## b
        make(left,right)()    // leftright()

    // 識別子を文字列に置き換える(#直後の１つだけが文字列に)
    #define make(a) #a
        make(hello)           // "hello"
    ```

* 特定の条件でプログラムを有効にしたり無効にしたりできる。`#if` `#elif` `#else` `#endif`を使う
    条件式として使えるものは、整数(0がfalse)、計算式、`defined(マクロ名)`(1か0)、マクロ名
    例えば`__LINE__` `__FILE__` `__cplusplus`などの暗黙的に定義される識別子も使える

    ```c++
    #include <iostream>
    using namespace std;

    #define PLUS(a,b) (a+b)
    #define HOGE

    int main(){
    #if PLUS( defined(HOGE),0 )
        cout << "PLUS( defined(HOGE),0 ) ---> true" << endl;
    #else
        cout << "PLUS( defined(HOGE),0 ) ---> false" << endl;
    #endif
    }
    ```

    ```c++
    #if defined(__cplusplus)
    extern "C"{
    #endif

        // C言語の宣言など

    #if defined(__cplusplus)
    }
    #endif
    ```

* 多くの場合で`#if`と`defined()`を組み合わせて使うため、合体させて`#ifdef`と`#ifndef`を用意した。

    ```c++
    #ifdef マクロ名     // 定義されていたらtrue
    #ifndef マクロ名    // 定義されていなかったらtrue
    ```

* インクルードガード：ライブラリの重複読込を回避するためのテクニック
    １回目のインクルードで、`#define`でマクロ定義したので、２回目以降は無効となる

    ```c++
    #ifndef UNIQUE_IDENTIFIER
    #define UNIQUE_IDENTIFIER
    // ヘッダーファイル本体
    #endif
    ```
