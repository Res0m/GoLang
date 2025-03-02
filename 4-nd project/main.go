package main

import (
	"Golang/password/account"
	"Golang/password/files"
	"Golang/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}


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
		choice := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выхода",
			"Выберите выриант",
		})
		menuFunc := menu[choice]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		// switch choice {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
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
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат")
		return
	}
	vault.AddAccount(*myAccount)
}

// Функция принимала слайс любого типа
// Выводит строкой каждый элемент, а последний - выводит через Printf(), добавляя:
func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

// 2 -> Найти аккаунт
func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool{ // анонимная функция
		return strings.Contains(acc.Url, str)
	})
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}


// func checkUrl(acc account.Account, str string) bool {
// 	return strings.Contains(acc.Url, str)
// }

// 3 -> Удалить аккаунт
func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для удаления"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Аккаунт удален")
		return
	} else {
		output.PrintError("Не найдено")
	}
}
