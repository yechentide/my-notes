# JUnit

## ユニットテスト

納期に間に合わない場合、テストの期間が短くなるケースがあるため、  
実装時からユニットテストを行ってテストを前倒しにする方が良い
- ユニットテスト: テストコードを使って、プログラマーが行う小さな単位の動作確認テスト
- ユニットテストの利点
    1. 意図した通りに動くか、常に即座に確認できる
    2. 最初に手周防コードを書くコストがいるが、テストの実行はコストなく何度も行える
    3. コードがどのような仕様に基づいて作成されたのかを表現できる
    4. 自分が書いたコードに自信を持てる
    5. これから書くコードに自信を持てる
- Java用のテスティングフレームワークはJUnit
    - その拡張はMockito, DbUnit, Jacocoがある

## JUnitの使い方①

- JUnitを使う前に、まず`junit-jupiter`という依存関係をMavenなどに追加する必要がある
- テストしたいクラスを開き、右クリックで`Generate`を選択して、テストを生成できる
- `import static`を使うことで、`クラス.メソッド()`ではなくメソッド名だけで呼び出せるようになる
- テストコードは、テスト対象クラスと同じパッケージにする
    - 修飾子は無指定または`protected`

### 例

```java
class OrderItemTest {
    @BeforeAll
    static void aaa() {
        // すべてのテストケースが実行される前に、一回だけ実行される(51ページ)
    }

    @AfterAll
    static void bbb() {
        // すべてのテストケースが実行される後に、一回だけ実行される(51ページ)
    }

    @BeforeEach
    void ccc() {
        // 各テストケースが実行される前に、実行される(49ページ)
    }

    @AfterEach
    void ddd() {
        // 各テストケースが実行される後に、実行される(49ページ)
    }

    @Test
    void テスト名() {
        OrderItem item = new OrderItem();
        型 実際値 = item.メソッド();
        assertEquals(期待値, 実際値);
    }
    
    @Test
    @DisplayName("デバッグエリアに表示させたいテスト名")
    @Disabled("このテストを無視する理由(48ページ)")
    void テスト名() {
        OrderItem item = new OrderItem();
        型 実際値 = item.メソッド();
        assertEquals(期待値, 実際値);
    }
}
```

### assertメソッドの種類

- `assertEquals(期待値, 実際値)`
- `assertNotEquals(期待値, 実際値)`
- `assertTrue(実際値)`
- `assertFalse(実際値)`
- `assertNull(実際値)`
- `assertNoNull(実際値)`
- `assertThrows(例外クラス.class, アロー関数)`
    ```java
    assertThrows(
        SQLException.class,
        () -> { /*処理*/ }
    );
    ```
- `assertTimeout(制限時間, アロー関数)`

## テストの抽出

- ブラックボックステスト: 仕様に注目してテストを行う
    - 同値分割: ==同じ結果になる値をグループ化==
    - 境界値分析: ==各グループの限界値を抽出==
    - エラー推測: ==エラーになりそうな値もテストする==
- ホワイトボックステスト: メソッドに注目してテストを行う

## JUnitの使い方②

テストデータをまとめて準備する
```java
class NumberPack {
    int a1; int a2; int a3;  // テストデータ
    int expected;            // 期待値
}

class SorterParamTest {
    private Sorter s = new Sorter();

    private static List<NumberPack> getNumberPackList() {
        return List.of(
                new NumberPack(1, 2, 3, 3),
                new NumberPack(30, 20, 10, 30),
                new NumberPack(200, 300, 100, 300)
        );
    }

    @ParameterizedTest
    @MethodSource("getNumberPackList")
    void findBigTheory(NumberPack pack) {
        int result = s.findBig(pack.a1, pack.a2, pack.a3);
        assertEquals(pack.expected, result);
    }
}
```

テストケースをグループ化
```java
class SorterEnclosedTest {

    private Sorter s = new Sorter();

    @Nested
    class 三つが同じ場合 {

        @Test
        void exceptionWhenSame3number() {
            assertThrows(
                    IllegalArgumentException.class,
                    () -> s.findBig(1, 1, 1)
            );
        }
    }

    @Nested
    class 三つの中の二つが同じ場合 {

        @Test
        void same2Number115() {
            int result = s.findBig(1, 1, 5);
            assertEquals(5, result);
        }

        @Test
        void same2Number511() {
            int result = s.findBig(5, 1, 1);
            assertEquals(5, result);
        }
    }
}
```

## 良いテストコードの条件

- 自動化されている
- 常に同じ結果になる
- 他のテストに依存しない
- 十分なテスト項目
- わかりやすい名前
- 遅すぎない

## モックを利用したテスト(Mockito)

JUnitと同じく、Mockitoを使うにはMavenなどに依存関係を記述し、`import static`する
```java
class AAATest {
    インターフェース env = mock(インターフェース.class);

    @Test
    void testCase001() {
        when(  env.メソッド()  ).thenReturn(値);            // 特定のメソッドが呼び出される際にセットする値を指定
        when(  env.メソッド("文字列")  ).thenReturn(値);     // 引数が特定の文字列の場合
        when(  env.メソッド(anyString())  ).thenReturn(値); // 引数が任意の文字列の場合

        when(  env.メソッド()  ).thenThrow(  new RuntimeException()  );

        doThrow(  new NullpointerException()  ).when(env).戻り値がないメソッド();
    }
}
```

