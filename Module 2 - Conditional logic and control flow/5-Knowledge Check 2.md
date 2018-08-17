# Module 2 Knowledge Check


## Question 1

True or false? `if` statements can omit their curly braces if there is only a single statement for the condition.

A. True

B. False

## Question 2

What will the code output?

```go
package main

import fmt

func main(){
    a ;= 2
    switch a {
        case 2:
            fmt.Print("1 ")
        case 2:
            fmt.Print("2 ")
        case 3:
            fmt.Print("3 ")
        default: 
            fmt.Print("0 ")
    }   
}
```

A. "1 1"

B. "1 "

C. "2 "

D. "0 "



## Question 3

What happens to a `for` loop when you omit the initial statement, conditional expression, and post statement?

A. It will just be skipped over

B. It will not work and will cause a compile error

C. It will iterate through the array

D. It loops forever


## Question 4

When do the arguments of a deferred function get evaluated?

A. Immediately

B. Right before the return value

C. Right before the next defer statement

C. Whenever the deferred function gets executed

