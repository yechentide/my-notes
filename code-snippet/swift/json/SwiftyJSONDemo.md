# SwiftJSONライブラリ

## 概要

`SwiftJSON`ライブラリの簡単な例。

## コード例

```swift
import SwiftyJSON

let token = "????????"

func httpRequestDemo01() {
    var urlComponents = URLComponents(string: "https://api.twitter.com/1.1/search/tweets.json")!
    urlComponents.queryItems = [
        URLQueryItem(name: "lang", value: "ja"),
        URLQueryItem(name: "include_rts", value: "false"),
        URLQueryItem(name: "count", value: "1"),
        URLQueryItem(name: "q", value: "コロナ")
    ]
    print("Request URL:\n\(urlComponents.url!)\n\n")

    var request = URLRequest(url: urlComponents.url!)
    request.addValue("application/json", forHTTPHeaderField: "Content-Type")
    request.addValue("application/json", forHTTPHeaderField: "Accept")
    request.setValue( "Bearer \(token)", forHTTPHeaderField: "Authorization")

    URLSession.shared.dataTask(with: request) { (data, response, error) in  //非同期で通信を行う
        if let error = error {
            print("Failed to get item info: \(error)")
            return;
        }

        if let response = response as? HTTPURLResponse {
            if !(200...299).contains(response.statusCode) {
                print("Response status code does not indicate success: \(response.statusCode)")
                return
            }
        }
        guard let data = data else {
            print("data is empty")
            return
        }
        do {
            let json = try JSON(data: data)
            print(json["statuses"][0]["text"])
        } catch let error {
            print(error)
        }
        print("\n\n")
        exit(0)
    }.resume()
    //task.resume()
}

httpRequestDemo01()
CFRunLoopRun()
```
