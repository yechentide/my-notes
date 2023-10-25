# UTIを調べる

## 概要

Document-basedアプリではUTIを使うが、公式サイトに説明されていないUTIがたくさんあるので、コードで調べる。

## コード例

```swift
import Foundation

if let str = UTTypeCreatePreferredIdentifierForTag(kUTTagClassFilenameExtension, "txt" as CFString, nil)?.takeUnretainedValue() {
    print("\(str)") // org.openxmlformats.spreadsheetml.sheet
}

if let str = UTTypeCreatePreferredIdentifierForTag(kUTTagClassFilenameExtension, "rar" as CFString, nil)?.takeUnretainedValue() {
    print("\(str)") // org.openxmlformats.spreadsheetml.sheet
}
```
