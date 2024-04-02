# 非同期処理 (Asynchronous)

## 従来のやり方

### GCD

```swift
import Dispatch
import Foundation

// MARK: 既存のdispatch queueの取得
//let queue = DispatchQueue.main
// QoS: 実行優先度。５つのレベルがある
let queue = DispatchQueue.global(qos: .userInitiated)

// MARK: 新規のdispatch queueの生成
//let queue = DispatchQueue(label: "myTestQueue", qos: .default, attributes: [.concurrent], autoreleaseFrequency: .inherit, target: nil)
    
// MARK: dispatch queueへのタスクの追加
queue.async {
    Thread.sleep(forTimeInterval: 2.0)
    if Thread.isMainThread {
        print("メインスレッド")
    } else {
        print("not メインスレッド")
    }
}
print("hhh")
```

#### クロージャでタスクを渡す

```swift
func download(from url: URL, completion: @escaping (Data) -> Void) { ... }

// 呼び出し
download(from: URL(".....")!) { data in
	// 処理
}
```

### Operation, OperationQueueクラス

```swift
import Foundation

/// - Operation: タスクとその情報
/// - OperationQueue: タスクを実行するキュー
class MyTask: Operation {
    let number: Int
    init(number: Int) { self.number = number }

    override func main() {
        Thread.sleep(forTimeInterval: 1.0)
        guard !isCancelled else {return}    // タスクがキャンセルされる時の対応
        print(number)
    }
}

let queue = OperationQueue()
queue.name = "my queue"
queue.maxConcurrentOperationCount = 2   // 並行に実行するタスクの最大数
queue.qualityOfService = .userInitiated

var tasks = [MyTask]()
for i in 0..<10 {
    tasks.append(MyTask(number: i))
    // tasks[i].addDependency(task[j])　を使うことで、task[j]が終わってからtask[i]が実行される
}

queue.addOperations(tasks, waitUntilFinished: false)
print("タスクが追加された")
tasks[6].cancel()       // タスクをキャンセル
```

### Threadクラス

```swift
import Foundation

class MyThread: Thread {
    override func main() {
        print("executed.")
    }
}
let thread = MyThread()
thread.start()
```

## 新しいやり方

Swift5.5から、`Swift Concurrency`が導入された。  
これにより`async/await`キーワードを使えるようになる

### async / await

`async`関数の呼び出し元は、同じく`async`関数か、`Task`を使う必要がある

```swift
func request(with urlString: String) async -> String {
	let url = URL(string: urlString)!
	do {
		let (data, urlResponse) = try await URLSession.shared.data(from: url, delegate: nil)
		let httpStatus = urlResponse as! HTTPURLResponse
		print(httpStatus.statusCode)
	} catch {
		print(error)
	}
}

// 呼び出し
func task01() async {
	let result = await request(...)
}
func task02() {
	Task {
		let result = await request(...)
	}
}
```

### Structured concurrency

```swift
// 並行処理の例
async let a = asyncFunc01()
aysnc let b = asyncFunc02()
print(await a + b)
```

### Task

```swift
func connectUser(to server: String) async { }

// SynchronousなコードからAsynchronousなコードを呼び出す
Task {
    await connectUser(to: "primary")
}

// タスクグループ
let userIDs = await withTaskGroup(of: Int.self) { group in
    for server in ["primary", "secondary", "development"] {
        group.addTask {
            return await fetchUserID(from: server)
        }
    }

    var results: [Int] = []
    for await result in group {
        results.append(result)
    }
    return results
}
```

### Actors

クラスと似ているが、異なる非同期関数が同時に同じアクターのインスタンスと安全に対話できることを保証できる
```swift
actor ServerConnection {
    var server: String = "primary"
    private var activeUsers: [Int] = []

    func connect() async -> Int {
        let userID = await fetchUserID(from: server)
        // ... communicate with server ...
        activeUsers.append(userID)
        return userID
    }
}
```

### 参考リンク

[Swift 5.5からの非同期処理について](https://zenn.dev/yimajo/scraps/b8e72a3f4e6e5e)
[Swift Concurrency まとめ](https://zenn.dev/akkyie/articles/swift-concurrency)
[iOSにおけるマルチスレッド](https://github.com/mixi-inc/iOSTraining/blob/master/Swift/pages/day3/1-2_Grand-Central-Dispatch.md)
