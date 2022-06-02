# libganache

Install `ganache` (`yarn global add ganache`) and then use `libganache` like this:

```go
package main

import (
	"fmt"
	"github.com/xrash/libganache"
)

func main() {
	g, err := libganache.RunGanache(nil)
	if err != nil {
		panic(err)
	}

	accounts, err := g.Accounts()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(accounts); i++ {
		fmt.Printf("accounts[%d] = %s\n", i, accounts[i].PublicKey)
	}
}
```

