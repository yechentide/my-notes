# Dataの拡張

## 概要

Dataの拡張

## コード例

```swift
import Foundation

extension Data {
    var hexString: String {
        var result: [String] = []
        for byte in self {
            let hexString = String(format: "%02X", byte)
            result.append(hexString)
        }
        return "0x" + result.joined(separator: "_")
    }
    
    var uint8: UInt8 {
        UInt8(littleEndian: self[ self.startIndex...self.startIndex ].withUnsafeBytes{
            $0.load(as: UInt8.self)
        })
    }
    
    var int8: Int8 {
        Int8(littleEndian: self[ self.startIndex...self.startIndex ].withUnsafeBytes{
            $0.load(as: Int8.self)
        })
    }
    
    var int32: Int32 {
        Int32(littleEndian: self[ self.startIndex+0...self.startIndex+3 ].withUnsafeBytes{
            $0.load(as: Int32.self)
        })
    }
}
```
