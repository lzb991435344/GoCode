/**
Map 是一种无序的键值对的集合。Map 最重要的一点是
通过 key 来快速检索数据，key 类似于索引，指向数据
的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样
迭代它。不过，Map 是无序的，我们无法决定它的返回
顺序，这是因为 Map 是使用 hash 表来实现的。

定义map
(1)var map_variable map[key_data_type]value_data_type
(2)map_variable := make(map[key_data_type]value_data_type)
 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来
存放键值对
*/
package main
import "fmt"

func main(){
   //define map
   var countryCapitalMap map[string]string
   countryCapitalMap = make(map[string]string)

   //insert values
   countryCapitalMap ["1"] = "1"
   countryCapitalMap ["2"] = "2"
   
   //println
   for country := range countryCapitalMap{
     fmt.Println(country,"number is",countryCapitalMap[country])
   }
   
   captical, ok :=countryCapitalMap["2"]
   if(ok){
     fmt.Println("number is",captical)
   }else{
     fmt.Println("not exist!\n")
   }
}




