package main 

import "fmt"

func main () {
    
    
    x := 12
    changeValue(&x)
    fmt.Println(x)
    
}

func changeValue(x *int) {
    
    *x = 7;
    
}

OUT PUT 
7
