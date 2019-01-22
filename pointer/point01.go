
/**
1.指针
一个指针变量指向了一个值的内存地址。
类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

var var_name *var-type
var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：

var ip *int        // 指向整型
var fp *float32    // 指向浮点型 
本例中这是一个指向 int 和 float32 的指针。
2.如何使用指针
指针使用流程：

定义指针变量。
为指针变量赋值。
访问指针变量中指向地址的值。
在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
*/

package main 

import "fmt"

func  main() {
	var a int = 10
	var ip *int
	var ptr *int

	ip = &a
	fmt.Println("变量的地址是：%x\n",&a)

	fmt.Println("ip指针变量的地址是：%x\n",ip)
	fmt.Println("ip指针变量的值是：%x\n",*ip)
	fmt.Println("ptr指针变量的值是：%x\n",ptr)
}