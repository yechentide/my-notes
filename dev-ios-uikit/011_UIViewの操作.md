# UIView

## UIViewをスナップさせる(タップした場所に移動)

```swift
import UIKit

class ViewController: UIViewController {

    // UIDynamicAnimatorのインスタンスを保存しなければアニメーションが実行されない.
    var animator : UIDynamicAnimator!
    var myLabel : UILabel!

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.


        // 背景を水色に設定.
        self.view.backgroundColor = UIColor.cyan

        // Labelを作成.
        myLabel = UILabel(frame: CGRect(x: 0, y: 0, width: 200, height: 50))
        myLabel.backgroundColor = UIColor.orange
        myLabel.layer.masksToBounds = true
        myLabel.layer.cornerRadius = 20.0
        myLabel.text = "Hello Swift!!"
        myLabel.textColor = UIColor.white
        myLabel.shadowColor = UIColor.gray
        myLabel.textAlignment = NSTextAlignment.center
        myLabel.layer.position = CGPoint(x: self.view.bounds.width/2,y: 200)
        self.view.backgroundColor = UIColor.cyan
        self.view.addSubview(myLabel)

        // UIDynamiAnimatorの生成とインスタンスの保存.
        animator = UIDynamicAnimator(referenceView: self.view)
    }


    /* タップを感知した時に呼ばれるメソッド. */
    override func touchesBegan(_ touches: Set<UITouch>, with event: UIEvent?) {

        for touch : AnyObject in touches {

            // タッチされた座標を取得.
            let location = touch.location(in: self.view)

            // animatorに登録されていたBahaviorを全て削除.
            animator.removeAllBehaviors()

            // UIViewをスナップさせるUISnapBehaviorを生成.
            let snap = UISnapBehavior(item: myLabel, snapTo: location)

            // スナップを適用させる.
            animator.addBehavior(snap)
        }
    }

}
```

## UIViewを動的に動かす

```swift
import UIKit

class ViewController: UIViewController {

    // UIDynamicAnimatorのインスタンスを保存しなければアニメーションが実行されない.
    var animator : UIDynamicAnimator!
    var continuousPush : UIPushBehavior!
    var instantaneousPush : UIPushBehavior!

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.


        // 背景を水色に設定.
        self.view.backgroundColor = UIColor.cyan

        // ContinuousButtonを作成.
        let myButton = UIButton(frame: CGRect(x: 0, y: 0, width: 200, height: 50))
        myButton.layer.position = CGPoint(x: self.view.center.x, y: self.view.bounds.maxY - myButton.bounds.midY)
        myButton.layer.masksToBounds = true
        myButton.layer.cornerRadius = 20.0
        myButton.setTitleColor(UIColor.white, for: UIControl.State.normal)
        myButton.setTitleColor(UIColor.black, for: UIControl.State.highlighted)
        myButton.backgroundColor = UIColor.red
        myButton.setTitle("Continuous!!", for: UIControl.State.normal)
        myButton.addTarget(self, action: #selector(ViewController.onClickMyButton(sender:)), for: UIControl.Event.touchUpInside)
        myButton.tag = 1
        self.view.addSubview(myButton)

        // InstaneousButtonを作成.
        let myButton2 = UIButton(frame: CGRect(x: 0, y: 0, width: 200, height: 50))
        myButton2.layer.position = CGPoint(x: self.view.center.x, y: self.view.bounds.minY + myButton.bounds.midY)
        myButton2.layer.masksToBounds = true
        myButton2.layer.cornerRadius = 20.0
        myButton2.setTitleColor(UIColor.white, for: UIControl.State.normal)
        myButton2.setTitleColor(UIColor.black, for: UIControl.State.highlighted)
        myButton2.backgroundColor = UIColor.red
        myButton2.setTitle("Instantaneous!!", for: UIControl.State.normal)
        myButton2.addTarget(self, action: #selector(ViewController.onClickMyButton(sender:)), for: UIControl.Event.touchUpInside)
        myButton2.tag = 2
        self.view.addSubview(myButton2)

        // UIDynamiAnimatorの生成とインスタンスの保存.
        animator = UIDynamicAnimator(referenceView: self.view)

        // UIViewを等加速度運動で動かすUIPushBehaviorを生成.
        continuousPush = UIPushBehavior(items: [myButton], mode: UIPushBehavior.Mode.continuous)
        continuousPush.pushDirection = CGVector(dx: 0.0, dy: -1.0)

        // UIViewを等速運動で動かすUIPushBehaviorを生成.
        instantaneousPush = UIPushBehavior(items: [myButton2], mode: UIPushBehavior.Mode.instantaneous)
        instantaneousPush.pushDirection = CGVector(dx: 0.0, dy: 1.0)
    }


    /* Buttonが押されたときに呼ばれるメソッド. */
    @objc func onClickMyButton(sender : UIButton){

        switch(sender.tag){
        case 1:
            animator.addBehavior(continuousPush)

        case 2:
            animator.addBehavior(instantaneousPush)

        default:
            print("Error")
        }
    }

}
```

