# UITableView

## UITableViewクラス

[UITableView](https://developer.apple.com/documentation/uikit/UITableView)

### 継承関係

```swift
class UITableView : UIScrollView        // : UIView : UIResponder : NSObject
```

### コンストラクタ

```swift
init()
init(frame: CGRect, style: UITableView.Style)
init?(coder: NSCoder)
```

### UITableViewDataSource

[UITableViewDataSource](https://developer.apple.com/documentation/uikit/uitableviewdatasource)

### UITableViewDelegate

[UITableViewDelegate](https://developer.apple.com/documentation/uikit/uitableviewdelegate)

## UITableViewの使い方

### 基本

- [デフォルトTableViewサンプル集](https://qiita.com/am10/items/9bbbe794e88a96e5420e)
- [コードのみでUITableViewを設置する](https://first-code.hatenablog.com/entry/2019/09/22/095301)
```swift
import UIKit

class MainVC: UIViewController, UITableViewDataSource, UITableViewDelegate {
    var table = UITableView()
    var dataList = ["", "", "", "", "", ""]

    override func viewDidLoad() {
        super.viewDidLoad()
        self.title = "素材"
        self.view.backgroundColor = .white

        self.view.addSubview(table)
        table.dataSource = self
        table.delegate = self
        table.allowsSelection = false
        // 内容がない行を隠す
        table.tableFooterView = UIView()
    }

    override func viewWillLayoutSubviews() {
        super.viewWillLayoutSubviews()

        table.translatesAutoresizingMaskIntoConstraints = false
        NSLayoutConstraint.activate([
            table.leadingAnchor.constraint(equalTo: self.view.leadingAnchor, constant: 0),
            table.trailingAnchor.constraint(equalTo: self.view.trailingAnchor, constant: 0),
            table.topAnchor.constraint(equalTo: self.view.topAnchor, constant: 0),
            table.bottomAnchor.constraint(equalTo: self.view.bottomAnchor, constant: 0)
        ])
    }
}

extension MainVC: UITableViewDataSource {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return dataList.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = UITableViewCell(style: .default, reuseIdentifier: "myCell")
        cell.textLabel!.text = dataList[indexPath.row]
        return cell
    }
}
```

### 編集モード

- [UITableViewの編集モードを使ってCellの削除を実装するまで](https://qiita.com/nasutaro211/items/50d1dbc89969d873b7da)
- [TableViewで複数セルを一気に複数削除する](https://qiita.com/vivayashi/items/9cf122ad625867db0ec9)
```swift
import UIKit

class MainVC: UIViewController {

    var table = UITableView()
    var dataList = ["AAA", "BBB", "CCC", "DDD", "EEE", "FFF"]

    override func viewDidLoad() {
        super.viewDidLoad()
        self.title = "Main"
        view.backgroundColor = .white

        self.view.addSubview(table)
        table.dataSource = self
        table.delegate = self
        table.allowsSelection = false
        table.tableFooterView = UIView()        // 内容がない行を隠す
        table.allowsMultipleSelectionDuringEditing = true    // 編集モードで複数選択可能に

        // 編集モードに入るボタン
        self.navigationItem.leftBarButtonItem = editButtonItem
        let addBtn = UIBarButtonItem(barButtonSystemItem: .add, target: self, action: nil)
        self.navigationItem.rightBarButtonItem = addBtn
    }

    override func viewWillLayoutSubviews() {
        super.viewWillLayoutSubviews()

        table.translatesAutoresizingMaskIntoConstraints = false
        NSLayoutConstraint.activate([
            table.leadingAnchor.constraint(equalTo: self.view.leadingAnchor, constant: 0),
            table.trailingAnchor.constraint(equalTo: self.view.trailingAnchor, constant: 0),
            table.topAnchor.constraint(equalTo: self.view.topAnchor, constant: 0),
            table.bottomAnchor.constraint(equalTo: self.view.bottomAnchor, constant: 0)
        ])
    }

    override func setEditing(_ editing: Bool, animated: Bool) {
        super.setEditing(editing, animated: animated)

        if editing {
            let deleteBtn = UIBarButtonItem(title: "Delete", style: .plain, target: self, action: nil)
            self.navigationItem.rightBarButtonItem = deleteBtn
        } else {
            let addBtn = UIBarButtonItem(barButtonSystemItem: .add, target: self, action: nil)
            self.navigationItem.rightBarButtonItem = addBtn
        }

        table.isEditing = editing
    }
}

extension MainVC: UITableViewDataSource {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return dataList.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = UITableViewCell(style: .default, reuseIdentifier: "myCell")
        cell.textLabel!.text = dataList[indexPath.row]
        return cell
    }

    @objc func add(){
        print("add")
    }

    @objc func deleteRows(){
        print("delete")
        guard let selectedIndexPaths = self.table.indexPathsForSelectedRows else {
            return
        }
        // 配列の要素削除で、indexの矛盾を防ぐため、降順にソートする
        let sortedIndexPaths =  selectedIndexPaths.sorted { $0.row > $1.row }
        for indexPathList in sortedIndexPaths {
            dataList.remove(at: indexPathList.row) // 選択肢のindexPathから配列の要素を削除
        }
        // tableViewの行を削除
        table.deleteRows(at: sortedIndexPaths, with: .automatic)
        setEditing(false, animated: true)
    }

    // 複数選択が有効にしない場合、下の関数を追加して、編集モードでの、１つずつ削除する機能を実装する
    // スワイプして削除
    func tableView(_ tableView: UITableView, commit editingStyle: UITableViewCell.EditingStyle, forRowAt indexPath: IndexPath) {
        dataList.remove(at: indexPath.row)
        table.deleteRows(at: [indexPath], with: .automatic)
    }
}

extension MainVC: UITableViewDelegate {
    // なし
}
```

### 並べ替え

[UITableView の編集モード](https://swift-ios.keicode.com/ios/uitableview-editing.php)
```swift
extension MainVC: UITableViewDataSource {
    func tableView(_ tableView: UITableView, moveRowAt sourceIndexPath: IndexPath, to destinationIndexPath: IndexPath) {
        let a = dataList[sourceIndexPath.row]
        dataList.remove(at: sourceIndexPath.row)
        dataList.insert(a, at: destinationIndexPath.row)
    }
}
```

### スワイプして削除

[【Swift5】UITableViewの編集モード、スワイプアクションについてまとめてみた](https://swallow-incubate.com/archives/blog/20200309)
[UITableViewのセルの編集制御を極める](https://blog.mothule.com/ios/uitableview/ios-uitableview-uitableviewcell-edit-mode)
```swift
extension MainVC: UITableViewDataSource {
    func tableView(_ tableView: UITableView, canEditRowAt indexPath: IndexPath) -> Bool {
        return true
    }

    func tableView(_ tableView: UITableView, commit editingStyle: UITableViewCell.EditingStyle, forRowAt indexPath: IndexPath) {
        dataList.remove(at: indexPath.row)
        table.deleteRows(at: [indexPath], with: .automatic)
    }
}
```

### スワイプアクションをカスタマイズ

```swift
// 調べる予定
```

### DataSource分離

```swift
import UIKit

class MainVC: UIViewController {

    let myTable = UITableView()
    let myTableDatasource = MyTableDatasource()
    let myTableDelegate = MyTableDelegate()



    override func viewDidLoad() {
        super.viewDidLoad()
        self.view.backgroundColor = .white

        self.navigationItem.title = "Main"
        self.navigationItem.leftBarButtonItem = editButtonItem
        let addBtn = UIBarButtonItem(barButtonSystemItem: .add, target: self, action: #selector(add))
        self.navigationItem.rightBarButtonItem = addBtn



        self.view.addSubview(myTable)
        myTable.allowsSelection = false
        myTable.tableFooterView = UIView()
        myTable.dataSource = myTableDatasource
        myTable.delegate = myTableDelegate
        myTable.register(MyTableCell.self, forCellReuseIdentifier: "myCell")
        myTable.rowHeight = 80
        myTable.separatorInset = UIEdgeInsets(top: 0, left: 80, bottom: 0, right: 0)

    }


    override func viewWillLayoutSubviews() {
        super.viewWillLayoutSubviews()

        myTable.translatesAutoresizingMaskIntoConstraints = false
        NSLayoutConstraint.activate([
            myTable.leadingAnchor.constraint(equalTo: self.view.leadingAnchor, constant: 0),
            myTable.trailingAnchor.constraint(equalTo: self.view.trailingAnchor, constant: 0),
            myTable.topAnchor.constraint(equalTo: self.view.topAnchor, constant: 0),
            myTable.bottomAnchor.constraint(equalTo: self.view.bottomAnchor, constant: 0)
        ])
    }


    override func setEditing(_ editing: Bool, animated: Bool) {
        super.setEditing(editing, animated: animated)

        if editing {
            let btn = UIBarButtonItem(title: "Mode", style: .plain, target: self, action: #selector(changeSelectMode(_:)))
            self.navigationItem.rightBarButtonItem = btn
        } else {
            myTable.allowsMultipleSelectionDuringEditing = false
            let addBtn = UIBarButtonItem(barButtonSystemItem: .add, target: self, action: #selector(add))
            self.navigationItem.rightBarButtonItem = addBtn
        }
        myTable.isEditing = editing
    }


    @objc func changeSelectMode(_ sender: UIBarButtonItem){
        myTable.allowsMultipleSelectionDuringEditing = true
        myTable.isEditing = false
        myTable.isEditing = true

        sender.title = "delete"
        sender.action = #selector(del)
    }

    @objc func add(){
        myTableDatasource.addRow(myTable)
    }

    @objc func del(){
        guard let _ = myTable.indexPathsForSelectedRows else {
            return
        }
        myTableDatasource.deleteRow(myTable)
        self.setEditing(false, animated: true)
    }

}


class MyTableDatasource: UITableView, UITableViewDataSource {

    var dataList = ["AAA", "BBB", "CCC", "DDD", "EEE", "FFF"]

    func addRow(_ table: UITableView){
        print("add")
        print("table.isEditing = \(table.isEditing)")
    }

    func deleteRow(_ table: UITableView){
        let selectedIndexPaths = table.indexPathsForSelectedRows!
        let sortedIndexPaths =  selectedIndexPaths.sorted { $0.row > $1.row }
        for indexPathList in sortedIndexPaths {
            dataList.remove(at: indexPathList.row)
        }
        table.deleteRows(at: sortedIndexPaths, with: .automatic)
    }



    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return dataList.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        // カスタムセル
        let cell = tableView.dequeueReusableCell(withIdentifier: "myCell", for: indexPath) as! MyTableCell
        cell.myLabel.text = dataList[indexPath.row]
        cell.myImageView.image = UIImage(named: "icon")

        cell.accessoryType = .disclosureIndicator
        return cell
    }

    func tableView(_ tableView: UITableView, commit editingStyle: UITableViewCell.EditingStyle, forRowAt indexPath: IndexPath) {
        dataList.remove(at: indexPath.row)
        tableView.deleteRows(at: [indexPath], with: .automatic)
    }

    func tableView(_ tableView: UITableView, moveRowAt sourceIndexPath: IndexPath, to destinationIndexPath: IndexPath) {
        let a = dataList[sourceIndexPath.row]
        dataList.remove(at: sourceIndexPath.row)
        dataList.insert(a, at: destinationIndexPath.row)
    }
}


class MyTableDelegate: UITableView, UITableViewDelegate {   }
```

### ボタン拡張

```swift
import UIKit

class ViewController: UIViewController, UITableViewDelegate, UITableViewDataSource {

    var myItems: [String] = ["TEST1", "TEST2", "TEST3"]
    var myTableView: UITableView = UITableView()

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.


        // Status Barの高さを取得.
        let barHeight: CGFloat = UIApplication.shared.statusBarFrame.size.height

        // Viewの高さと幅を取得.
        let displayWidth: CGFloat = self.view.frame.width
        let displayHeight: CGFloat = self.view.frame.height

        // TableViewの生成( status barの高さ分ずらして表示 ).
        myTableView.frame = CGRect(x: 0, y: barHeight, width: displayWidth, height: displayHeight - barHeight)

        // Cellの登録.
        myTableView.register(UITableViewCell.self, forCellReuseIdentifier: "MyCell")

        // DataSourceの設定.
        myTableView.dataSource = self

        // Delegateを設定.
        myTableView.delegate = self

        // Viewに追加する.
        self.view.addSubview(myTableView)
    }


    /* Cellが選択された際に呼び出される. */
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        print("Num: \(indexPath.row)")
        print("Value: \(myItems[indexPath.row])")
    }

    /* Cellの総数を返す. */
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        print("numberOfRowsInSection")
        return myItems.count
    }

    /* Editableの状態にする. */
    func tableView(_ tableView: UITableView, canEditRowAt indexPath: IndexPath) -> Bool {
        print("canEditRowAtIndexPath")

        return true
    }

    /* 特定の行のボタン操作を有効にする. */
    func tableView(_ tableView: UITableView, commit editingStyle: UITableViewCell.EditingStyle, forRowAt indexPath: IndexPath) {
        print("commitEdittingStyle:\(editingStyle)")
    }

    /* Cellに値を設定する. */
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        print("cellForRowAtIndexPath")

        let cell: UITableViewCell = tableView.dequeueReusableCell(withIdentifier: "MyCell", for: indexPath as IndexPath)
        //tableView.dequeueReusableCellWithIdentifier("MyCell", forIndexPath: indexPath as IndexPath)

        cell.textLabel?.text = "\(myItems[indexPath.row])"

        return cell
    }

    /* Buttonを拡張する. */
    func tableView(_ tableView: UITableView, editActionsForRowAt indexPath: IndexPath) -> [UITableViewRowAction]? {

        // Shareボタン.
        let myShareButton: UITableViewRowAction = UITableViewRowAction(style: .normal, title: "Share") { (action, index) -> Void in

            tableView.isEditing = false
            print("share")

        }
        myShareButton.backgroundColor = UIColor.blue

        // Archiveボタン.
        let myArchiveButton: UITableViewRowAction = UITableViewRowAction(style: .normal, title: "Archive") { (action, index) -> Void in

            tableView.isEditing = false
            print("archive")

        }
        myArchiveButton.backgroundColor = UIColor.red

        return [myShareButton, myArchiveButton]
    }

}
```
