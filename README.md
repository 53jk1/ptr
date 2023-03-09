# Ptr

# Usage example

```go
package main

import (
    "fmt"
    "github.com/53jk1/ptr"
)

func main() {
    var i int = 1
    var s string = "hello"
    var b bool = true
    var f float64 = 1.1
    var c complex128 = 1 + 1i

    fmt.Println(*ptr.Int(i))
    fmt.Println(*ptr.String(s))
    fmt.Println(*ptr.Bool(b))
    fmt.Println(*ptr.Float64(f))
    fmt.Println(*ptr.Complex128(c))
}
```

# License

MIT