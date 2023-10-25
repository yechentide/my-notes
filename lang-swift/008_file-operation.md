# ファイル出入力

## ファイル操作

### 新規作成

```swift
// ファイル
let fileManager = FileManager.default
fileManager.createFile(atPath: "file path",  contents: Data(), attributes: nil)

// ディレクトリ
fileManager.createDirectory(at: "dir path", withIntermediateDirectories: true, attributes: nil)
```

### 存在チェック

```swift
let fileManager = FileManager.default
if fileManager.fileExists(atPath:"パス") {
    print("存在する")
}
```

### ディレクトリかどうか

ディレクトリが存在しない場合も`false`
```swift
var isDirectory: ObjCBool = false
FileManager.default.fileExists(atPath: "path", isDirectory: &isDirectory)
print(isDirectory.boolValue)
```

### 削除

```swift
FileManager.default.fileManager.removeItem(atPath: "path")
```

### 移動

移動先に同名ファイルが存在する場合はエラー。  
これでファイル名・ディレクトリ名を変更できる。
```swift
FileManager.default.fileManager.moveItem(atPath: "path", toPath: "path")
```

### コピー

コピー先に同名ファイルが存在する場合はエラー
```swift
FileManager.default.fileManager.copyItem(atPath: "path", toPath: "path")
```

### ディレクトリ内のファイル列挙

ディレクトリ内部のコンテンすを再帰的に取得する場合、`fileManager.subpathsOfDirectory`を使う。
```swift
let fileManager = FileManager.default
do {
    guard let contentURLs = try fileManager.contentsOfDirectory(at: "URL", includingPropertiesForKeys: nil) else {
        return
    }

    for url in contentURLs {
        print("contentUrl: \(url)")
        print("contentUrlPath: \(url.path)")
        print("contentUrlLastPathComponent: \(url.lastPathComponent)")
        print("FileName: \(try FileManager.default.contentsOfDirectory(atPath: url.path))")
    }
} catch {
    print("フォルダが空です。")
}
```

### ファイルの情報を取得

```swift
let attributes: [FileAttributeKey : Any] = FileManager.default.attributesOfItem(atPath: "path")
```

### ファイルサイズ取得

```swift
func checkSize(url: URL) -> Int64 {
    //print("target path: \(url.path)")
    var bytes: Int64 = 0
    let keys : [URLResourceKey] = [.nameKey, .fileSizeKey, .isDirectoryKey]
    var isDir: ObjCBool = false
    FileManager.default.fileExists(atPath: url.path, isDirectory: &isDir)
    //print(isDir.boolValue ? "is dir" : "not dir")

    do {
        if isDir.boolValue, let items = FileManager.default.enumerator(atPath: url.path) {
            for item in items {
                let fileURL = url.appendingPathComponent(item as! String)
                let attributes = try! fileURL.resourceValues(forKeys: Set(keys))
                if let b = attributes.fileSize {
                    bytes += Int64(b)
                }
            }
        } else {
            //let attributes: [FileAttributeKey : Any] = try! FileManager.default.attributesOfItem(atPath: filePath)
            //let fileSize = attributes[FileAttributeKey.size] as! NSNumber
            let resources = try url.resourceValues(forKeys: Set(keys))
            if let fileSize = resources.fileSize {
                bytes += Int64(fileSize)
            }
        }
    } catch {
        print(error)
    }

    //let kb = Double(bytes) / 1000.0
    //print("size: \(kb) KB")
    return bytes
}
```

## データ操作

### 書き込み

Stringの`write(toFile: atomically: encoding:)`メソッドを使えば良い  
ファイルの保存は失敗する可能性があるので、例外処理に組み込む  
第２引数のatomicallyをtrueにすると、書き出し中に異常終了してもファイルが壊れないように、  
一時ファイルを使って書き込む処理を行う
```swift
do{
    try str.write(toFile: "パス", atomically: true, encoding: String.Encoding.utf8)
}catch let error as NSError{
    print(error)
}
```

### 読み込み

`String(contentsOfFile: encoding:)`メソッドを使う  
書き込みと同じく、例外処理に組み込む
```swift
do{
    let data = try String(contentsOfFile: "パス", encoding: String.Encoding.utf8)
}catch let error as NSError{
    print(error)
}
```

## 参考サイト

- [いまさらだけどiOSのファイル操作まとめ（Swift）](https://qiita.com/am10/items/3b2eb3d9f6c6955455b6)
- [【Swift】ファイル操作で保存されたファイルやパスを色んな形式で確認してみる](https://qiita.com/ta9yamakawa/items/dbebfcdad34b66b6a98a)
