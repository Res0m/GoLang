package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!@#$%^&*()_+")

type Account struct {
	Login     string    `json:"login"` // `` - теги
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
func NewAccount(log, passw, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("URL is not valid")
	}
	if log == "" {
		return nil, errors.New("login is empty")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     log,
		Password:  passw,
		Url:       urlString,
	}
	if passw == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil

}

// Вывод данных
/* метод для структуры account */
func (acc *Account) Output() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
}

// Генерация пароля
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}
