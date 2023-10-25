# 単色の画像ファイル生成

## 概要

単色の`CGImage`の生成、そしてそれを利用してファイルを生成する。

## コード例

```swift
import Foundation

struct GridTile {
    let color: CGColor
    init(red: CGFloat, green: CGFloat, blue: CGFloat, alpha: CGFloat) {
        color = CGColor(red: red, green: green, blue: blue, alpha: alpha)
    }

    func createCGImage(data: [UTF8]) -> CGImage? {
        let image = CGDataProvider(dataInfo: nil, data: data, size: data.count, releaseData: {_,_,_ in }).flatMap {
            CGImage(width: 256, height: 256,
                    bitsPerComponent: 8, bitsPerPixel: 24, bytesPerRow: 3,
                    space: CGColorSpaceCreateDeviceRGB(),
                    bitmapInfo: [],
                    provider: $0,
                    decode: nil,
                    shouldInterpolate: false,
                    intent: .defaultIntent)
        }
        return image
    }

    @discardableResult func writeCGImage(_ image: CGImage, to destinationURL: URL) -> Bool {
        guard let destination = CGImageDestinationCreateWithURL(destinationURL as CFURL, kUTTypePNG, 1, nil) else { return false }
        CGImageDestinationAddImage(destination, image, nil)
        return CGImageDestinationFinalize(destination)
    }
}
```
