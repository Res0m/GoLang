package main

import (
	"fmt"
	"Golang/password/account"
	"Golang/password/files"
)


func main() {
	defer fmt.Println(1) // defer - отложенный вызов функции
	defer fmt.Println(2) // Сначала выполнится 2, потом 1
	// -----------------------------------------------------------------------------------------
	// str :=[]int32(rune)("Привет!)")
	// for _, ch := range string(str){
	// 	fmt.Println(ch, string(ch))
	// }

	// password := generatePassword(8)
	// fmt.Println(password)
	// -----------------------------------------------------------------------------------------
	files.WriteFile("Привет! Я файл", "file.txt")
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.OutputPassword()

	files.ReadFile()
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

func promptData(prompt string) string {
	fmt.Print(prompt, ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
