package main 

import (
	"fmt"
)


func main(){
	transactions := []int{0, 20, 35} // [] - длина + тип (например float) + {}содержимое масива (например: {1, 2, 4})
	temp := transactions

	// partial := transactions[1:] // [1:n] - срез масива (от 1 до n)	
	transactions = append(transactions, 100)
	fmt.Println(temp)
	fmt.Println(transactions)

	





}