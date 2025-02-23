package main

import (
	"fmt"
	"Golang/password/account"
	"Golang/password/files"
)


func main() {
	//1. Создать аккаунт
	//2. Найти аккаунт
	//3. Удалить аккаунт
	//4. Выход


	fmt.Println("Программа для хранения паролей")
Menu: // label -> какой-то лейбл (часть кода)
	for {
		menuForAccount()
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch choice{
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
	// -----------------------------------------------------------------------------------------
	// str :=[]int32(rune)("Привет!)")
	// for _, ch := range string(str){
	// 	fmt.Println(ch, string(ch))
	// }
	// -----------------------------------------------------------------------------------------
	

	
	// -----------------------------------------------------------------------------------------
	// //
	// // account2 := account{
	// // 	login: login,
	// // 	password: password,
	// // }
	// // account3 := account{}
	// outputPassword(&myAccount)
	// myAccount.generatePassword(8)
	// -----------------------------------------------------------------------------------------

}

//1 -> Создать аккаунт
func createAccount(){
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := myAccount.ToBytes()
	if err != nil {	
		fmt.Println(err)
		return
	}
	files.WriteFile(file, "data.json")
}


func promptData(prompt string) string {
	fmt.Print(prompt, ": ")
	var res string
	fmt.Scanln(&res)
	return res
}


//2 -> Найти аккаунт
func findAccount(){

}
//3 -> Удалить аккаунт
func deleteAccount(){

}
func menuForAccount(){
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
}
