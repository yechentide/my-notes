# UITabBarController

## TabBarとNavigation

### AppDelegate.swift

```swift
import UIKit

@UIApplicationMain
class AppDelegate: UIResponder, UIApplicationDelegate {

    var window: UIWindow?
    var viewController = MyTabbarC()

    func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {

        self.window = UIWindow(frame: UIScreen.main.bounds)
        self.window?.rootViewController = viewController
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
    var viewController = MyTabbarC()

    func scene(_ scene: UIScene, willConnectTo session: UISceneSession, options connectionOptions: UIScene.ConnectionOptions) {
        guard let scene = (scene as? UIWindowScene) else { return }

        self.window = UIWindow(windowScene: scene)
        self.window?.rootViewController = viewController
        self.window?.makeKeyAndVisible()
    }

    func sceneDidDisconnect(_ scene: UIScene) {}
    func sceneDidBecomeActive(_ scene: UIScene) {}
    func sceneWillResignActive(_ scene: UIScene) {}
    func sceneWillEnterForeground(_ scene: UIScene) {}
    func sceneDidEnterBackground(_ scene: UIScene) {}
}
```

### MyTabbarC.swift

```swift
import UIKit

class MyTabbarC: UITabBarController {

    override func viewDidLoad() {
        super.viewDidLoad()

        // 背景色
        tabBar.barTintColor = UIColor(red: 0.5, green: 0.8, blue: 0.9, alpha: 0.8)
        // アイテムの色
        tabBar.tintColor = UIColor.white

        let firstVC = MyNaviC(rootViewController: MainVC())
        firstVC.tabBarItem = UITabBarItem(tabBarSystemItem: .bookmarks, tag: 0)

        let secondVC = MyNaviC(rootViewController: SecondVC())
        secondVC.tabBarItem = UITabBarItem(tabBarSystemItem: .contacts, tag: 1)

        let thirdVC = MyNaviC(rootViewController: ThirdVC())
        thirdVC.tabBarItem = UITabBarItem(tabBarSystemItem: .more, tag: 2)

        self.viewControllers = [firstVC, secondVC, thirdVC]

    }

}
```

### MyNaviC.swift

```swift
import UIKit

class MyNaviC: UINavigationController {

    override func viewDidLoad() {
        super.viewDidLoad()

        // 背景色
        navigationBar.barTintColor = UIColor(red: 0.5, green: 0.8, blue: 0.9, alpha: 0.8)
        // アイテムの色
        navigationBar.tintColor = .blue
        // テキスト
        navigationBar.titleTextAttributes = [.foregroundColor: UIColor.black]
    }
}
```

### MainVC.swift

### SecondVC.swift

### ThirdVC.swift