## テストの容易性を考えた設計とは

- 他のコンポーネントとの依存関係を最小限にする
- 依存しているコンポーネントをテスト用のものに差し替えられる
    - インターフェースや継承などを前提にした設計(==ポリモーフィズム==を利用)
    - 外部から依存コンポーネントをセットできるようにする
- 参照透過性を確保する
    - 戻り値が引数んみによって決定される
    - 引数が同じなら常に戻り値は同じである
- しかし大抵のシステムはファイルIOやデータベースCRUDなど副作用があるものを必要とする
    - 参照透過性を確保できるものと副作用があるものを意識的に分けて設計すべき

## DbUnit

- `IDatabaseTester`の実装クラス
    - `JdbcDatabaseTester`
    - `PropertiesBasedJdbcDatabaseTester`
    - `DataSourceDatabaseTester`
    - `JndiDatabaseTester`

```java
class MemberDaoTest {

    private static IDatabaseTester tester = null;

    @BeforeAll
    static void beforeTest() throws Exception {
        tester = new JdbcDatabaseTester("ドライブクラスの文字列", "DBのURL", "ユーザ名", "パスワード", "DBの名前") {
            // MySQL 8.0の場合、以下のように書かないといけない
            @Override
            public IDatabaseConnection getConnection() throws Exception {
                return new MySqlConnection(super.getConnection().getConnection(), getSchema());
            }
        };
    }

    @Test
    void testAddOneMember() throws Exception {
        // 準備：初期データの設定
        //       テスト対象のオブジェクトの用意
        tester.setDataSet(new XlsDataSet(new File("src/test/resources/initial_dataset.xlsx")));
        tester.onSetup();
        MemberDao dao = new MemberDao();
        Member m = new Member("tarou@mail.com", "太郎", "kibidango");

        // 実行：テスト対象のメソッドの実行
        dao.add(m);

        //検証①：テスト対象メソッド実行後の実際のテーブル内の状態を取得
        IDatabaseConnection conn = tester.getConnection();
        ITable actualTable = conn.createTable("members");

        //検証②： 期待するデータセットをExcelファイルから取得
        IDataSet expectedDataSet = new XlsDataSet(new File("src/test/resources/expected_dataset.xlsx"));
        //検証③： 続いて期待するテーブルの状態を取得
        ITable expectedTable = expectedDataSet.getTable("members");

        //検証④：テーブルの実際の状態と期待する状態を比較
        Assertion.assertEquals(expectedTable, actualTable);

    }

    @AfterEach
    void tearDown() throws Exception {
        tester.onTearDown();
    }

}

```

### データの準備や後処理の方法を変える場合

普通は`onSetUP`と`onTearDown`を呼び出せば十分
```java
@BeforeAll
static void aaaa() {
    tester = new JdbcDatabaseTester(/*.......*/);
    tester.setSetUpOperation(DatabaseOperation.INSERT);    // デフォルトの設定
    tester.setTearDownOperation(DatabaseOperation.DELETE); // デフォルトの設定
    // UPDATE, INSERT, DELETE, DELETE_ALL
    // CLEAN_INSERT, REFRESH, TRUNCATE_TABLE, NONE
}
```

### フィルタリング

actualTableの方にいらない列がある場合、フィルターリングを行う
```java
ITable actualTable = tester.getConnection().createDataSet().getTable("テーブル名");
ITable expectedTable = new XlsDataSet(new File("aaa.xlsx")).getTable("テーブル名");

ITable filteredTable =
    DefaultColumnFilter.includedColumnsTable(actualTable, expectedTable.getTableMetaData().getColumns());

Assertion.assertEquals(expectedTable, filteredTable);
```

### データセットの置き換え

```java
IDataSet original = new FlatXlsDataSet(new FileReader("aaa.xlsx"));
ReplacementDataSet replaced = new ReplacementDataSet(original);
replaced.addReplacementObject("[null]", null);   // null を 文字列の[null] に変換
```

## プログラムでテストデータセットを作る

```java
void 全テーブルの全データを出力() throws Exception {
    Connection con = DriverManager.getConnection("URL", "ユーザ名", "パスワード");
    IDatabaseConnection icon = new DatabaseConnection(con);

    IDataSet fullDataSet = icon.createDataSet();
    XlsDataSet.write(fullDataSet, new FileOutputStream("ファイルのパス"));
}
    
void 特定のテーブルの全データを出力() throws Exception {
    Connection con = DriverManager.getConnection("URL", "ユーザ名", "パスワード");
    IDatabaseConnection icon = new DatabaseConnection(con);
    String[] tableNames = {"players"};

    IDataSet dataSet = icon.createDataSet(tableNames);
    XlsDataSet.write(dataSet, new FileOutputStream("ファイルのパス"));
}
```
