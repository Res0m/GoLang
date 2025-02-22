package main

import (
	"fmt"
	"Golang/password/account"
)


func main() {

	// -----------------------------------------------------------------------------------------
	// str :=[]int32(rune)("Привет!)")
	// for _, ch := range string(str){
	// 	fmt.Println(ch, string(ch))
	// }

	// password := generatePassword(8)
	// fmt.Println(password)
	// -----------------------------------------------------------------------------------------

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.OutputPassword()

	
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
