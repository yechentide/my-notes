# その他

## 参考サイト

- [Golang 入門](https://qiita.com/y-kaanai/items/e022b7c316cd8a6bb6d2)
- [【Go】基本文法](https://qiita.com/k-penguin-sato/items/1d0e1c6b4bf937996cd3)

## ネットワーク

```go
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("引数の数が違います")
        os.Exit(1)
    }

    url := os.Args[1]
    response, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "[error] %v\n", err)
        os.Exit(1)
    }

    body, err := ioutil.ReadAll(response.Body)
    response.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "[error] %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("%s", body)

}
```

## ゴルーチン＆チャンネル

```go
import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    ch := make(chan string)

    for _, url := range os.Args[1:] {
        go fetch(url, ch) // goroutine開始
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // chチャンネルから受信
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    response, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    bytes, err := io.Copy(ioutil.Discard, response.Body)
    response.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    seconds := time.Since(start).Seconds()
    ch <- fmt.Sprintf("時間：%.2fs、サイズ：%7d、From：%s", seconds, bytes, url)
}
```

## Webサーバ

```go
import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler) // 個々のリクエストに対してhandler関数が呼ばれる
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handlerは、リクエストされたURLのPath要素を返す
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "method = %q\n", r.Method)
    fmt.Fprintf(w, "url = %q\n", r.URL)
    fmt.Fprintf(w, "protocol = %q\n\n", r.Proto)

    for k, v := range r.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "Remote Address = %q\n", r.RemoteAddr)

    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }
}
```

## JSON

JSONのエンコーディングとデコーディングに使うパッケージ：`encoding/json`  
Goのデータ構造をJSONに変換することを、マーシャリング(marshaling)という <---> アンマーシャリング(unmarshaling)  
公開されているフィールドだけ変換される  

### フィールドタグ

フィールドタグによって、Goのフィールドに対する、JSONでの代替名を指定できる  
フィールドタグはどんな文字列でも良いが、慣習的に`key:"value"`のペアの、空白区切りのリストとして解釈される  
`"`が入っているので、大体`` `key:"value"` ``の形で書く  
追加のオプション`omitempty`は、そのフィールドがゼロ値or空であれば、JSONへ変換しないことを意味する

```go
type Movie struct {
    // フィールド宣言の後ろの文字列 = フィールドタグ
    Title  string
    Year   int  `json:"released"`
    Color  bool `json:"color,omitempty"`
    Actors []string
}

var movies = []Movie{
    {Title: "A", Year: 2020, Color: false, Actors: []string{"a1", "a2"}},
    {Title: "B", Year: 2021, Color: true, Actors: []string{"b1", "b2"}},
}
```

### マーシャリング

```go
// json.Marshal()は余分な空白を取り除いた長い1行になる
// json.MarshalIndent()の方が読みやすい
// data, err := json.MarshalIndent(movies, "各行の接頭辞", "インデント")
data, err := json.Marshal(movies)
if err != nil {
    log.Fatalf("JSONに変換失敗：%s", err)
}
fmt.Printf("%s\n", data)
```

### アンマーシャリング

アンマーシャリングの過程で、JSONの名前をGoの構造体フィールド名との関連付けるは、大文字小文字を区別しない

```go
// １つのフィールドしかない構造体のスライスへアンマーシャリング
// JSONのTitleという項目だけデコード、それ以外は破棄
var titles []struct{ Title string }
if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSONから変換失敗：%s", err)
}
fmt.Println(titles)
```
