package main

import "fmt"

//if语句
func main() {
   /* 定义局部变量 */
   var a int = 10
 
   /* 使用 if 语句判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)
}


//if ... else
func main() {
   /* 局部变量定义 */
   var a int = 100;
 
   /* 判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" );
   } else {
       /* 如果条件为 false 则执行以下语句 */
       fmt.Printf("a 不小于 20\n" );
   }
   fmt.Printf("a 的值为 : %d\n", a);
}


//if嵌套
func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200
 
   /* 判断条件 */
   if a == 100 {
       /* if 条件语句为 true 执行 */
       if b == 200 {
          /* if 条件语句为 true 执行 */
          fmt.Printf("a 的值为 100 ， b 的值为 200\n" );
       }
   }
   fmt.Printf("a 值为 : %d\n", a );
   fmt.Printf("b 值为 : %d\n", b );
}

//switch
func main() {
   /* 定义局部变量 */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }

   switch {
      case grade == "A" :
         fmt.Printf("优秀!\n" )     
      case grade == "B", grade == "C" :
         fmt.Printf("良好\n" )      
      case grade == "D" :
         fmt.Printf("及格\n" )      
      case grade == "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" );
   }
   fmt.Printf("你的等级是 %s\n", grade );      
}

//type switch
/**switch 语句还可以被用于 type-switch 来判断某个 interface 变量中
实际存储的变量类型。
switch x.(type){
    case type:
       statement(s);      
    case type:
       statement(s); 
    // 你可以定义任意个数的case 
    default: // 可选 
       statement(s);
}
*/
func main() {
   var x interface{}
     
   switch i := x.(type) {
      case nil:      
         fmt.Printf(" x 的类型 :%T",i)                
      case int:      
         fmt.Printf("x 是 int 型")                       
      case float64:
         fmt.Printf("x 是 float64 型")           
      case func(int) float64:
         fmt.Printf("x 是 func(int) 型")                      
      case bool, string:
         fmt.Printf("x 是 bool 或 string 型" )       
      default:
         fmt.Printf("未知型")     
   }   
}


//select 
/**
select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s); 
    // 你可以定义任意数量的 case 
    default : // 可选 
       statement(s);
}

(1)每个case都必须是一个通信
(2)所有channel表达式都会被求值
(3)所有被发送的表达式都会被求值
(4)如果任意某个通信可以进行，它就执行；其他被忽略。
(5)如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。 
否则：
 a.如果有default子句，则执行该语句。
 b.如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值
进行求值。
*/

func main() {
   var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    
}