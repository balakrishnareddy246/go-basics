package main 

import (
         "fmt"
)

func main () {
    n := 1 == 1
    m := 2 == 3
    fmt.Printf("%v, %T\n", n, n)
    fmt.Printf("%v, %Tn", m, m)
}
