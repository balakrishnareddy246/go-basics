package main 

import "fmt"

func main () {
    
    var name string = "Balakrishna reddy"
    const pi float64 =3.141615139

    win := true
    x := 6
    fmt.Println("%.3f \n",pi)
    fmt.Println("%T \n", name)
    fmt.Println("%t \n", win)
    fmt.Println("%d \n", x)
    fmt.Println("%b \n", 26)
    fmt.Println("%c \n", 53)
    fmt.Println("%x \n", 42)
    fmt.Println("%e\n", pi)
}


OUT PUT 
$go run main.go
%.3f 
 3.141615139
%T 
 Balakrishna reddy
%t 
 true
%d 
 6
%b 
 26
%c 
 53
%x 
 42
%e
 3.141615139
