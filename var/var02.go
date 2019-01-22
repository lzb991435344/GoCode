/**
 1.多变量声明
 (1)声明多个类型形同的变量，非全局变量
   var vname1,vname2,vname3 type
      vname1,vname2,vname3  = v1, v2, v3
    
 (2)不需要显示声明类型，自动推断
   var vname1,vname2,vname3  = v1, v2, v3
 (3)//出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误
   vname1, vname2, vname3 := v1, v2, v3 
 (4)因式分解关键字的写法用于声明全局变量
 var (
    vname1 v_type1
    vname2 v_type2
 )
*/

 package main


 var x,y int

//声明全局变量
 var (
 	a int
 	b bool
 	)
 var c, d int = 1, 2
 var e, f = 123,"hello"

//不带格式声明的只能在函数体中出现
//g,h := 123,"hello"

 func main() {
 	g,h := 123,"hello"
 	println(x, y, a, b, c, d, e, f, g, h)
 }
