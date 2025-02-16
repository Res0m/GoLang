package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a) // Меняет порядок элементов на обратный
	fmt.Println(a)
}

func reverse(array *[4]int){
	for index, value := range *array{
		(*array)[len(*array) - 1 - index] = value

	}
}