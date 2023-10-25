# ネットから画像をダウンロード

## 概要

インターネット上の画像をダウンロードする。

## コード例

```swift
import AppKit
import Foundation

// https://stackoverflow.com/questions/29262624/nsimage-to-nsdata-as-png-swift
// https://qiita.com/IKEH/items/af88048caa59a6055e2b
extension NSBitmapImageRep {
    var png: Data? { representation(using: .png, properties: [:]) }
    var gif: Data? { representation(using: .gif, properties: [:]) }
}
extension Data {
    var bitmap: NSBitmapImageRep? { NSBitmapImageRep(data: self) }
}
extension NSImage {
    var png: Data? { tiffRepresentation?.bitmap?.png }
    var gif : Data? { tiffRepresentation?.bitmap?.gif  }
}

func downloadImage(downloadURL : String) {

    let image = NSImage(contentsOf: URL(string: downloadURL)!)!

    var imgData: Data?
    if downloadURL.hasSuffix(".png") {
        imgData = image.png
    } else if downloadURL.hasSuffix(".gif") {
        imgData = image.gif
    } else {
        print("wrong url: \(downloadURL)")
        return
    }

    if let file = imgData {
        let startIndex = downloadURL.lastIndex(of: "/")!
        let endIndex = downloadURL.endIndex
        let fileName = String(   downloadURL[startIndex..<endIndex]   )
        let savePath = FileManager.default.urls(for: .downloadsDirectory, in: .userDomainMask).first!.appendingPathComponent("ImgDownloader"+fileName)
        do {
            try file.write(to: savePath)
            print("OK --- \(savePath)")
        } catch {
            print(error)
        }
    }
}

let urlOfGoogleLogImage = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
downloadImage(downloadURL: urlOfGoogleLogImage)
```
