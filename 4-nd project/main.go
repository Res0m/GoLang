package main

import (
	"fmt"
	"Golang/password/account"
	"Golang/password/files"
)


func main() {
	createAccount()

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
