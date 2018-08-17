# Functions

## Function syntax

Functions are declared using the `func` statement and have the following syntax:

```go
func myFunction(){
   
}
```

## Function parameters

Functions can have any number of parameters. The parameter type must come after the parameter name. 

```go
func myFunction(param1 int, param2 string) {

}
```

If the function has multiple parameters with the same return type, you can omit the type from all but the last parameter.

```go

func myFunction(param1, param2 int){

}

```
## Function return types

Function return types are specified after the last parenthesis and before the first bracket. 

```go
func add(x, y int) int {
    return x + y
}

```

## Multiple return types

Functions can have multiple return types. To do so use the following syntax:

```go
func myFunction() (int, string){
    return 5, "five"
}
```

You can get the multiple results by assigning them to multiple variables:

```go
var a,b = myFunction() // a = 5, b = "five"
```

If you want to omit one of the return values, use a _ to replace one of the variables. Having less variables than the number of return values doesn't work.

```go
var a,_ = myFunction() //a = 5
var _,b = myFunction() //b = "five"
var a = myFunction() //doesn't work
```




## Named return types

Functions can have named return types. 

Named return types are essentially the same as having variables declared at the top of the function.

Like with parameters, if the return types have the same type, you can omit the type from all but the last.

If the function returns without arguments, it will return the named return values instead. This is known as a naked return.

```go
func addSub(x,y int) (add, sub int){
    add = x + y
    sub = x - y
    return //returns add, sub
}
```
You should only use named returns in short functions. In longer functions they make readability more difficult. 


