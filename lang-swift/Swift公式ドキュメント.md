# Swift公式ドキュメント

[公式ドキュメント - Apple](https://developer.apple.com/documentation/swift)
[公式ドキュメント - Swift](https://www.swift.org)

## ClassにするかStructureにするか

[Choosing Between Structures and Classes](https://developer.apple.com/documentation/swift/choosing-between-structures-and-classes)

- まずはstructを使おう
- Protocolをstructに適用して、実装を共有しよう
- 以下のことをしたいなら、classに変える
    - Objective-Cを使いたい
    - データのIDをコントロールしたり、インスタンスを共有したいとき

## Enumで状態を管理する

[Maintaining State in Your Apps](https://developer.apple.com/documentation/swift/maintaining-state-in-your-apps)

```swift
class App {
    enum State {
        case unregistered
        case loggedIn(User)
        case sessionExpired(User)
    }
    var state: State = .unregistered
    // ...
}
```
