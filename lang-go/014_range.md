# Range

## Rangeの基本

`range` は、`Slices(スライス)` や、`Maps(マップ)` をひとつずつ反復処理するために使われる

indexやvalueを使わない場合は、変数名を`_`にすることで明示できる

```go
for index, value := range aSlice {
    fmt.Println(index, value)
}
```
