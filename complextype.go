package main 

import (
  "fmt"
  )
  func main () {
  
  var n complex64 = 1+2i
  fmt.Println(%v, %T\n", real(n), real(n))
  fmt.Printf("%v, %T\n", imag(n), imag(n))
  }
