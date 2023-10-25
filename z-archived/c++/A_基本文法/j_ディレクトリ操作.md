# ディレクトリ操作

```c++
#include <iostream>
#include <filesystem>
using namespace std;

// https://nompor.com/2019/02/16/post-5089/
// https://cpprefjp.github.io/reference/filesystem.html
int main() {
    using Path = std::filesystem::path;

    cout << "\n\n==========test==========\n" << endl;

    // アプリケーションの実行階層を取得
    Path path = filesystem::current_path();
    cout << "実行ファイルの場所：　" << path.string() << endl;

    // 相対パスを絶対パスに変換
    Path result = filesystem::absolute("test.txt");
    cout << "test.txtを絶対パスに変換：　" << result.string() << endl;

    // ファイルの存在チェック
    string filename = "test.txt";
    bool isExists = filesystem::exists(filename);
    cout << filename << "は存在する？　" << (isExists ? "true" : "false") << endl;

    // ファイル種類のチェック
    bool isDirectory = filesystem::is_directory(filename);
    cout << filename << "はディレクトリ？　" << (isDirectory ? "true" : "false") << endl;
    bool isFile = filesystem::is_regular_file(filename);
    cout << filename << "はファイル？　" << (isFile ? "true" : "false") << endl;

    // ファイル、ディレクトリなどのコピー
    // filesystem::copy("test.txt", "test2.txt");      // ディレクトリの内容までコピーする場合、第３引数std::filesystem::copy_options::recursiveを指定する

    // ファイル、ディレクトリなどの移動＆名前変更
    // filesystem::rename("test2.txt", "test3.txt");   // パスを変えれば、移動できる

    // ファイル、ディレクトリなどの削除
    // filesystem::remove("./test.txt");      //ファイル削除
    // filesystem::remove_all("./test");      //ファイル削除（全階層）

    // ディレクトリの作成
    filesystem::create_directory("./testtest");
    filesystem::create_directories("./testtest/inner/inin");

    // ディレクトリの中身を取得
    // フォルダの中身を取得するにはdirectory_iteratorを使用
    // 全階層取得する場合はrecursive_directory_iteratorを使用
    cout << "\n\nディレクトリの中身：　" << endl;
    filesystem::directory_iterator e = filesystem::directory_iterator("./testtest");
    for (auto f : e) {
        cout << f.path().filename() << endl;
    }


    cout << endl;
    return 0;
}
```
