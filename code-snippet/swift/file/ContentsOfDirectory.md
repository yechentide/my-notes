# ディレクトリ内のファイル

## 概要

ディレクトリ内のファイルを調べる。  
そのファイル関連の情報も出力する。

## コード例

```swift
import Foundation

let directoryUrl = URL(string: "/tmp/work/study_world")!
let keys : [URLResourceKey] = [.nameKey, .isDirectoryKey, .fileSizeKey]

do {
    let urls = try FileManager.default.contentsOfDirectory(at: directoryUrl, includingPropertiesForKeys: keys)

    for url in urls {
        let attributes = try url.resourceValues(forKeys: Set(keys))
        print(attributes.name ?? "nil")
        print(attributes.isDirectory ?? "nil")
        print(attributes.fileSize ?? "nil")
        print()
    }
} catch let error {
    print(error)
}
```
