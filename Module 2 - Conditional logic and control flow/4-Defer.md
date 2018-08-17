# Defer

## Defer

The defer statement delays the execution of a function until its surrounding function returns.

The arguments to the deferred function are evaluated immedietely, but the actual function call is not executed until its surrounding function returns.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("later")
	fmt.Println("now")
}
```

## Stacking defers

If multiple functions are deferred, they get added to a stack and get executed LIFO.

```go
package main

import "fmt"

func main() {
	defer fmt.Print("1 ")
    defer fmt.Print("2 ")
    defer fmt.Print("3 ")
    defer fmt.Print("4 ")
}
// outputs 4 3 2 1
```