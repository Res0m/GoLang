package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)



const IMTPower = 2
func main() {

    //-------------------------------------------
	//Циклы: 
	// for i := 0; i < 10; i++ {
	// 	if i == 5{
	// 		continue;
	// 	}
	// 	fmt.Printf("%d\n", i)
	// 	i++;
	// }
	//-------------------------------------------



	//--------------------------------
	// Переменные: 
	// var userHeight = 1.85
	// var userKg float64 = 82
	// userHeight1 := 1.0
	// const IMTpower = 2
	// var c float64
	// c = 200
	
	// var a, b float64 = 2.0, 10
	// g, b := 50, 15

	//var stroka string = "legenda"
	//--------------------------------

    //--------------------------------
	// Практика:
	// const IMTpower = 2
	// var userHeight float64
	// var userKg float64
	// fmt.Println("___Калькулятор индекса массы тела___")
	// fmt.Print("Введите свой рост в сантиметрах: ")
	// fmt.Scan(&userHeight)
	// fmt.Print("Введите свой вес: ")
	// fmt.Scan(&userKg)
	// var IMT = userKg / math.Pow(userHeight/100, IMTpower)
	// // fmt.Printf("Вам индекс массы тела: %v", IMT)
	// result := fmt.Sprintf("Вам индекс массы тела: %.0f", IMT)
	// fmt.Print(result)
	//------------------------------------------------

	// Функции: 
	// a := 10
	// b := 5
	// c := adder(a, b)
	// fmt.Println(c)
	// c = subtractor(a, b)
	// fmt.Println(c)
	// c = multiplier(a, b)
	// fmt.Println(c)
	// c = divider(a, b)
	// fmt.Println(c)
	//------------------------------------------------


	var x float64
	reader := bufio.NewReader(os.Stdin)

	for {
		menu()
		_, err := fmt.Scan(&x)
		if err != nil{
			fmt.Println("Ошибка ввода. Пожалуйста введите число (1 или 2)")
			reader.ReadString('\n') // очищаем буфер
			continue
		}


		if x == 1{
			userKg, userHeight := getUserInput()
			IMT, error := calculateIMT(userKg, userHeight)
			if error != nil{
				// fmt.Println(error)
				// continue
				panic("Произошла ошибка")
			}
			outputResult((IMT))
			switch{
			case IMT < 16:
				fmt.Println("У вас сильный дефицит массы тела")
			case IMT < 18.5:
				fmt.Println("У вас дефицит массы тела")
			case IMT < 25:
				fmt.Println("У вас нормальный вес")
			case IMT < 30:
				fmt.Println("У вас избыточный вес")
			default:
				fmt.Println("У вас степень ожирения")
			}
		} else if x == 2{
			return
	}
	}
}
	// -----------------------------------------------------------
	// if, else if, else : 
	// if IMT < 16 {
	// 	fmt.Println("У вас сильный дефицит массы тела")
	// } else if IMT < 18.5{
	// 	fmt.Println("У вас дефицит массы тела")
	// } else if IMT < 25{
	// 	fmt.Println("У вас нормальный вес")
	// } else if IMT < 30{
	// 	fmt.Println("У вас избыточный вес")
	// } else{
	// 	fmt.Println("У вас степень ожирения")
	// }
	// -----------------------------------------------------------


	//------------------------------------------------------------
	//Switch конструкция:
	// switch{
	// case IMT < 16:
	// 	fmt.Println("У вас сильный дефицит массы тела")
	// case IMT < 18.5:
	// 	fmt.Println("У вас дефицит массы тела")
	// case IMT < 25:
	// 	fmt.Println("У вас нормальный вес")
	// case IMT < 30:
	// 	fmt.Println("У вас избыточный вес")
	// default:
	// 	fmt.Println("У вас степень ожирения")
	// }
	//------------------------------------------------------------







//------------------------------------------
//Функции:
// func adder(a, b int) int{
// 	return a + b
// }


// func subtractor(a, b int) int{
// 	return a - b 
// }

// func multiplier(a, b int) int{
// 	return a * b
// }

// func divider(a, b int) int{
// 	return a / b
// }
//------------------------------------------




func outputResult(imt float64){
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
}

func calculateIMT(kg, height float64) (float64, error){
	if kg <= 0 || height <= 0{
		return 0, errors.New("Вес или рост не могут быть меньше или равны 0")
	}
	IMT := kg / math.Pow(height/100, IMTPower) 
	return IMT, nil
}


func getUserInput() (float64, float64){
	var height float64
	fmt.Printf("Введите свой рост в сантиметрах: ")
	fmt.Scan(&height)
	var kg float64
	fmt.Printf("Введите свой вес в кг: ")
	fmt.Scan(&kg)
	return kg, height
}


func menu() {
	fmt.Printf("___________Меню___________\n")
	fmt.Printf("Выберите вариант: \n")
	fmt.Printf("1. Посчитать индекс массы тела \n")
	fmt.Printf("2. Завершить работу \n")
	fmt.Print("Напишите номер задачи: ")
}