## UIViewアニメーションまとめ

```swift
import UIKit

class ViewController: UIViewController {

    var myLabel: UILabel!

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.


        // Labelを生成
        myLabel = UILabel(frame: CGRect(x: 0, y: 0, width: 200, height: 50))
        myLabel.backgroundColor = UIColor(red: 0.561, green: 0.737, blue: 0.561, alpha: 1.0)
        myLabel.center = self.view.center
        myLabel.text = "*･゜ﾟ･*:.｡..｡.:*･゜"
        myLabel.textAlignment = NSTextAlignment.center
        myLabel.textColor = UIColor.white

        // SegmentesControllerを生成.
        let mySegcon = UISegmentedControl(items: ["Spring", "Invert", "Rotate", "Scaling", "Move"])
        mySegcon.layer.position = CGPoint(x: self.view.frame.width/2, y: self.view.frame.height - 50)
        mySegcon.tintColor = UIColor.blue
        mySegcon.addTarget(self, action: #selector(ViewController.changedValue(sender:)), for: UIControl.Event.valueChanged)

        // Labelをviewに追加
        self.view.addSubview(myLabel)

        // SegmentedControllerをviewに追加.
        self.view.addSubview(mySegcon)
    }


    /* SegmentedControllerの値が変わった時に呼ばれるメソッド. */
    @objc func changedValue(sender: UISegmentedControl) {

        myLabel.center = self.view.center

        // 各アニメーションの処理.
        switch(sender.selectedSegmentIndex) {
            /* バネのような動きをするアニメーション. */
        case 0:
            // アニメーションの時間を2秒に設定.
            UIView.animate(withDuration: 2.0,

                delay: 0.0, // 遅延時間.

                // バネの弾性力. 小さいほど弾性力は大きくなる.
                usingSpringWithDamping: 0.2,

                // 初速度.
                initialSpringVelocity: 1.5,

                // 一定の速度.
                options: UIView.AnimationOptions.curveLinear,

                animations: { () -> Void in

                    self.myLabel.layer.position = CGPoint(x: self.view.frame.width-50, y: 100)

                    // アニメーション完了時の処理
            }) { (Bool) -> Void in
                self.myLabel.center = self.view.center
            }

            /*
             X, Y方向にそれぞれ反転するアニメーション.
             */
        case 1:
            // アニメーションの時間を3秒に設定
            UIView.animate(withDuration: 3.0,

                                       // アニメーション中の処理
                animations: { () -> Void in

                    // X方向に反転用のアフィン行列作成
                    self.myLabel.transform = self.myLabel.transform.scaledBy(x: -1.0, y: 1.0)

                    // 連続したアニメーション処理.
            }) { (Bool) -> Void in
                UIView.animate(withDuration: 3.0,

                                           // アニメーション中の処理
                    animations: { () -> Void in

                        // Y方向に反転用のアフィン行列作成
                        self.myLabel.transform = self.myLabel.transform.scaledBy(x: 1.0, y: -1.0)

                        // アニメーション完了時の処理
                }) { (Bool) -> Void in
                }
            }

            /* 回転アニメーション. */
        case 2:
            // 初期化.
            self.myLabel.transform = CGAffineTransform(rotationAngle: 0)

            // radianで回転角度を指定(90度).
            let angle:CGFloat = CGFloat(Double.pi/2)

            // アニメーションの秒数を設定(3秒).
            UIView.animate(withDuration: 3.0,

                                       animations: { () -> Void in

                                        // 回転用のアフィン行列を生成.
                                        self.myLabel.transform = CGAffineTransform(rotationAngle: angle)
                },
                                       completion: { (Bool) -> Void in
            })

            /* 拡縮アニメーション. */
        case 3:
            self.myLabel.transform = CGAffineTransform(scaleX: 1, y: 1)

            // アニメーションの時間を3秒に設定.
            UIView.animate(withDuration: 3.0,

                                       animations: { () -> Void in
                                        // 縮小用アフィン行列を作成.
                                        self.myLabel.transform = CGAffineTransform(scaleX: 1.5, y: 1.5)
                }) // 連続したアニメーション処理.
            { (Bool) -> Void in
                UIView.animate(withDuration: 3.0,
                                           // アニメーション中の処理.
                    animations: { () -> Void in
                        // 拡大用アフィン行列を作成.
                        self.myLabel.transform = CGAffineTransform(scaleX: 0.5, y: 0.5)
                    }) // アニメーション完了時の処理.
                { (Bool) -> Void in
                    // 大きさを元に戻す.
                    self.myLabel.transform = CGAffineTransform(scaleX: 1, y: 1)
                }
            }

            /* 移動するアニメーション. */
        case 4:
            myLabel.layer.position = CGPoint(x: -30, y: -30)

            // アニメーション処理
            UIView.animate(withDuration: TimeInterval(CGFloat(3.0)),
                                       animations: {() -> Void in

                                        // 移動先の座標を指定する.
                                        self.myLabel.center = CGPoint(x: self.view.frame.width/2,y: self.view.frame.height/2);

                }, completion: {(Bool) -> Void in
            })

        default:
            print("error!")
        }
    }

}
```

