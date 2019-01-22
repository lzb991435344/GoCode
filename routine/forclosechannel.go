/**
Go 通过 range 关键字来实现遍历读取道的数据，
类似于与数组或切片。格式如下：

v, ok := <-ch
如果通道接收不到数据后 ok 就为 false，这时通道
就可以使用 close() 函数来关闭。
*/
package main

import("fmt")

func fibonacci(n int, c chan int){
  x,y := 0, 1
  for i := 0; i < n ; i++{
    c <- x
	x, y = y, x+y
  }
  close(c)
}

func main(){
  c:=make(chan int, 10)
  go fibonacci(cap(c), c)
  for i := range c{
   fmt.Println(i)
  }
}
