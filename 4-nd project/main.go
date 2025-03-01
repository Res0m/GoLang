package main

import (
	"Golang/password/account"
	"Golang/password/files"
	"Golang/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {

	//1. Создать аккаунт
	//2. Найти аккаунт
	//3. Удалить аккаунт
	//4. Выход
	vault := account.NewVault(files.NewJsonDb("data.json"))
	// vault := account.NewVault(cloud.NewCloudDb("https://aaaa.ru"))

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
		switch choice {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

// 1 -> Создать аккаунт
func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt, ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

// 2 -> Найти аккаунт
func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

// 3 -> Удалить аккаунт
func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для удаления")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Аккаунт удален")
		return
	} else {
		output.PrintError("Не найдено")
	}
}
func menuForAccount() {
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Print("Введите номер задачи: ")
}
