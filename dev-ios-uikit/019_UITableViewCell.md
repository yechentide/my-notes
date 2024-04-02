# UITableViewCell

## カスタムセル

- [UITableViewCellの高さを可変にする](https://blog.mothule.com/ios/uitableview/ios-uitableview-uitableviewcell-height-customize)
- [カスタムセルの作り方](https://turedureengineer.hatenablog.com/entry/2018/12/06/170045)
```swift
class MainVC: UIViewController, UITableViewDataSource, UITableViewDelegate {
    override func viewDidLoad() {
        table.dataSource = self
        table.delegate = self
        table.register(MyTableCell.self, forCellReuseIdentifier: "myCell")
        table.tableFooterView = UIView()
        self.view.addSubview(table)
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "myCell") as! MyTableCell
        cell.myLabel.text = humanList[indexPath.row].name
        cell.myImageView.image = humanList[indexPath.row].photo
        cell.accessoryType = .disclosureIndicator
        return cell
    }
}
```

```swift
import UIKit

class MyTableCell: UITableViewCell {

    var myLabel: UILabel!
    var myImageView: UIImageView!

    override func awakeFromNib() {
        super.awakeFromNib()
        // Initialization code
    }

    override func setSelected(_ selected: Bool, animated: Bool) {
        super.setSelected(selected, animated: animated)
        // Configure the view for the selected state
    }

    override init(style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
        super.init(style: style, reuseIdentifier: reuseIdentifier)

        myLabel = UILabel()
        myLabel.textAlignment = .left
        contentView.addSubview(myLabel)

        myImageView = UIImageView()
        myImageView.contentMode = .scaleAspectFit
        myImageView.backgroundColor = .gray
        contentView.addSubview(myImageView)
    }

    required init(coder aDecoder: NSCoder) {
        fatalError("init(coder: ) has not been implemented")
    }

    override func layoutSubviews() {
        super.layoutSubviews()

        myLabel.frame = CGRect(x: 80, y: 0, width: frame.width - 100, height: frame.height)
        myImageView.translatesAutoresizingMaskIntoConstraints = false
        myImageView.centerYAnchor.constraint(equalTo: self.contentView.centerYAnchor).isActive = true
        myImageView.leadingAnchor.constraint(equalTo: self.contentView.leadingAnchor, constant: 10).isActive = true
        myImageView.widthAnchor.constraint(equalToConstant: 60).isActive = true
        myImageView.heightAnchor.constraint(equalToConstant: frame.height*0.8).isActive = true
    }
}

```
