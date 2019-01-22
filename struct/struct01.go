/**
 1.定义结构体
  结构体定义需要使用 type 和 struct 语句。
  struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。
  type 语句设定了结构体的名称

type struct_variable_type struct {
   member definition;
   member definition;
   ...
   member definition;
}

一旦定义了结构体类型，它就能用于变量的声明
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

*/

package main

import "fmt"

type Books struct{
   title string
   author string
   subject string
   book_id int
}
 

func main() {
	//创建一个新结构体
	fmt.Println(Books{"Go语言", "www.baidu.com", "Go语言教程", 64985 })
      
    //可以使用 key => value格式
    fmt.Println(Books{title:"Go语言", author:"www.baidu.com", subject:"Go语言教程", book_id:64985})
   
   //忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}