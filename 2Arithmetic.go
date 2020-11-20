package main 

import "fmt"

func main () {

x,y := 3,9

fmt.Println(x+y)
fmt.Println(x-y)
fmt.Println(x*y)
fmt.Println(x/y)

isbool := true
hate := false 
fmt.Println(isbool && hate)
fmt.Println(isbool || hate)
fmt.Println(!isbool)
}

OUT PUT

12
-6
27
0
false
true
false
