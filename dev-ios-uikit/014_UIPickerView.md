# UIPickerView

## UIPickerViewクラス

[UIPickerView](https://developer.apple.com/documentation/uikit/UIPickerView)  
複数の選択肢を回転ドラム式のコンポーネントで表示し、  
コンポーネントごとに値を１つ選ぶための部品

### 継承関係

```swift
class UIPickerView : UIView     // : UIResponder : NSObject
```

### コンストラクタ

```swift
init()
init(frame: CGRect)
```

### クラス定義

```swift
open class UIPickerView : UIView, NSCoding {

    weak open var dataSource: UIPickerViewDataSource? // default is nil. weak reference

    weak open var delegate: UIPickerViewDelegate? // default is nil. weak reference

    @available(iOS, introduced: 2.0, deprecated: 13.0, message: "This property has no effect on iOS 7 and later.")
    open var showsSelectionIndicator: Bool

    // info that was fetched and cached from the data source and delegate
    open var numberOfComponents: Int { get }

    open func numberOfRows(inComponent component: Int) -> Int

    open func rowSize(forComponent component: Int) -> CGSize

    // returns the view provided by the delegate via pickerView:viewForRow:forComponent:reusingView:
    // or nil if the row/component is not visible or the delegate does not implement
    // pickerView:viewForRow:forComponent:reusingView:
    open func view(forRow row: Int, forComponent component: Int) -> UIView?

    // Reloading whole view or single component
    open func reloadAllComponents()

    open func reloadComponent(_ component: Int)

    // selection. in this case, it means showing the appropriate row in the middle
    open func selectRow(_ row: Int, inComponent component: Int, animated: Bool) // scrolls the specified row to center.

    open func selectedRow(inComponent component: Int) -> Int // returns selected row. -1 if nothing selected
}
```

### UIPickerViewDataSource

[UIPickerViewDataSource](https://developer.apple.com/documentation/uikit/uipickerviewdatasource)

```swift
public protocol UIPickerViewDataSource : NSObjectProtocol {

    // returns the number of 'columns' to display.
    @available(iOS 2.0, *)
    func numberOfComponents(in pickerView: UIPickerView) -> Int

    // returns the # of rows in each component..
    @available(iOS 2.0, *)
    func pickerView(_ pickerView: UIPickerView, numberOfRowsInComponent component: Int) -> Int
}
```

### UIPickerViewDelegate

[UIPickerViewDelegate](https://developer.apple.com/documentation/uikit/uipickerviewdelegate)

```swift
public protocol UIPickerViewDelegate : NSObjectProtocol {

    // returns width of column and height of row for each component.
    @available(iOS 2.0, *)
    optional func pickerView(_ pickerView: UIPickerView, widthForComponent component: Int) -> CGFloat

    @available(iOS 2.0, *)
    optional func pickerView(_ pickerView: UIPickerView, rowHeightForComponent component: Int) -> CGFloat

    // these methods return either a plain NSString, a NSAttributedString, or a view (e.g UILabel) to display the row for the component.
    // for the view versions, we cache any hidden and thus unused views and pass them back for reuse.
    // If you return back a different object, the old one will be released. the view will be centered in the row rect
    @available(iOS 2.0, *)
    optional func pickerView(_ pickerView: UIPickerView, titleForRow row: Int, forComponent component: Int) -> String?

    @available(iOS 6.0, *)
    optional func pickerView(_ pickerView: UIPickerView, attributedTitleForRow row: Int, forComponent component: Int) -> NSAttributedString? // attributed title is favored if both methods are implemented

    @available(iOS 2.0, *)
    optional func pickerView(_ pickerView: UIPickerView, viewForRow row: Int, forComponent component: Int, reusing view: UIView?) -> UIView


    @available(iOS 2.0, *)
    optional func pickerView(_ pickerView: UIPickerView, didSelectRow row: Int, inComponent component: Int)
}
```
