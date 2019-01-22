/**
切片截取
可以通过设置下限及上限来设置截取切片 
[lower-bound:upper-bound]
*/


package main

import "fmt"

func printSlice(x []int){
  fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}


func main(){
  //创建，打印原始切片
  numbers :=[]int{0,1,2,3,4,5,6,7,8}
  printSlice(numbers)
  
  fmt.Println("numbers==", numbers)
  
  //前闭后开区间
  fmt.Println("numbers[1:4]==", numbers[1:4])

  //默认下限为0
  fmt.Println("numbers[:3]", numbers[:3])
  
  //默认上限为len(s)

  fmt.Println("numbers[4:]==", numbers[4:])

  numbers1 := make([]int , 0 ,5)
  printSlice(numbers1)

  //打印切片从0-2的索引，[0,2)
  numbers2 := numbers[:2]
  printSlice(numbers2)

  number3 :=numbers[2:5]
  printSlice(number3)
 }