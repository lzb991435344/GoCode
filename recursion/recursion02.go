package main

import "fmt"

func fibonacci(n int) int{
  if n < 2{
  return n
  }
  return fibonacci(n - 1) + fibonacci(n - 2)
}

func main(){
 var i int
 for i = 0; i < 10; i++{
    fmt.Println(fibonacci(i))
 }
}