## ドラッグでViewを移動させる

UIViewだけじゃない
```swift
import UIKit

class ViewController: UIViewController {

    var myLabel: UILabel!

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.


        // 背景を黒色に設定.
        self.view.backgroundColor = UIColor.black

        // Labelを生成.
        myLabel = UILabel(frame: CGRect(x: 0, y: 0, width: 80, height: 80))
        myLabel.text = "Drag!"
        myLabel.textAlignment = NSTextAlignment.center
        myLabel.backgroundColor = UIColor.red
        myLabel.layer.masksToBounds = true
        myLabel.center = self.view.center
        myLabel.layer.cornerRadius = 40.0

        // Labelをviewに追加.
        self.view.addSubview(myLabel)
    }


    /* タッチを感知した際に呼ばれるメソッド. */
    override func touchesBegan(_ touches: Set<UITouch>, with event: UIEvent?) {
        print("touchesBegan")

        // Labelアニメーション.
        UIView.animate(withDuration: 0.06,
                                   // アニメーション中の処理.
            animations: { () -> Void in
                // 縮小用アフィン行列を作成する.
                self.myLabel.transform = CGAffineTransform(scaleX: 0.9, y: 0.9)
            })
        { (Bool) -> Void in
        }
    }

    /* ドラッグを感知した際に呼ばれるメソッド. (ドラッグ中何度も呼ばれる) */
    override func touchesMoved(_ touches: Set<UITouch>, with event: UIEvent?) {

        print("touchesMoved")

        // タッチイベントを取得.
        let aTouch: UITouch = touches.first!

        // 移動した先の座標を取得.
        let location = aTouch.location(in: self.view)

        // 移動する前の座標を取得.
        let prevLocation = aTouch.previousLocation(in: self.view)

        // CGRect生成.
        var myFrame: CGRect = self.view.frame

        // ドラッグで移動したx, y距離をとる.
        let deltaX: CGFloat = location.x - prevLocation.x
        let deltaY: CGFloat = location.y - prevLocation.y

        // 移動した分の距離をmyFrameの座標にプラスする.
        myFrame.origin.x += deltaX
        myFrame.origin.y += deltaY

        // frameにmyFrameを追加.
        self.view.frame = myFrame
    }

    /* 指が離れたことを感知した際に呼ばれるメソッド. */
    override func touchesEnded(_ touches: Set<UITouch>, with event: UIEvent?) {

        print("touchesEnded")

        // Labelアニメーション.
        UIView.animate(withDuration: 0.1,

                                   // アニメーション中の処理.
            animations: { () -> Void in
                // 拡大用アフィン行列を作成する.
                self.myLabel.transform = CGAffineTransform(scaleX: 0.4, y: 0.4)
                // 縮小用アフィン行列を作成する.
                self.myLabel.transform = CGAffineTransform(scaleX: 1.0, y: 1.0)
            })
        { (Bool) -> Void in

        }
    }

}
```

## モーダル表示させる

