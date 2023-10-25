# ファイル操作

## 全体の説明

### ファイルの入出力

* ファイルストリーム
    1. `std::ifstream in`：入力用
    2. `std::ofstream out`：出力用
    3. `std::fstream inout`：両用
* ファイルを開くには、コンストラクタにファイル名を渡すか、オブジェクト作成後`open()`を使う
    閉じるには`close()`を使う

    ```c++
    #include <fstream>
    int main(){
        std::ofstream out;
        out.open("aaa.txt");
        out << "Line 1" << std::endl;
        out.close();

        std::ifstream in{"aaa.txt"};
        std::string line;
        std::getline(in, line);    // ファイルから１行入力
        in >> line;    // ファイルから１行入力
        in.close();
    }
    ```

* ファイルオープンのモードフラグ
    モードはどのような用途・状態でオープンするかを表す
    `open()`の第２引数、またはコンストラクタの第２引数で渡す。
    or演算子 `|` で複数組み合わせることができる
    途中でモードを変えたいとき、開き直す必要がある

    |   モードフラグ   |                   意味                   |
    | :--------------: | :--------------------------------------: |
    |   std::ios::in   |           入力可能なように開く           |
    |  std::ios::out   |           出力可能なように開く           |
    |  std::ios::app   |  **常に**出力がファイルの末尾に追加する  |
    |  std::ios::ate   | **開くときに**ファイルの末尾にシークする |
    | std::ios::trunc  |        ファイルを空にしてから開く        |
    | std::ios::binary |           バイナリモードで開く           |
* ファイルオープンの失敗
    `is_open()`で確認できる。真理値を返す
* ファイルの終端
    ファイルの終わり＝終端(EOF)。
    ファイルに出力する時に自動的に伸びていくので気にしなくて良い。
    ファイルから入力する時に、`eof()`で確認すべき。`eof()`は真理値を返す
* 書式不定のバイナリ入出力
    バイナリモードで開くことをおすすめ。
    様々な方法で書式を持たないファイル出入力があるが、その中一番低水準なものは`get()` `put()`。
    １バイトだけ入力/出力できる
    `get()`は変数に代入するのに対して、`peek()`は戻り値として返す
    `get()`や`peek()`はint型の値を返すのに、なぜchar型を使うのか？
    それはこの２つの関数では`eof()`関数を使えないので、char型では収まらないEOFとして返すことで判別している

    ```c++
    std::ofstream out{"aaa.bin", std::ios::binary};
    out.put(1);
    out.put(2).put(3);
    out.close()

    char a, b, c;
    std::ifstream in{"aaa.bin", std::ios::binary};
    in.get(&a);
    in.get(&b).get(&c);
    a = in.peek();
    ```

    １引数の`get()`と引数なしの`get()`

    ```c++
    // ファイルの中身は１行：　Hello, EOF
    std::ifstream in{"aaa.txt"};

    char c;
    if(!in.get(c).eof()){  }

    int ci = in.get();
    if(ci!=EOF){  }
    ```

* ブロックのバイナリ入出力

    ```c++
    std::istream&   read(char* buf, std::streamsize num);
    std::ostream&   write(const char* buf, std::streamsize num);

    // 読み込んだサイズを調べる
    std::streamsize gcount() const;
    ```

* ランダムアクセス
    現在の位置を知る

    ```c++
    std::ios::pos_type tellg() const;    // 入力ストリームの位置
    std::ios::pos_type tellp() const;    // 出力ストリームの位置
    ```

    現在の位置を変更する

    ```c++
    std::istream&   seekg(std::ios::pos_type position);    // 入力ストリームの位置変更
    std::ostream&   seekp(std::ios::pos_type position);    // 出力ストリームの位置変更

    in.seekg(  in.tellg()+std::streamoff{12}  );    // 現在の位置から12バイト後ろに

    std::istream&   seekg(std::ios::off_type offset, std::ios::seekdir origin); // 相対位置
    std::ostream&   seekp(std::ios::off_type offset, std::ios::seekdir origin); // 相対位置
    ```

    originに指定できる値：
    1. `std::ios::beg`：ファイルの先頭
    2. `std::ios::cur`：現在の位置
    3. `std::ios::end`：ファイルの最後

* 入出力状態のチェック
    エラーなどの状態に関する情報を`std::ios::iostate`型という値で保存している
    | `std::ios::iostate`の値 |                    意味                    |
    | :---------------------: | :----------------------------------------: |
    |    std::ios::goodbit    |            エラーが起きなかった            |
    |    std::ios::eofbit     |          ファイルの終端に到達した          |
    |    std::ios::failbit    |        致命的ではないエラーが起きた        |
    |    std::ios::badbit     | (復旧できないような)致命的なエラーが起きた |
    現在の情報の取得

    ```c++
    // 全ての情報を取得できるため、ビット演算で結果を分析する必要がある
    std::ios::iostate   rdstate() const;
    // 特定な状態を調べる
    bool good() const;
    bool eof() const;
    bool fail() const;
    bool bad() const;
    ```

    エラーフラグは通常自動的にクリアされない。
    もしエラーを解決できたら、エラーフラグもクリアする必要がある

    ```c++
    void clear(std::ios::iostate flags = std::ios::goodbit);
    // 例：
    std::ifstream in{"aaa.txt"};
    in.clear(std::ios::badbit | std::ios::failbit);    // フラグのセット
    in.clear();    // フラグのリセット
    ```
