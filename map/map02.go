/**
 map：
 delete() 函数用于删除集合的元素, 参数为 map 
和其对应的 ke
*/

package main

import "fmt"

func main(){
    
	countryCapitalMap := map[string]string{"1":"1", "2":"2", "3":"3", "4":"4", "5":"5"}
	
	//print map
	fmt.Println("begin Map")
    for country := range countryCapitalMap{
	fmt.Println(country, "number is", countryCapitalMap[country])
	}
	
	fmt.Println("delete Map")
	delete(countryCapitalMap, "2")
	
	fmt.Println("after delete Map")
    for country := range countryCapitalMap{
	fmt.Println(country, "number is", countryCapitalMap[country])
	}
}