# Module 3 Knowledge Check

## Question 1

What does the following code print out?

```go
m := map[string] string{"apple": 3, "banana":4, "carrot": 2} 

for key, val := range array{
    fmt.Println(key)
}
```
A. 
```go
apple
banana
carrot

```

B.
```
3
4
2
```
C.
```
apple 3
banana 4
carrot 2
```
D.
```
0
1
2
```

## Question 2
True or false? When an array is passed into a function it is passed by reference.

A. True

B. False

## Question 3:

What should the slice start and end index be so that the code prints out [8 10 12 14 16]?

```go

array := [100]int{2,4,6,8,10,12,14,16}
s := array[?:?]
fmt.Println(s)
```

A. [3:7]

B. [3:]

C. [4:8]

D. [4:7]

## Question 4:

What will the following code print?

```go
    m := map[string] string{"bob": "barns", "joe": "johnson", "sally": "sue"} 
    val := m["barns"] 
    fmt.println(val) 
```
A. It will cause a compile error

B. ""

C. "bob"

D. `nil`