SecondViewController.swift
```swift
import UIKit

class SecondViewController : UIViewController{

    override func viewDidLoad() {

        self.view.backgroundColor = UIColor.black

        // もどるButtonを生成.
        let myButton = UIButton()
        myButton.frame = CGRect(x: 0, y: 0, width: 200, height: 40)
        myButton.backgroundColor = UIColor.red
        myButton.layer.masksToBounds = true
        myButton.setTitle("もどる", for: .normal)
        myButton.setTitleColor(UIColor.white, for: .normal)
        myButton.setTitleColor(UIColor.black, for: .highlighted)
        myButton.layer.cornerRadius = 20.0
        myButton.layer.position = CGPoint(x: self.view.frame.width/2, y:200)
        myButton.tag = 1
        myButton.addTarget(self, action: #selector(SecondViewController.onClickMyButton(sender:)), for: .touchUpInside)

        // viewにButtonを追加.
        self.view.addSubview(myButton)
    }

    /* Buttonを押した時に呼ばれるメソッド. */
    @objc func onClickMyButton(sender : UIButton){

        // viewを閉じる.
        self.navigationController?.dismiss(animated: true, completion: nil)
    }
}
```
ViewController.swift
```swift
import UIKit

class ViewController: UIViewController {



    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.


        // PopButtonを生成.
        let myButton = UIButton()
        myButton.frame = CGRect(x: 0, y: 0, width: 200, height: 40)
        myButton.backgroundColor = UIColor.red
        myButton.layer.masksToBounds = true
        myButton.setTitle("PopOver", for: .normal)
        myButton.setTitleColor(UIColor.white, for: .normal)
        myButton.setTitleColor(UIColor.black, for: .highlighted)
        myButton.layer.cornerRadius = 20.0
        myButton.layer.position = CGPoint(x: self.view.frame.width/2, y:200)
        myButton.tag = 0
        myButton.addTarget(self, action: #selector(ViewController.onClickMyButton(sender:)), for: .touchUpInside)

        // viewにButtonを追加.
        self.view.addSubview(myButton)
    }


    /* Buttonが押された時に呼ばれるメソッド. */
    @objc func onClickMyButton(sender : UIButton){

        // secondViewControllerのインスタンス生成.
        let second = SecondViewController()

        // navigationControllerのrootViewControllerにsecondViewControllerをセット.
        let nav = UINavigationController(rootViewController: second)

        // 画面遷移.
        self.present(nav, animated: true, completion: nil)
    }

}
```

## 現Viewの上に別のViewを表示

```swift
import UIKit

class ViewController: UIViewController {

    var myView: UIView!
    var myButton: UIButton!
    var flag: Bool!

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.


        // view表示・非表示のためのフラグ.
        flag = false

        // viewの背景を青色に設定.
        self.view.backgroundColor = UIColor.cyan

        // Viewを生成.
        myView = UIView(frame: CGRect(x: 0, y: 0, width: 100, height: 100))

        // myViewの背景を緑色に設定.
        myView.backgroundColor = UIColor.green

        // 透明度を設定.
        myView.alpha = 0.5

        // 位置を中心に設定.
        myView.layer.position = CGPoint(x: self.view.frame.width/2, y: self.view.frame.height/2)

        // myViewを非表示.
        myView.isHidden = true

        // ボタンを生成.
        myButton = UIButton(frame: CGRect(x: 0, y: 0, width: 100, height: 50))
        myButton.backgroundColor = UIColor.red
        myButton.layer.cornerRadius = 20.0
        myButton.layer.position = CGPoint(x: self.view.frame.width/2, y: self.view.frame.height-50)
        myButton.setTitle("Appear", for: .normal)
        myButton.setTitleColor(UIColor.white, for: .normal)
        myButton.addTarget(self, action: #selector(ViewController.onClickMyButton(sender:)), for: .touchUpInside)

        // myViewをviewに追加.
        self.view.addSubview(myView)

        // ボタンをviewに追加.
        self.view.addSubview(myButton)
    }


    /* ボタンイベント */
    @objc func onClickMyButton(sender: UIButton) {

        // flagがfalseならmyViewを表示.
        if !flag {
            // myViewを表示.
            myView.isHidden = false

            // ボタンのタイトルを変更.
            myButton.setTitle("Disappear", for: .normal)
            flag = true
        }
            // flagがtrueならmyViewを非表示.
        else {

            // myViewを非表示.
            myView.isHidden = true

            // ボタンのタイトルを変更.
            myButton.setTitle("Appear", for: .normal)
            flag = false
        }
    }

}
```
