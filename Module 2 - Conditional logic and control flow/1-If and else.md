# if and else

## if and else
`if` statements have the following syntax:

```go
if x < 0 {
    fmt.Println("negative")
}else if x > 0 {
    fmt.Println("positive")
}else {
    fmt.Println("zero")
}
```

Curly braces are required and must always be used. However, there are no parenthesis, unlike other languages.

## Short statements

Like `for` loops, `if` statements can start with a short statement to execute before the condition. 

```go

if x := getStatus(); x == true {
    fmt.Println("Active")
}else {
    fmt.Println("Not active")
}

```

Short statements are useful when a function returns multiple values with one value being a status boolean value.

```go
if x, ok := dict["foo"]; ok {
    fmt.println(x)
}
```

The scope of the short statement is only within the if-else block.

