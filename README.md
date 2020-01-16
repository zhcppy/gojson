# GoJSON

Json serialization with color, color use of `github.com/fatih/color` implementation, most of the code from `https://github.com/hokaccha/go-prettyjson`

## Use

```bash
go get -v -u github.com/zhcppy/gojson
```

```go
import (
    "fmt"
    "github.com/zhcppy/gojson"
)

func main() {
    var data = map[string]interface{}{
        "name": "gojson",
    }
    res, err := gojson.Marshal(data)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(res)
}
```

## Run shows

```bash
git clone https://github.com/zhcppy/gojson
cd gojson
go test -v -test.run Test_Marshal
```