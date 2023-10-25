# UIAlertController

## 中心にあるアラート

```swift
import UIKit

class ViewController: UIViewController, UITextFieldDelegate {

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
    }

    @IBAction func showAlert(_ sender: Any) {
        // アラートを作る
        let alert = UIAlertController(title: nil, message: nil, preferredStyle: .alert)
        alert.title = "タイトル"
        alert.message = "メッセージ文"

        // テキストフィールド
        alert.addTextField(configurationHandler: {(textField) -> Void  in
            textField.delegate = self   // テキストフィールドのデリゲートになる
            // 入力された文字を非表示モードにする.(passwd形式)
            // myTextField.isSecureTextEntry = true
        })

        // OKボタン
        alert.addAction(
            UIAlertAction(
                title: "OK",
                style: .default,
                handler: {(action) -> Void in
                    print(action.title!)
            })
        )

        // キャンセル
        alert.addAction(
            UIAlertAction(
                title: "キャンセル",
                style: .cancel,
                handler: nil)
        )

        // アラートを表示する
        self.present(
            alert,
            animated: true,
            completion: {
                print("アラートが表示された")       // 表示完了後に実行
            }
        )
    }

    // テキストフィールドの編集終了（キーボードを下げる）
    func textFieldDidEndEditing(_ textField: UITextField) {
        print(textField.text ?? "")
    }

}
```

## 下にあるアラート

```swift
import UIKit

class ViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
    }

    @IBAction func showActionSheet(_ sender: Any) {

        // アクションシートを作る
        let actionSheet = UIAlertController(
            title: "タイトル",
            message: "メッセージ文",
            preferredStyle: .actionSheet
        )

        // ボタン1
        actionSheet.addAction(
            UIAlertAction(
                title: "ボタン１です",
                style: .default,
                handler: {(action) -> Void in
                    print(action.title!)
            })
        )

        // ボタン２
        actionSheet.addAction(
            UIAlertAction(
                title: "ボタン２です",
                style: .default,
                handler: {(action) -> Void  in
                    print(action.title!)
            })
        )

        // キャンセル（追加順にかかわらず最後に表示される）
        actionSheet.addAction(
            UIAlertAction(
                title: "キャンセル",
                style: .cancel,
                handler: nil)
        )


        // 赤色のボタン
        actionSheet.addAction(
            UIAlertAction(
                title: "削除します",
                style: .destructive,
                handler: {(action) -> Void in
                    print(action.title!)
            })
        )

        // アクションシートを表示する
        self.present(
            actionSheet,
            animated: true,
            completion: {
                // 表示完了後に実行
                print("アクションシートが表示された")
        }
        )


    }

}
```
