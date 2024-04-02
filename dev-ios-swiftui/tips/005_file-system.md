# iOSのファイルシステム

## 概要

iOSアプリごとにSandboxがシステムに用意されている。  
基本的にアプリがアクセスできる領域は自分のSandbox内のみ。  
他のアプリからファイルを受け取ったり、他のアプリのSandbox内のファイルにアクセスするには、Appleが用意した機能を使う必要がある。  

Sandbox内は、大きく３つの部分がある。  
`Bundle Container`と`Data Container`を合わせてホームディレクトリという。
1. `Bundle Container`  
    Xcodeプロジェクト内のリソースがここに保存されている。  
    メインバンドルと呼ばれている。
2. `Data Container`  
    ユーザがPCからコピーしたファイルやキャッシュファイル、設定ファイルなどを保存する。
3. `iCloud Container`  
    iCloudとデータのやり取りを行う場所。

また、データを保存する手段として、４つほどある。  
1. `UserDefaults`
2. `ファイルとして保存`  
3. `キーチェーン`
4. `データベース`

## Data Container

`Data Container`には、７つのディレクトリが最初から用意されている。
1. `Documents/`  
    ユーザが作成した文書やその他のデータ、  
    アプリで再生成できないようなデータはここに保存する。  
    FinderやiTunesで共有したファイルはこの場所に保存される。
2. `Documents/Inbox/`  
    他のアプリから送信されたファイルの保存場所。  
    OpenInという機能で使用される。
3. `Library/`  
    ユーザのデータファイルを保存する場所。  
    Application SupportとCaches以外のサブフォルダを生成することもできる。  
    ユーザに見せたくない(バックアップしたくない)ファイルがある場合は、  
    サブフォルダを生成してバックアップ対象外に設定することもできる。
4. `Library/Application Support/`  
    アプリで生成されたデータファイル、設定ファイルなどを保存する場所。
5. `Library/Caches/`  
    バックアップの対象にならない  
    再ダウンロードや再生成可能なデータはここに保存する。  
    使用しなくなったファイルはアプリが削除する責任がある。  
    一般的にインターネットやデータベースなどのキャッシュファイルに使用される。  
    非常に稀だが、ストレージ容量が少ない場合システムがファイル削除する可能性がある。
6. `Library/Preferences/`  
    UserDefaultsのデータがplistファイルとして保存される場所。
7. `tmp/`  
    バックアップの対象にならない  
    Library/Caches/と同様一時ファイルを保存する場所。  
    システムがファイルを定期的に削除する。

### パス

`NSHomeDirectory()`：それぞれのアプリが利用できるホームディレクトリを取得
```swift
// 自分でパスの文字列を作る
NSHomeDirectory()+"/Documents/test.txt"

// FileManagerで取得する
FileManager.default.urls(for: .documentDirectory, in: .userDomainMask)[0]           // URL型, Documents
FileManager.default.urls(for: .libraryDirectory , in: .userDomainMask)[0]           // URL型, Library
FileManager.default.urls(for: .cachesDirectory  , in: .userDomainMask)[0]           // URL型, Library/Caches
FileManager.default.urls(for: .applicationSupportDirectory, in: .userDomainMask)[0] // URL型, Library/Application Support/

NSTemporaryDirectory()                  // String型, tmp/
FileManager.default.temporaryDirectoy   // String型, tmp/
```

## iCloud Container

iCloudでは3つのストレージタイプがあり、用途に応じて使い分ける。
1. Key-value storage  
   アプリの設定や状態などを保存するためのKVS。最大容量は1アプリ1MB。
2. iCloud Documents  
   ファイルベースのストレージ(Core Data含む)。  
   ドキュメントやドローファイル、または複雑なアプリの状態などを格納するために利用可能。  
   最大保存容量はユーザーのiCloudアカウントに依存。
3. CloudKit storage  
   Appleが提供するBaaSの内のストレージ機能。  
   構造化されたデータを管理したり、ユーザー同士で共有したいファイルやデータを保存したい際に利用する。  
   データの管理にはデータベースが使用され、各レコードはKVSとして扱う。

