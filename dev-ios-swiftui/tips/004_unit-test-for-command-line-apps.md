# コマンドラインアプリにもユニットテストを

## 概要

コマンドラインアプリのプロジェクトでは、`main.swift`が存在するが、  
その中に記述したクラス・関数などに対してユニットテストができない。  
ユニットテストをするために、フレームワーク用のターゲットを新しく生成し、  
そのターゲットのユニットテストを用意する形になる。

## やり方

1. プロジェクト設定画面で、新しいフレームワークを作る
2. その際、`Include Tests`にチェックを入れる
3. テストが必要なコードはフレームワークに書く(`public`にすること)
4. `main.swift`の方ではフレームワークを`import`して使う
5. ユニットテストを書く

## 参考サイト

[What is a good or common way to structure a CLI appliction in Swift](https://stackoverflow.com/questions/67613494/what-is-a-good-or-common-way-to-structure-a-cli-appliction-in-swift)
