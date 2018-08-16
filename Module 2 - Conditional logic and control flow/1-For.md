# For

## For

The for loop in Go has the following syntax:

```go
for i := 0; i < 10; i++ {
	fmt.println(i)
}
```

The for loop has 3 parts

## For as while

There is no while loop in Go. However, you can make a for loop into a while loop by omitting the pre and post condition.

```go
i := 0
for i < 10 {
    fmt.println(i)
    i++
}
```

## Forever loop

A for loop can loop forever if you omit all of its conditions.

```go
for {
    //break somewhere?
}
```

You can still break, continue, or goto out of the infinite for loop



