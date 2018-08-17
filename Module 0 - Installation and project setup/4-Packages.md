# Packages

## Importing packages

In our `hello.go` example we imported the `fmt` package in order to use its logging methods:

```go
package main

import "fmt"

func main() {
    fmt.Println("hello, world")
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

## Exporting packages

It is useful to create your own packages so you can reuse commonly used code. It also helps to keep your application organized.

As an example, let's make a package named `mylib` that contains a few helper functions.

### Creating a package directory
Create a directory named `mylib` within `/src`:

```
workspace
  /src
    /mylib
  /pkg
  /bin
```

Each package should be in a separate directory with the same name. You can have multiple files in a package directory as long as they have the same package name.

### Adding files to a package and exporting functions
Next, create a file named `mylib.go` within `/mylib`. Copy the following code into `mylib.go`:

```go
package mylib

import "fmt"

func Hello(){
   fmt.Println("hello")
}

func World(){
	fmt.Println("world")
} 
```
The `package` statement determines what package a Go file will be a part of. In this example, it is part of the `mylib` package.

By default, all function names that start with a capital letter are exported to the package. 

### Importing a package into your main package

Next, create a file named `main.go` within `/src`. Copy the following code into `main.go`:

```go
package main

import "mylib"

func main(){
   mylib.Hello()
   myLib.World()
}
```

The `main.go` file will contain your `main` package. The `mylib` package will be imported using the `import` statement.

We can access the functions we built in our `mylib` package by using `mylib.Hello()` and `mylib.World()`.

### Building your main package

Your final folder structure should look like this:
```
workspace
  /src
    /mylib
      mylib.go
    main.go
  /pkg
  /bin
```

Navigate to your `./src` directory and build and run your application using `go run main.go`. You should see `Hello` and `World` printed to the console on separate lines.

