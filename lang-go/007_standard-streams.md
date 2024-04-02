# 標準出入力

## 標準出力

### コンソールに出力

```go
package main
import "fmt"

func main() {
    fmt.Print("hello, world\n")
    fmt.Println("hello, world")
    name := "yechentide"
    fmt.Printf("My name is %s", name)
}
```

## 標準入力

### Command Line Tool

何種類も方法があるっぽい
```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    in := scanner.Text()
    fmt.Println("input: ", in)
}
```
