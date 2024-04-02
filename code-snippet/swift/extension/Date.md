# Dateの拡張

## 概要

Dateの拡張

## コード例

```swift
import Foundation

extension Date {
    static func generateCurrentTimeString() -> String {
        let formatter = DateFormatter()
        formatter.locale = Locale.current
        formatter.dateFormat = "yyMMdd-HHmmss"
        
        return formatter.string(from: Date())
    }
}
```
