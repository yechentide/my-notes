# Swiftでコマンド実行

## 概要

Swiftコードで、システムのコマンドを実行する。

## コード例

```swift
import Foundation

let task = Process()
task.launchPath = "/usr/local/bin/brew"
task.arguments = ["list", "--cask"]
let pipe = Pipe()
task.standardOutput = pipe
task.standardError = pipe
task.launch()
task.waitUntilExit()
let data = pipe.fileHandleForReading.readDataToEndOfFile()
let output: String = NSString(data: data, encoding: String.Encoding.utf8.rawValue)! as String

print(output)
```
