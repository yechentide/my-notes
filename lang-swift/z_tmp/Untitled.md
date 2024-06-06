# いろいろ

## Decode (Dataから別の型に変換)

### Data  -->  FixedWidthInteger

```swift
self.init(bigEndian: 
    data[data.startIndex..<(data.startIndex + length)]
    .advanced(by: 0)
    .withUnsafeBytes {
        $0.load(as: Self.self)
    })
self.init(littleEndian: nbtData.withUnsafeBytes{
    $0.load(as: Self.self)
})
```

### Data  -->  Float

```swift
self.init(bitPattern: 
    data[data.startIndex..<(data.startIndex + length)]
    .advanced(by: 0)
    .withUnsafeBytes {
        UInt32(bigEndian: $0.load(as: UInt32.self))         // Doubleの場合はUInt64.self
    })
self.init(bitPattern: nbtData.withUnsafeBytes{
    $0.load(as: UInt32.self)                                // Doubleの場合はUInt64.self
})
```

### Data  -->  String

```swift
self.init(data: nbtData, encoding: .utf8)!
```

## Encode (Dataに変換)

### FixedWidthInteger  -->  Data

```swift
withUnsafeBytes(of: bigEndian) { Data($0) }
```

### Float  -->  Data

```swift
withUnsafeBytes(of: bitPattern.bigEndian) { Data($0) }
```

## プロトコル

### 整数タグ

- FixedWidthInteger

### 文字列タグ

- RawRepresentable
- Hashable

### リストタグ

- RandomAccessCollection
- RangeReplaceableCollection

## プログラムから画像生成

### CGImage

```swift
func makeCGImage(color: CGColor) -> CGImage? {
    let size = CGSize(width: 1, height: 1)
    guard let cgContext = CGContext(
        data: nil,
        width: Int(size.width),
        height: Int(size.height),
        bitsPerComponent: 8,
        bytesPerRow: 4 * Int(size.width),
        space: CGColorSpaceCreateDeviceRGB(),
        bitmapInfo: CGImageAlphaInfo.premultipliedLast.rawValue
    ) else {
        return nil
    }
    cgContext.setFillColor(color)
    cgContext.fill(CGRect(origin: .zero, size: size))
    // cgContext.stroke(CGRect(origin: .zero, size: size))
    return cgContext.makeImage()
}
```

### NSImage

```swift
let nsImage = NSImage(
    cgImage: cgImage, 
    size: NSSize(width: cgImage.width, height: cgImage.height)
)
```

## 画像をローカルファイルとして保存

### CGImage

```swift
@discardableResult func save(cgImage: CGImage, to path: URL) -> Bool {
    guard let destination = CGImageDestinationCreateWithURL(path as CFURL, kUTTypePNG, 1, nil) else { return false }
    CGImageDestinationAddImage(destination, cgImage, nil)
    return CGImageDestinationFinalize(destination)
}
```

### NSImage

```swift
@discardableResult func save(nsImage: NSImage, to path: URL) -> Bool {
    let imageRep = NSBitmapImageRep(data: nsImage.tiffRepresentation!)
    let pngData = imageRep?.representation(using: .png, properties: [:])
    do {
        try pngData?.write(to: path)
        return true
    } catch {
        print(error)
        return false
    }
}
```
