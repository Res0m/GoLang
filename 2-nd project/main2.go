package main 

import (
	"fmt"
)
// В цикле спрашиваем ввод транзакции: -10, 10, 40.5
// Добавлять каждую в массив транзакций
// Вывести массив
// Вывести сумму всех транзакций

func main(){
	// tr1 := []int{1, 2, 3}
	// tr2 := []int{4, 5, 6}
	// tr1 = append(tr1, tr2...) //  в метод append() можно передавать любое количество значений
	// fmt.Println(tr1)		  // ... -> unpack, т.е распаковывает слай и берет из него значения



	// Цикл по массиву (слайсу) структура:
	// for index, value := range + massive, range -  обозначает что мы берем все значения массива (элементы)
	// index - индекс, value - значение
	// for _, value := range + massive, _ - если не нужен индекс, только значение
	// for index, _ := range + massive, _ - если не нужно значение, только индекс
	// for index := range + massive - если нужен только индекс

	// for index, value := range tr1{
	// 	fmt.Println(index, value)
	// }
	// for _, value := range tr1{
	// 	fmt.Println(value)
	// }
	// for index, _ := range tr1{ // --> Также если хотим только индекс, только с _
	// 	fmt.Println(index) 
	// }

	// for index := range tr1{ //--> Если хотим чисто индекс
	// 	fmt.Println(index)
	// }


	// transactions := []int{0, 20, 35} // [] - длина + тип (например float) + {}содержимое масива (например: {1, 2, 4})
	// temp := transactions

	// // partial := transactions[1:] // [1:n] - срез масива (от 1 до n)	
	// transactions = append(transactions, 100)
	// fmt.Println(temp)
	// fmt.Println(transactions)
	fmt.Print("Введите количество транзакций: ")
	var n int
	fmt.Scan(&n)
	transactions := []int{}
	for i := 0; i < n; i++ {
		fmt.Print("Введите транзакцию: ")
		var transaction int
		fmt.Scan(&transaction)
		transactions = append(transactions, transaction)
	}
	var sum int
	for _, value := range transactions{
		sum += value
	}
	fmt.Println(transactions)
	fmt.Println("Сумма всех транзакций: ", sum)
}