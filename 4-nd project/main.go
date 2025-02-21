package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)


type account struct {
	login string
	password string
	url string
}



type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}
// 1. Если логина нет, то ошибка 
// 2. Если пароля нет, то генерируем
// func newAccount(log, passw, urlString string) (*account, error){
// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil{
// 		return nil, errors.New("URL is not valid")
// 	}
// 	if log == ""{
// 		return nil, errors.New("Login is empty")
// 	}
// 	newAcc := &account{
// 		login: log,
// 		password: passw,
// 		url: urlString,	
// 	}
// 	if passw == ""{
// 		newAcc.generatePassword(12)
// 	}
// 	return newAcc, nil

// }

func newAccountWithTimeStamp(log, passw, urlString string) (*accountWithTimeStamp, error){
	_, err := url.ParseRequestURI(urlString)
	if err != nil{
		return nil, errors.New("URL is not valid")
	}
	if log == ""{
		return nil, errors.New("Login is empty")
	}
	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login: log,	
			password: passw,
			url: urlString,
	},
}
	if passw == ""{
		newAcc.generatePassword(12)
	}
	return newAcc, nil

}

// Вывод данных
func /* метод для структуры account */ (acc *accountWithTimeStamp) outputPassword(){
	fmt.Println(acc.login, acc.password, acc.url, acc.createdAt, acc.updatedAt)
}

// Генерация пароля
func (acc *account) generatePassword(n int){
	res := make([]rune, n)
	for i := range res{
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!@#$%^&*()_+")



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

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil{
		fmt.Println(err)
		return
	}

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


	myAccount.outputPassword()


}


func promptData(prompt string) string { 
	fmt.Print(prompt, ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

