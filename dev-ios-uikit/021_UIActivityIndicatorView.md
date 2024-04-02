# UIActivityIndicatorView

## コードで生成

```swift
import UIKit

class ViewController: UIViewController {

    private var myActivityIndicator: UIActivityIndicatorView!
    private var myButton: UIButton!

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.


        // 背景色を黒に設定する.
        self.view.backgroundColor = UIColor.black

        // インジケータを作成する.
        myActivityIndicator = UIActivityIndicatorView()
        myActivityIndicator.frame = CGRect(x:0, y:0, width:50, height:50)
        myActivityIndicator.center = self.view.center

        // アニメーションが停止している時もインジケータを表示させる.
        myActivityIndicator.hidesWhenStopped = false
        myActivityIndicator.style = .white

        // アニメーションを開始する.
        myActivityIndicator.startAnimating()

        // インジケータをViewに追加する.
        self.view.addSubview(myActivityIndicator)

        // ボタンを生成する.
        myButton = UIButton(frame: CGRect(x:0, y:0, width:60, height:60))
        myButton.backgroundColor = UIColor.red
        myButton.layer.masksToBounds = true
        myButton.layer.cornerRadius = 30.0
        myButton.setTitle("Stop", for: .normal)
        myButton.layer.position = CGPoint(x: self.view.bounds.width/2, y: self.view.bounds.height-50)
        myButton.addTarget(self, action: #selector(ViewController.onClickMyButton(sender:)), for: .touchUpInside)

        // ボタンをViewに追加する.
        self.view.addSubview(myButton)
    }


    /* ボタンイベント. */
    @objc internal func onClickMyButton(sender: UIButton){

        if myActivityIndicator.isAnimating {
            myActivityIndicator.stopAnimating()
            myButton.setTitle("Start", for: .normal)
            myButton.backgroundColor = UIColor.blue
        }
        else {
            myActivityIndicator.startAnimating()
            myButton.setTitle("Stop", for: .normal)
            myButton.backgroundColor = UIColor.red
        }
    }

}

```
