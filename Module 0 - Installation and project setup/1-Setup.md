## Setup

### Downloading Go

To download Go, nagivate to https://golang.org/dl/ and download the appropriate installer for your operating system.

The installer should edit your PATH environment variable to allow you to use the `go` command to build your Go files in the command line.

### Running your first Go application

To run your first program, make a new file called `hello.go` in your workspace and copy in the following code:

```go

package main

import "fmt"

func main() {
    fmt.Println("hello, world")
}

```

Next, build and run your code with `go run hello.go`. 

You should see `hello, world` printed to the console.