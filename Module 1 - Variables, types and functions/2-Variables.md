# Variables

## Declaring a variable

To declare a variable, use the `var` statement followed by the variable name and variable type:

```go
var a int // a = 0
var b string // b = ""
var c bool // c = false
```

Variables declared without being initialized are assigned to their type's `zero value`:

* Numeric types: 0
* Boolean types: false
* String types : ""

## Declaring and initializing a variable

To declare and initialize a variable at the same time:

```go
var a int = 123
```

If the type is left out, the variable will take the type of the initializer:

```go
var a = 123 // a is of type int
var b = 3.0 // b is of type float64
var c = true // c is of type bool
```

## Reassigning a variable

Variables can be reassigned, but only to values of their original type:

```go
var a int = 123 // a was initialized to 123
a = 456 // a was reassigned to 456
a = 5.0 // can't do this!
a = "hello" // can't do this!
```

## Multiple variables

To declare a list of variables with the same type:

```go
var a,b,c int 
```

To declare and initialize multiple variables of the same type:

```go
var a,b,c int = 1,2,3 
```

To reassign multiple variables at the same time:

```go
var a,b = 1,2 //a = 1, b = 2
a,b = 3,4 // a = 3, b = 4
a,b = b,a //swaps values of a and b 
```

## Short variable declarations

The `:=` short assignment statement is used to declare a variable with an implicit type. It only works inside functions.

```go
func myFunc(){
    a := 4 // same as `var a = 4`
}
```

Reassignments to variables declared using the short assignment statement still need to be of the original type. 

```go
func myFunc(){
    a := 4 // as is of type int
    a = 5 // ok
    a = 5.0 // can't do this!
    a = "hello" // can't do this!
}
```

## Constants

Constants are declared using the `constant` statement instead of the `var` statement. Constants can not be reassigned.

```go
constant Pi float64 = 3.14159
Pi = 3.14 // can't do this!
```
If the type is left out, the variable will take the type of the initializer:

```go
constant Url = "http://website.com" // Url is of type string
```

Constants can not be declared using the `:=` statement.

