/**
func recursion() {
   recursion() // 函数调用自身
}

func main() {
   recursion()
}
Go 语言支持递归。但我们在使用递归时，
开发者需要设置退出条件，否则递归将陷入
无限循环中
*/

package main

import "fmt"

func Factorial (n uint64)(result uint64){
 if(n > 0){
 result = n * Factorial(n - 1)
   return result
 }
  return 1
}

func main(){
  var i int = 15
  fmt.Println("%d的阶乘是%d的\n",i,Factorial(uint64(i)))
}





