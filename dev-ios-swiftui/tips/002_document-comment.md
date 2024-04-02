# ドキュメントコメント

## Document Items

以下の３つは`///`で表記すると機能しなくなる
- `// MARK: - hoge -`
- `// TODO: - hoge -`
- `// FIXME: - hoge -`

## Document comment

`///`と`/**   */`を使う。  
`///`の場合、コメントが複数行でも、一行になって表示されるし、Markdown形式には対応していない。  
`/**   */`の場合、2行開けるとここからはDiscussionになる。  
使えるMarkdown記述:
- `*italic*`
- `**bold**`
- `*** (平行線)`
- `# paragraph`
- `[リンク](https://google.com)`
- リスト: `1.` `-`
- コードブロック: ` ``` `

### 関数コメント

関数にカーソルを合わせた状態で `option + ⌘ + /`で関数のコメントを追加できる
```swift
/// <#Description#>
/// - Parameters:
///   - application: <#application description#>
///   - launchOptions: <#launchOptions description#>
/// - Returns: <#description#>
func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {
    // ...
}
```

### 他のキーワード

`Parameters`や`Returns`以外にも、たくさんのキーワードがある。
- Precondition: 前提条件
- Postcondition: 事後条件
- Requires: 必要条件
- Invariant: 不変
- Complexity: 複雑度
- Important: 重要
- Warning: 警告
- Author: 著者
- Authors: 著者複数
- Copyright: コピーらいと
- Date: 日付
- SeeAlso: これもみてね
- Since: いつから
- Version: バージョン
- Attention: 注意
- Bug: バグ
- Experiment: 実験
- Note: メモ
- Remark: 気付き
- ToDo: TODO

## 参考サイト

[【Xcode】ドキュメントコメントについて知ってる事全てまとめてみた](https://qiita.com/ashdik/items/0019b6ba5ed228d41e66)