## UserDefaults

単純な値を格納するために使われる。  
最大1MBだが, 512KBを超えないことをAppleが推奨しているらしい。  
暗号化されていないため, 機密データは`UserDefaults`で保存してはいけない。  
`UserDefaults`クラスの`standard`プロパティで、辞書のようにデータを管理している。  
また、[`Property Wrappers`](https://github.com/apple/swift-evolution/blob/master/proposals/0258-property-wrappers.md#user-defaults)も用意されている。

### データの保存

```swift
// let defaults = UserDefaults.standard
let defaults = UserDefaults(suiteName: "com.test.myapp")
defaults.set(データ, forKey: "キー")
```

### データの読み込み

読み込むデータに応じて、メソッドを利用する
```swift
func bool(forKey: String) -> Bool
func integer(forKey: String) -> Int
func float(forKey: String) -> Float
func double(forKey: String) -> Double
func string(forKey: String) -> String?
func stringArray(forKey: String) -> [String]?
func array(forKey: String) -> [Any]?
func dictionaryRepresentation() -> [String : Any]
func dictionary(forKey: String) -> [String : Any]?
func url(forKey: String) -> URL?
func data(forKey: String) -> Data?
func object(forKey: String) -> Any?
```

### データの削除

```swift
func removeObject(forKey: String)
```

### Property Wrapperを利用する

Swift5.0までは、計算型プロパティの`get`や`set`内でコードを記述する。  
Swift5.1からは、`Property Wrapper`を利用して、少し簡単にできる。  
```swift
@propertyWrapper
struct UserDefault<T> {
    private let key: String
    private let defaultValue: T
    private let userDefaults: UserDefaults

    init(_ key: String, defaultValue: T, userDefaults: UserDefaults = .standard) {
        self.key = key
        self.defaultValue = defaultValue
        self.userDefaults = userDefaults
    }

    var wrappedValue: T {
        get {
            userDefaults.object(forKey: key) as? T ?? defaultValue
        }
        set {
            userDefaults.set(newValue, forKey: key)
        }
    }
}

// 利用例
struct AppData {
    /// 初回起動フラグ
    @UserDefault(key: "FirstLaunchFlag", defaultValue: true)
    static var firstLaunchFlag: Bool

    /// Bool や Int などの基本型以外も格納できる（Codable に準拠している場合のみ）
    @UserDefault<Hoge?>(key: "hoge", defaultValue: nil)
    static var hoge: Hoge?
}
```

## ファイルとして保存

通常、大量のデータはディスクに保存する。  
ユーザが作成したドキュメントも、可能であればディスクに保存する。  
データ構造を`Codable`に適合すると、簡単に`Data`型に変換でき、保存がやりやすくなる。

## キーチェーン

キーチェーンは、ユーザーの機密データを保存することを目的としている。  
APIがC言語ベースで、扱うのがめんどくさいので、ライブラリの利用がおすすめ。

## データベース

データはDocumentsディレクトリに保存されることが多い
- `SQLite`
- `CoreData`
- `Realm`

## 参考サイト

- [iOS Data Storage Guidelines](https://developer.apple.com/icloud/documentation/data-storage/index.html)
- [File System Basics](https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileSystemOverview/FileSystemOverview.html)
- [iOSアプリケーションのフォルダ構成](https://www.techpit.jp/courses/80/curriculums/83/sections/627/parts/2182)
- [iOSでデータを保存する場所を決定する](https://ichi.pro/ios-de-de-ta-o-hozonsuru-basho-o-ketteisuru-69758420013256)
- [How to save data in iOS](https://gist.github.com/y-takagi/9f2cea659fb3f55b56aa04530bf0af39)
- [【iOS】デバイス(ローカル)にデータを保存する方法](https://qiita.com/shiz/items/c7a9b3218269c5c92fed)
- [iOSのファイル保存場所について](https://wingdoor.co.jp/blog/ios/)
- [【Swift】Property Wrapper の基礎から実装まで](https://kobatech-blog.com/swift-property-wrapper/)
