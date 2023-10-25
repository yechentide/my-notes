# APIを叩く

## 概要

RESTの簡単な利用例

## コード例

```swift
import Foundation
import SwiftyJSON

let url = URL(string: "https://itunes.apple.com/search?term=Pixelmator+Pro&country=jp&media=software&entity=macSoftware&limit=1&lang=ja_jp")!

Task {
    let (jsonData, _) = try await URLSession.shared.data(from: url)
    let keywords = ["trackName", "formattedPrice", "version", "currentVersionReleaseDate"]
    
    do {
        let json = try JSON(data: jsonData)
        
        for keyword in keywords {
            if let value = json["results"][0][keyword].string {
                print("\(keyword): \(value)")
            } else {
                print("\(keyword) not found")
            }
        }
    }
    
    exit(0)
}

CFRunLoopRun()
```
