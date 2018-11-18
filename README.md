# faketime
Mock time.Now() in golang

## Examples

time.Date style

```go
package main

import (
	"fmt"
	"github.com/tkuchiki/faketime"
	"time"
)

func main() {
	fmt.Println(time.Now()) // 2017-06-07 18:59:35.01959464 +0900 JST
	f := faketime.NewFaketime(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	defer f.Undo()
	f.Do()
	fmt.Println(time.Now()) // 2009-11-10 23:00:00 +0000 UTC
}
```

### NewFaketimeWithTime

with time.Time

```go
package main

import (
	"fmt"
	"github.com/tkuchiki/faketime"
	"time"
)

func main() {
	fmt.Println(time.Now()) // 2017-06-07 18:59:35.01959464 +0900 JST
    t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	f := faketime.NewFaketimeWithTime(t)
	defer f.Undo()
	f.Do()
	fmt.Println(time.Now()) // 2009-11-10 23:00:00 +0000 UTC
}
```
