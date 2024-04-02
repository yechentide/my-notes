# Spring MVC

## Spring MVCについて

- ==Spring MVC==はSpring Framework関連のプロダクトの一つ
- プロジェクトの初期構築は[Webサイト](https://start.spring.io)で行える
- 依存ライブラリは`spring-boot-starter-web` (Spring MVC + Tomcat + ...)
- JSONの変換ライブラリの一つとして、Jacksonがある
- ==application.properties==に設定を書く
    - JSONの整形: `spring.jackson.serialization.indent-output=true`

### レスポンスクラス・リクエストクラスの定義

```json
{
    "id": 101,
    "name": "山田太郎",
    "email": "yamada@example.com"
}
```
```java
public class EmployeeResponse {
    private Integer id;
    private String name;
    private String email;
    
    // コンストラクタ (引数あり)

    // getterを準備
}
public class EmployeeRequest {
    // private Integer id;
    private String name;
    private String email;
    
    // コンストラクタ (引数なし / デフォルトコンストラクタ)

    // getterを準備
    // setterを準備
}
```

### 実行

- `@SpringBootApplication` = `@Configuration` + `@ComponentScan` + `@EnableAutoConfiguration`
    - `@ComponentScan`でbasepackagesを指定してない場合、スキャン範囲は今のパッケージの配下
        - 所属パッケージ: `com.example`
        - スキャン範囲: `com.example.*`
        - スキャンパッケージの例: com.example, com.example.web.rest
```java
@SpringBootApplication
public class EduSpringRestIntroApplication {
    public static void main(String[] args) {
        SpringApplication.run(EduSpringRestIntroApplication.class, args)
    }
}
```

### コントローラクラスの定義

- `@Controller`: コントローラクラスにつける。実体は`@Component`である
- `@RequestMapping`: リクエストのパスを指定する
- `@GetMapping`: 指定したURIにGETリクエストが来たときに実行するメソッドにつける
- `@ResponseBody`: メソッドの戻り値を適切なフォーマットに変換し、レスポンスとしてクライアントに返す
```java
@Controller
@RequestMapping("/employees")    // リクエストのパスを指定
public class EmployeeRestController {
    @GetMapping
    @ResponseBody
    public List<EmployeeResponse> findAll() {
        // 処理
    }
}
```

- `@RestController` = `@Controller` + `@ResponseBody`
    - レスポンスを自動的にJSONに変換するため、`@ResponseBody`をつけずに済む
```diff
+ @RestController
- @Controller
  @RequestMapping("/employees")    // リクエストのパスを指定
  public class EmployeeRestController {
      @GetMapping
-     @ResponseBody
      public List<EmployeeResponse> findAll() {}
  }
```

- パラメータの受け取り
    - `@GetMapping`の中に追加のパスを指定する。取得したい値を`{}`で囲む
    - `@PathVariable`でパスの中から値を取得する, `{}`内の名前と引数名は同じにする必要がある
```java
@RestController
@RequestMapping("/employees")    // リクエストのパスを指定
public class EmployeeRestController {
    @GetMapping("/{id}")
    public EmployeeResponse findById(@PathVariable Integer id) {}
}
```

### Mappingアノテーションについて

Spring Framework 2.5から使える
- `@RequestMapping`: 完全形は`@RequestMapping(path="/employees", method=RequestMethod.GET)`

Spring Framework 4.3から使える、`@RequestMapping`をラップしたアノテーション
- `@GetMapping`
- `@PostMapping`
- `@PutMapping`
- `@DeleteMapping`
- `@PatchMapping`

### 更新用コントローラ(Putメソッド)

- `@ResponseStatus`: メソッド処理が正常終了した際に、クライアントに送るレスポンスコードを設定
- `@RequestBody`: クライアントから送られてきたリクエストボディを引数と関連づける
```java
@RestController
@RequestMapping("/employees")    // リクエストのパスを指定
public class EmployeeRestController {
    @PutMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)   // コードは204
    public void update(@PathVariable Integer id, @RequestBody EmployeeRequest request) {
        // 引数requestからデータを抽出
        // 引数idと対応するデータを更新
    }
}
```

## アノテーションの紹介

- `@RequiredArgsConstructor`
    - (プレゼンテーション層)コントローラクラス、(ビジネスロジック層)サービスクラスに付けるとよい
    - このアノテーションをクラスに付与することにより、finalなフィールドを初期化するコンストラクタが自動生成される
        - つまりfinal(定数)フィールドしか持たない場合、コンストラクタを書かなくて済む
- `@Value`
    - リクエストクラスに付けるとよい
    - 各フィールドのGetterを自動生成してくれる
- `@Data`
    - レスポンスクラス、entityクラスに付けるとよい
    - 各フィールドのGetter、Setterなどを自動生成してくれる
- `@Builder`
    - レスポンスクラス、リクエストクラスに付けるとよい
    - これを付けたクラスは、メソッドチェーン形式で、フィールドを１つずつ初期化して、クラスのインスタンスを作れる
- `@JsonProperty("snake_case")`
    - レスポンスクラスに付けるとよい
    - 変数に`@JsonProperty("snake_case")`を付与することでsnake_caseに変換
- `@Jacksonized`
    - リクエストクラスに付けるとよい
    - Requestクラスのインスタンスを生成する際のデシリアライズエラーを回避できる
