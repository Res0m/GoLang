package main

import (
	"Golang/password/account"
	"Golang/password/encrypter"
	"Golang/password/files"
	"Golang/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выхода",
	"Выберите выриант",
}

func menuCounter() func() {
	i := 0
	return func() {
		i += 1
		fmt.Println(i)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти env файл")
	}
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
	// vault := account.NewVault(cloud.NewCloudDb("https://aaaa.ru"))

	fmt.Println("Программа для хранения паролей")
Menu: // label -> какой-то лейбл (часть кода)
	for {
		choice := promptData(menuVariants...)
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

// Функция принимала слайс любого типа
// Выводит строкой каждый элемент, а последний - выводит через Printf(), добавляя:
func promptData(prompt ...string) string {
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
func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска: ")
	accounts := vault.FindAccounts(url, func(acc account.Account, url string) bool {
		return strings.Contains(acc.Url, url)
	})
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}
func findAccountByLogin(vault *account.VaultWithDb) {
	url := promptData("Введите Логин для поиска: ")
	accounts := vault.FindAccounts(url, func(acc account.Account, login string) bool {
		return strings.Contains(acc.Login, login)
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

// func checkLog(acc account.Account, str string) bool {
// 	return strings.Contains(acc.Login, str)
// }

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
