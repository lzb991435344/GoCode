
/**
 1.变量声明
  (1)该语言变量由字母，数字和下划线组成，首个字符不能是数字
  (2)指定变量类型，声明后不赋值，则使用默认值
    var v_name v_value
    v_name = value
  (3)根据值自行判断变量类型
    var v_name = value
  (4)省略var，:=左侧的变量不应该是已经声明过的，否则编译出错
   var_name := value

   var a int = 10
   var b = 10
   c := 10
*/

 package main 

 var a = "i am blake"
 var b string  = "www.blake.com"

 var c bool 

 func main() {
 	print(a, b, c)
 }