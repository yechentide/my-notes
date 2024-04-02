# 数値型の拡張

## 概要

数値型の拡張

## コード例

```swift
import Foundation

extension FixedWidthInteger {
    var data: Data {
        return withUnsafeBytes(of: self) { Data($0) }
    }
    
    var binaryString: String {
        var result: [String] = []
        for i in 0..<(Self.bitWidth / 8) {
            let byte = UInt8(truncatingIfNeeded: self >> (i * 8))
            let byteString = String(byte, radix: 2)
            let padding = String(repeating: "0", count: 8 - byteString.count)
            result.append(padding + byteString)
        }
        return "0b" + result.joined(separator: "_")
    }
    
    mutating func bitOn(offset: UInt8) {
        guard (0..<Self.bitWidth).contains(Int(offset)) else { return }
        let newValue = (self >> offset | 0x1) << offset | self
        self = newValue
    }
    
    mutating func bitOff(offset: UInt8) {
        guard (0..<Self.bitWidth).contains(Int(offset)) else { return }
        let left  = (self >> (offset + 1)) << (offset + 1)
        let n = UInt8(MemoryLayout.size(ofValue: self) * 8)
        let right = (self << (n-offset)) >> (n-offset)
        self = left | right
    }
    
    func isBitOn(offset: UInt8) -> Bool {
        return (self >> offset) & 0x1 == 1
    }
    
    var bitArray: [UInt8] {
        // 0b0010 ---> [0b0, 0b1, 0b0, 0b0]
        var array = [UInt8]()
        for offset in 0..<self.bitWidth {
            let flag = (self >> offset) & 0x1 == 1
            array.append(flag ? 1 : 0)
        }
        return array
    }
}
```
