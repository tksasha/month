# Month
```go
package main

import (
  "fmt"

  "github.com/tksasha/month"  
)

func main() {
  m := month.New(2024, 2)

  fmt.Println(m.Begin()) // 2024-02-01 00:00:00.00 +0000 UTC

  fmt.Println(m.End()) // 2024-02-29 23:59:59.999999999 +0000 UTC
} 
```
