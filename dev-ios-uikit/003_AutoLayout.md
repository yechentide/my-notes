# Auto Layout

## コードでAuto Layoutを作る

### 前提

- `view.addSubview(部品)`  
    Auto Layoutでは、部品同士は共通の祖先を持っている必要があるため  
    制約を付けるのは`addSubview`の後でなければならない  
    `addSubview`より前に制約を付けようとすると実行時にエラーが発生してしまう
- `部品.translatesAutoresizingMaskIntoConstraints=false`  
    コードからUI部品を生成すると  
    デフォルトで`translatesAutoresizingMaskIntoConstraints`がtrueになってしまう  
    このプロパティは、AutoresizingMaskをAuto Layoutの制約に置き換えるかどうか指定する値で  
    trueのままだと意図しない制約がついてしまう

### オススメのやり方

- [コードで作るAuto Layoutまとめ](https://medium.com/@shiba1014/%E3%82%B3%E3%83%BC%E3%83%89%E3%81%A7%E4%BD%9C%E3%82%8Bauto-layout%E3%81%BE%E3%81%A8%E3%82%81-274f14043393)
- [Auto Layoutをコードで書いてみた](https://qiita.com/dddisk/items/8001598ea7951bcdcc30)
1. やり方1
   ```swift
    myImageView.translatesAutoresizingMaskIntoConstraints = false
    myImageView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor, constant: 10.0).isActive = true
    myImageView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor, constant: -10.0).isActive = true
    myImageView.topAnchor.constraint(equalTo: self.view.topAnchor, constant: 30.0).isActive = true
    myImageView.bottomAnchor.constraint(equalTo: self.view.bottomAnchor, constant: -30.0).isActive = true
    ```
2. やり方2
    ```swift
    myImageView.translatesAutoresizingMaskIntoConstraints = false
    NSLayoutConstraint.activate([
        myImageView.leadingAnchor.constraint(equalTo: self.view.leadingAnchor, constant: 10.0),
        myImageView.trailingAnchor.constraint(equalTo: self.view.trailingAnchor, constant: -10.0),
        myImageView.topAnchor.constraint(equalTo: self.view.topAnchor, constant: 30.0),
        myImageView.bottomAnchor.constraint(equalTo: self.view.bottomAnchor, constant: -30.0)
    ])
    ```
