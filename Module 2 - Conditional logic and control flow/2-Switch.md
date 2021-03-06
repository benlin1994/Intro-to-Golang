# Switch

## Switch

Switch statements have the following syntax:

```go
a ;= 2

switch a {
    case 1:
        fmt.Println("case 1")
    case 2:
        fmt.Println("case 2")
    case 3:
        fmt.Println("case 3")
    default: 
        fmt.Println("none")
}
```

The cases are evaluated from top to bottom until one matches the switch condition expression. 

The condition expression doesn't have to be a constant or an int, it can be of any type. The case values don't have to be constants either, they can be evaluated expressions.

If none of the cases match, the default case will be chosen. Only one case will ever be run, so there is no need to use break statements. 

## Cases with multiple values

Switch statement cases can have multiple values separated by commas. If any of them match the switch condition expression, then that case will be chosen.

```go
a ;= 5

switch a {
    case 1,2,3:
        fmt.Println("small")
    case 4,5,6:
        fmt.Println("medium")
    case 7,8,9:
        fmt.Println("large")
    default: 
        fmt.Println("out of range")
}
```

## Short statements

Switch statements can have short statements before the switch condition expression.

```go
switch a:=2; a {
    case 1:
        fmt.Println("case 1")
    case 2:
        fmt.Println("case 2")
    case 3:
        fmt.Println("case 3")
    default: 
        fmt.Println("none")
}
```

## Switch with no condition

Switch statements with no condition are treated the same as `switch true`. This means that the first case that evaluates to `true` will be chosen. This effectively acts like an if-else statement with different syntax.

```go

a := 42 
switch {
    case a > 0:
        fmt.Println("positive")
    case a < 0 :
        fmt.Println("negative")
    default: 
        fmt.Println("zero")
}
```
