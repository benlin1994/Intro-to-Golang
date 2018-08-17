# For

## For

The `for` loop in Go has the following syntax:

```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}
```

There are three parts to a basic for loop:
* initial condition - executed before the first iteration
* conditional expression - evaluated before every iteration
* post statement - executed after every iteration

## For as while

There is no `while` loop in Go. However, you can make a `for` loop into a `while` loop by omitting the initial and post condition.

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

## Forever loop

A `for` loop can loop forever if you omit all of its conditions.

```go
for {

}
```

You can still `break` out of an infinite `for` loop.



