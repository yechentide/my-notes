# ファイル出入力

## ファイル操作

### 書き込み

### 読み込み

必要に応じて1行ずつ読み込む
```go
import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    file, err := os.Open("fileName")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error")
    }
    input := bufio.NewScanner(file)
    for lineNum := 1; input.Scan(); lineNum++ {
        fmt.Printf("%d行目: %s\n", lineNum, input.Text())
    }
}
```
一気に読み込んで、メモリ上に保持
```go
import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    data, err := ioutil.ReadFile("aaa.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }
    for lineNum, line := range strings.Split(string(data), "\n") {
        fmt.Printf("%d行目: %s\n", lineNum, line)
    }
}
```
