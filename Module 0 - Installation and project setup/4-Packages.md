## Packages

### Importing packages

In our `hello.go` example we import the `fmt` package in order to use its logging methods:

```go

package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}

```

We can import more packages by using the following syntax:

```go

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Printf("Time: %s Rand:%d\n", time.Now(), rand.Intn(5))
}


```

### Exporting packages

To create a package, make a new folder within `/src` and give it a name that you want to refer to it as.

As an example, let's make a package that will contain two methods: Add() and Sub().



Next, inside the folder, make