# UINavigationController

## コードで生成

### AppDelegate.swift

```swift
import UIKit

@UIApplicationMain
class AppDelegate: UIResponder, UIApplicationDelegate {

    var window: UIWindow?
    var myNavigationController: UINavigationController?

    func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {

        let first = MainVC()
        myNavigationController = UINavigationController(rootViewController: first)
        self.window = UIWindow(frame: UIScreen.main.bounds)
        self.window?.rootViewController = myNavigationController
        self.window?.makeKeyAndVisible()


        return true
    }
}
```

### SceneDelegate.swift

```swift
import UIKit

@available(iOS 13.0, *)
class SceneDelegate: UIResponder, UIWindowSceneDelegate {

    var window: UIWindow?
    var myNavigationController: UINavigationController?

    func scene(_ scene: UIScene, willConnectTo session: UISceneSession, options connectionOptions: UIScene.ConnectionOptions) {
        guard let scene = (scene as? UIWindowScene) else { return }

        let first = TestViewController()
        myNavigationController = UINavigationController(rootViewController: first)
        self.window = UIWindow(windowScene: scene)
        self.window?.rootViewController = myNavigationController
        self.window?.makeKeyAndVisible()
    }

    func sceneDidDisconnect(_ scene: UIScene) {}
    func sceneDidBecomeActive(_ scene: UIScene) {}
    func sceneWillResignActive(_ scene: UIScene) {}
    func sceneWillEnterForeground(_ scene: UIScene) {}
    func sceneDidEnterBackground(_ scene: UIScene) {}
}
```

### MainVC.swift

```swift
import UIKit

class MainVC: UIViewController {

    var addBtn: UIBarButtonItem!

    override func viewDidLoad() {
        super.viewDidLoad()
        self.title = "home"
        view.backgroundColor = UIColor.gray

        addBtn = UIBarButtonItem(barButtonSystemItem: .add, target: self, action: #selector(onClick)   )
        self.navigationItem.rightBarButtonItem = addBtn

    }


    override func viewDidLayoutSubviews() {
        super.viewDidLayoutSubviews()
    }


    @objc func onClick(){
        let second = SecondVC()
        self.navigationController?.pushViewController(second, animated: true)
    }


}
```

### SubVC.swift

```swift
import UIKit

class SecondVC: UIViewController {

    var aBtn = UIButton()

    override func viewDidLoad() {
        super.viewDidLoad()
        self.title = "Second"
        self.view.backgroundColor = .cyan

        aBtn.addTarget(self, action: #selector(backToMain), for: .touchUpInside)
        aBtn.backgroundColor = .red
        aBtn.setTitle("Click me back", for: .normal)
        self.view.addSubview(aBtn)

        aBtn.translatesAutoresizingMaskIntoConstraints = false
        aBtn.centerYAnchor.constraint(equalTo: view.centerYAnchor).isActive = true
        aBtn.centerXAnchor.constraint(equalTo: view.centerXAnchor).isActive = true


    }


    @objc func backToMain(){
        self.navigationController?.popViewController(animated: true)
    }

}
```
