/**
1.定义切片,不需要说明长度
 var identifier []type
2.make创建切片
  var slice []type = make([]type, len)
  or: slice1 := make([]type, len)
3.指定容量，capacity为可选参数
  make([]T, length,capacity)
4.初始化切片  
 (1) s :=[] int {1,2,3} ,len = 3
 (2) s := arr[startIndex:endIndex],数组的引用
     范围是startIndex到endIndex - 1
 (3) s := arr[startIndex:],缺省endIndex时将表示
 一直到arr的最后一个元素
 (4) s := arr[:endIndex],缺省startIndex时将表示
 从arr的第一个元素开始
 (5)s1 := s[startIndex:endIndex],切片之间初始化
 (6)s := make([]int, len ,cap)
 通过内置函数make()初始化切片s,[]int 标识为其元
 素类型为int的切片
 */
 
 /**len(),cap()函数
 切片是可索引的，并且可以由 len() 方法获取长度。
 切片提供了计算容量的方法 cap() 可以测量切片最长
 可以达到多少
 */
 
 package main
 
 import "fmt"
 
 func printSlice(x []int){
    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
 }
 
 func main(){
   var numbers = make([]int, 3, 5)
   printSlice(numbers)
 }
 //len=3 cap=5 slice=[0 0 0]
 //一个切片在未初始化之前默认为 nil，
 //长度为 0
 
 
 
 