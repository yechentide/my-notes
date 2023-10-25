# UIKitとの併用

SwiftUIは、UIKitなどのネイティブフレームワークと設計が違うため、そのままでは一緒に利用できない。  
したがって、各ネイティブフレームワーク内で、SwiftUIのビューを動かすためのクラスを、  
SwiftUIの中で、ネイティブフレームワークのビューやビューコントローラを動かすためのビューといった、  
橋渡しになる仕組みを用意されている。

## SwiftUIの中でUIKitを使う

### UIViewをラップするビュー

UIKitのビューは、`UIView`クラスかそのサブクラスである。  
SwiftUIでこれらのビューを利用するために、`UIViewRepresentable`に適合するビューを用意する。  
３つの実装が必要:
1. `UIViewType`タイプエイリアス: ラップするUIKitのビュークラス
2. `makeUIView()`メソッド: ラップするUIKitのビューを作成するメソッド
3. `updateView()`メソッド: バインディング更新時や、ビューの状態更新時などに呼ばれるメソッド
```swift
import UIKit
import SwiftUI

struct MyWrappedView: UIViewRepresentable {
    typealias UIViewType = UIKitのビュー
    func makeUIView(context: Context) -> UIKitのビュー {
        let view = UIKitのビューのインスタンスの生成
        return view
    }
    func updateView(_ uiView: UIKitのビュー, context: Context) {}
}
```

### コーディネーターについて

UIKitで頻繁に登場する設計手法の中で、次のような３パターンがある:
1. デリゲート
2. データソース
3. ターゲット＆アクション

これらのパターンはいずれも専用のプロトコルを適合したクラス＆メソッドを実装し、  
フレームワーク側に渡されるインスタンスを知らなくても、必要なメソッドだけ呼び出して、  
必要な情報やオブジェクトを取得し機能を実現している。  
SwiftUIのビューは`NSObject`を継承してないため、このようなプロトコルを実装できない。  
そこで、`NSObject`クラスを継承しているクラスで実装したコーディネーターを使う。  
UIKitのデリゲートやデータソース、ターゲットなどを、ラップしているビューの中で、UIKit側のプロトコルに合わせて実装。

### デリゲートを実装

```swift
struct MapView: UIViewRepresentable {
    @Binding var message: String
    typealias UIViewType = MKMapView

    // コーディネーターのインスタンスを確保する
    func makeCoordinator() -> Coordinator {
        Coordinator(self)
    }

    // UIKitのビューのインスタンスを確保する
    func makeUIView(context: Context) -> MKMapView {
        let view = MKMapView()
        // viewのデリゲート、データソース、ターゲットとしてコーディネーターを指定する
        // makeCoordinatorメソッドで確保したインスタンスは、context.coordinatorで取得する
        view.delegate = context.coordinator
        return view
    }

    // ビューの状態更新
    func updateUIView(_ uiView: MKMapView, context: Context) {}

    // コーディネーター定義、必要に応じて適合するプロトコルを宣言
    class Coordinator: NSObject, MKMapViewDelegate {
        var parent: MapView

        init(_ parent: MapView) { self.parent = parent }

        // 必要なメソッドを実装する
        func mapViewDidChangeVisibleRegion(_ mapView: MKMapView) {
            // デリゲートの実装
        }
    }
}
```

### ターゲットとアクションを実装

UIKitでは、ボタンをタップした時の処理やタイマーなどで、ターゲットとアクションが使われる。  
SwiftUIでの実装では、ターゲットを`Coordinator`クラスにし、メソッドは`Coordinator`クラスのメソッドを設定。
```swift
import UIKit
import SwiftUI

struct ButtonView: UIViewRepresentable {
    @Binding var showingSheet: Bool
    typealias UIViewType = UIButton

    func makeCoordinator() -> Coordinator {
        Coordinator(self)
    }

    func makeUIView(context: Context) -> UIButton {
        let button = UIButton(type: .system)
        button.setTitle("Show Sheet", for: .normal)
        button.addTarget(context.coordinator, action: #selector(Coordinator.showSheet(_:)), for: .touchUpInside)
        return button
    }

    func updateUIView(_ uiView: UIButton, context: Context) {}

    class Coordinator: NSObject {
        var parent: ButtonView
        init(_ parent: ButtonView) { self.parent = parent }
        @objc func showSheet(_ sender: Any) {
            self.parent.showingSheet = true
        }
    }
}
```

