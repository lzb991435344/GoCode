
//append() and copy()函数
//增加切片的容量，把切片的内容拷贝过来

package main

import "fmt"
 
 func printSlice(x []int) {
 	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
 }
 func main() {
  var numbers []int 
  
  //允许追加空切片
  numbers = append(numbers, 0)
  printSlice(numbers)

  //添加一个元素
  numbers = append(numbers, 1)
  printSlice(numbers)
  

  //向切片添加多个元素
  numbers = append(numbers, 2, 3, 4)
  printSlice(numbers)

  //容量增加两倍
  numbers1 := make([]int, len(numbers), (cap(numbers) )* 2)
 
 //拷贝内容
  copy(numbers1, numbers)
  printSlice(numbers1)
}