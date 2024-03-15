# Xcode Tips

## ショートカット

- 見た目
    - `command` + `option` + `←`: コードブロックを折り畳む
    - `command` + `option` + `→`: コードブロックを展開
    - `command` + `shift` + `option` + `←`: 全てのコードブロックを折り畳む
    - `command` + `shift` + `option` + `→`: 全てのコードブロックを展開
    - `command` + `shift` + `;`: フォントを大きくする
    - `command` + `-`: フォントを小さくする
    - `command` + `control` + `0`: フォントの大きさを元に戻す
- フォーマット
    - `control` + `i`: 選択範囲内のインデントを揃える
    - `control` + `m`: 関数パラメータの整形
- その他
    - `command` + `control` + `e`: スコープ内の変数名・定数名変更
    - `command` + `control` + `shift` + `w`: 全タブを閉じる
    - `command` + `option` + `enter`: プレビュー表示の切り替え
    - `command` + `option` + `p`: プレビューのResume

## アノテーションコメント

`:`の次に `-` をつけると、線が引かれる。(`// MARK: - コメント`)
- `// MARK:`
- `// FIXME:`
- `// TODO:`
- `// !!!:`
- `// ???:`

## Docコメント

メソッドなどの説明を書くためのコメント  
カーソルを合わせて、`command` + `option` + `/`で生成できる
```swift
/// Docコメント
func test() {}

/**
Docコメント
*/
func test() {}

/// <#Description#>
/// - Parameter text: <#text description#>
/// - Returns: <#description#>
func test001(text: String) -> String { return "" }
```

## その他

- `option`押しながらドラッグ: マルチカーソル & 矩形選択
- `shift` + `control` + `↓`: マルチカーソル
- シミュレーターの時刻を変える: `xcrun simctl status_bar "iPhone 11 Pro Max" override --time '9:41'`
- パラメータのプレースホルダー: `Text(<#ここにパラメータを入れるよ#>)`

## 参考サイト

- [【知らなきゃ損！】Xcodeのちょっとしたテクニックや裏技、設定を33個紹介！](https://ios-docs.dev/xcode-technic/)
- [【Xcode】ドキュメントコメントについて知ってる事全てまとめてみた](https://qiita.com/ashdik/items/0019b6ba5ed228d41e66#ドキュメントコメントの書き方)
- [【Xcode】ヘッダーコメントのカスタマイズ](https://capibara1969.com/1151/)
- [macOSで空き容量を増やす〜xcode関連の不要なファイル削除](https://101010.fun/programming/delete-garbage-files-macos.html)
