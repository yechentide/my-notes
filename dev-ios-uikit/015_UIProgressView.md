# UIProgressView

## UIProgressViewクラス

[UIProgressView](https://developer.apple.com/documentation/uikit/UIProgressView)

### 継承関係

```swift
class UIProgressView : UIView       // : UIResponder : NSObject
```

### コンストラクタ

```swift
init()
init(progressViewStyle: UIProgressView.Style)
init(frame: CGRect)
init?(coder: NSCoder)
```

### クラス定義

```swift
open class UIProgressView : UIView, NSCoding {

    public init(frame: CGRect)

    public init?(coder: NSCoder)

    public convenience init(progressViewStyle style: UIProgressView.Style) // sets the view height according to the style

    open var progressViewStyle: UIProgressView.Style // default is UIProgressViewStyleDefault

    open var progress: Float // 0.0 .. 1.0, default is 0.0. values outside are pinned.

    @available(iOS 5.0, *)
    open var progressTintColor: UIColor?

    @available(iOS 5.0, *)
    open var trackTintColor: UIColor?

    @available(iOS 5.0, *)
    open var progressImage: UIImage?

    @available(iOS 5.0, *)
    open var trackImage: UIImage?

    @available(iOS 5.0, *)
    open func setProgress(_ progress: Float, animated: Bool)

    @available(iOS 9.0, *)
    open var observedProgress: Progress?
}
```
