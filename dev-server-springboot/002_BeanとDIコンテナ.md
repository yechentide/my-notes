# BeanとDIコンテナ

Spring Frameworkでは、アプリケーション起動時に必要な設定を読み込み、インスタンスを生成してDIコンテナに保持する。  
==Bean==: DIコンテナで管理されているインスタンス

## DIコンテナの機能

- 必要なインスタンスの保持
- 必要があれば、あるBeanを別のBeanに代入(==DI==)
- 必要な箇所に、トランザクション管理などの割り込み処理を織り込む(==AOP==)
- Beanインスタンスのライフサイクルの管理
    - Beanインスタンスはデフォルトで==シングルトン==

## Bean定義

1. クラスにアノテーションを付ける
2. Java Configクラスにメソッドを作成する
3. 関数型Bean定義を利用する(新しくて使われないことが多い？)
4. XML設定ファイルに記述する(古すぎる)

基本的にはアノテーションを使うが、  
第三者が作ったライブラリ内のクラスなどアノテーションを使えない場合、  
Configクラスにメソッドを追加する形でBean定義をする  
この二つの方法は==併用できる==

## アノテーションを使う方法

- まずは依存ライブラリの導入: `org.springframework.boot`の`spring-boot-stater`ライブラリ
- 設定クラスの作成
    - `@Configuration`: このクラスがJava Configであることを示す
    - `@ComponentScan`: basepackageで指定したパッケージから、`@Component`がついているクラスをインスタンス化し、DIコンテナに保持
    ```java
    @Configuration
    @ComponentScan(basePackage = {  "パッケージ名", "パッケージ名"  })
    public class AppConfig {
        /*..........*/
    }
    ```
- Beanの作成
    - `@ComponentScan`で指定したパッケージの中に定義する必要がある
    ```java
    public interface DataLogic {
        public void search();
    }
    @Component
    public class MyBatisDataLogic implements DataLogic {
        @Override
        public void search() { /*..........*/ }
    }
    ```
    - インターフェースを作成する理由:
        - 後から実装クラスの差し替えが可能 (つまり処理を変えられる)
        - インターフェースを使わなくてもBeanを作れる
- DIコンテナの作成: `SpringApplication.run()`
    ```java
    // AppConfigが作成したJava Configクラスである
    public class MyAPP {
        public static void main(String[] args){
            ApplicationContext context = SpringApplication.run(AppConfig.class, args);
            // Java Configを読み込んで
            // コンポーネントをスキャン
            // Beanインスタンスを生成
            // 生成したBeanをDIコンテナに保存
        }
    }
    ```
- Beanの取り出し: `getBean()`
    ```java
    // 引数には、取得したいBeanの型を与える (シングルトンだから型でわかる)
    DataLogic dataLogic = context.getBean(DataLogic.class);
    // 取り出した後は普通にインスタンスメソッドを呼び出して処理させる
    dataLogic.search();
    ```

## Configクラスにメソッドを作成する方法

アノテーションを使えない場合(他人が書いたクラスを扱う場合)はこの方法を使う
- 依存ライブラリを導入
- クラスに`@Componentをつけない`
- 設定クラスの作成
    - `@ComponentScanをつけない`
    - Beanインスタンスを返すメソッドを定義し、`@Bean`をつける
    ```java
    @Configuration
    public class AppConfig {
        /*..........*/
        @Bean
        public DataLogic datalogic() { return new MyBatisDatalogic(); }
    }
    ```
