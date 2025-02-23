package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/fatih/color"
)


var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!@#$%^&*()_+")


type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
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
// Если с маленькой буквы, то метод приватный
// Если с большой буквы, то публичный
func NewAccountWithTimeStamp(log, passw, urlString string) (*AccountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("URL is not valid")
	}
	if log == "" {
		return nil, errors.New("login is empty")
	}
	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login:    log,
			password: passw,
			url:      urlString,
		},
	}
	if passw == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil

}

// Вывод данных
/* метод для структуры account */
func  (acc *AccountWithTimeStamp) OutputPassword() {
	color.Cyan(acc.login, acc.password, acc.url)
	fmt.Println(acc.login, acc.password, acc.url)
}

// Генерация пароля
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}
