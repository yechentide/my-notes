# 脱Storyboard

## 参考リンク

- [Storyboardを出来るだけ使わずに書き始める](https://qiita.com/nagisawks/items/222c881d6798c46a390f)
- [【Xcode11】Storyboardを使わずコードだけで画面を生成する方法](https://qiita.com/edasan/items/68cbe9ab63d48ee71594)

## やり方

1. Main.storyboardを削除
2. プロジェクトのTARGETS > Infoにある`Main storyboard file base name`という項目を削除  
    または、TARGETにある、Deployment Info / Main Interface にある Main を空欄にする
3. プロジェクトのTARGETS > Infoにある`Application Scene Manifest` > `Scene Configuration` > `Application Session Role` > `Item0(Default Configura...` > `Storyboard Name`を削除
4. AppDelegate.swiftで初回の起動画面を設定
    ```swift
    class AppDelegate: UIResponder, UIApplicationDelegate {

        var window: UIWindow?

        func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {
            window = UIWindow(frame: UIScreen.main.bounds)
            window?.rootViewController = ViewController()
            window?.makeKeyAndVisible()
            return true
        }
    }
    ```
5. SceneDelegate.swiftで初回の起動画面を設定
    ```swift
    class SceneDelegate: UIResponder, UIWindowSceneDelegate {

        var window: UIWindow?

        func scene(_ scene: UIScene, willConnectTo session: UISceneSession, options connectionOptions: UIScene.ConnectionOptions) {
            guard let scene = (scene as? UIWindowScene) else { return }
            window = UIWindow(windowScene: scene)
            window?.rootViewController = ViewController()
            window?.makeKeyAndVisible()
        }
    }
    ```
