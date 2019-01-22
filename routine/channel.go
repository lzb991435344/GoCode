/**
(1)ch <- v    // 把 v 发送到通道 ch
   v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v
(2)ch := make(chan int)
*/

package main

import "fmt"

func sum(s []int, c chan int){
  sum := 0
  for _, v := range s {
    sum += v
  }
  c <- sum
}
func main(){
  s := []int{7, 2, 8, -9, 4, 0}
  c := make(chan int)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:],c)
  x,y := <-c,<-c
  
  fmt.Println(x, y, x+y)
}