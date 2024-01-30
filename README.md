# Go logic gates
## Description
This is a simple implementation of logic gates in Go meant for learning purposes.

## Usage
```go
package main

import (
    "fmt"
    gates "github.com/torbenconto/logic-gates"
)

func main() {
    fmt.Println(gates.And(true, false))
}
```

## Gates
- And
- Or
- Not
- Nand
- Nor
- Xor
- Xnor