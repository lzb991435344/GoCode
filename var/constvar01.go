/**
 语言常量
(1)常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
(2)格式：
   const identifier [type] = value
   显示类型定义：const b string = "abc"
   隐式：const b = "abc"
(3)多个相同类型定义
   const c_name1, c_name2 = value1, value2

*/

package main

import "fmt"

func main() {
 	const LENGTH int = 10
 	const WIDTH int = 5
 	var area int
 	const a, b, c = 1, false, "str" //多重赋值

    area = LENGTH*WIDTH
    fmt.Println("面积为:%d\n", area)
    print(a, b, c)
 } 