### UIViewへの値変更の反映処理を実装

SwiftUIではバインディングを使って、コントロールで行われた値の変更が、自動的に他のビューに適用される。  
`UIViewRepresentable`を適合したビューでもこの仕組みを利用するために、`updateUIView`メソッドとバインディングを使う。  
バインディングをプロパティにし、そのバインディングを変更すると、`updateUIView`メソッドが呼ばれる。
```swift
struct MapView: UIViewRepresentable {
    @Binding var showsBuilding: Bool
    @Binding var showsCompass: Bool
    @Binding var showsScale: Bool
    @Binding var showsTraffic: Bool
    typealias UIViewType = MKMapView

    func makeCoordinator() -> Coordinator {
        Coordinator(self)
    }

    func makeUIView(context: Context) -> MKMapView {
        let view = MKMapView()
        view.delegate = context.coordinator
        applyOptions(view: view)
        return view
    }

    func updateUIView(_ uiView: MKMapView, context: Context) {
        applyOptions(view: uiView)
    }

    func applyOptions(view: MKMapView) {
        view.showsBuildings = showsBuilding
        view.showsCompass = showsCompass
        view.showsScale = showsScale
        view.showsTraffic = showsTraffic
    }

    class Coordinator: NSObject, MKMapViewDelegate {
        var parent: MapView
        init(_ parent: MapView) { self.parent = parent }
    }
}
```

### UIViewControllerをラップするビュー

UIKitのビューコントローラは、`UIViewController`クラスかそのサブクラスである。  
SwiftUIでこれらのビューを利用するために、`UIViewControllerRepresentable`に適合するビューを用意する。
```swift
struct WrappedView: UIViewControllerRepresentable {
    typealias UIViewControllerType = ラップするビューコントローラのクラス

    func makeCoordinator() -> Coordinator { return Coordinator(self) }
    func makeUIViewController(context: Context) -> some ビューコントローラ {}
    func updateUIViewController(_ uiViewController: UIViewControllerType, context: Context) {}

    class Coordinator: NSObject {
        var parent: WrappedView
        init(_ parent: WrappedView) { self.parent = parent }
    }
}
```

### データソースを実装

```swift
struct TableView: UIViewRepresentable {
    typealias UIViewType = UITableView
    var items: [String]

    func makeCoordinator() -> Coordinator { return Coordinator(self) }

    func makeUIView(context: Context) -> UITableView {
        let tableView = UITableView()
        tableView.dataSource = context.coordinator
        tableView.reloadData()
        return tableView
    }

    func updateUIView(_ uiView: UITableView, context: Context) {}

    class Coordinator: NSObject, UITableViewDataSource {
        var parent: TableView
        init(_ parent: TableView) { self.parent = parent }

        func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
            return self.parent.items.count
        }
        func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
            let cell = UITableViewCell(style: .default, reuseIdentifier: nil)
            cell.textLabel?.text = self.parent.items[indexPath.row]
            return cell
        }
    }
}
```

## UIKitの中でSwiftUIを使う

### SwiftUIのビューを表示する

UIKitの中でSwiftUIのビューを表示する時は、`UIHostingController`を使うだけで実現できる
```swift
struct ImageView: View {
    var body: some View {
        VStack {
            Image("Sample")
            Text("Now Thinking...")
        }
    }
}

class ViewController: UIViewController {
    var imageViewController: UIHostingController<ImageView>!

    override func viewDidLoad() {
        super.viewDidLoad()
        // SwiftUIのビューを作って、全体に広げて表示する
        self.imageViewController = UIHostingController(rootView: ImageView())
        self.imageViewController.view.frame = self.view.bounds
        // サブビューとして追加
        self.addChild(self.imageViewController)
        self.view.addSubview(self.imageViewController.view)
        // 自動レイアウトで追従するように設定する
        self.imageViewController.view.autoresizingMask = [.flexibleWidth, .flexibleHeight]
    }
}
```

### SwiftUIのビュー内の変更に対応する

```swift
// SwiftUI側
import Combine
class MyObject {
    @Pubilshed var value: Int = 0
}

// UIKit側
let obj = MyObject()
let objSink = obj.$value
    .sink() { value in
    // 値変更時に行いたい処理。変更後の値はvalueに入っている。
}
```
