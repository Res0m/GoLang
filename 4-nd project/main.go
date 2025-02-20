package main





import (
	"fmt"
	"math/rand/v2"
)






type account struct {
	login string
	password string
	url string
}

func main() {
	fmt.Println(rand.IntN(10))
	str :=[]int32("Привет!)")
	for _, ch := range string(str){
		fmt.Println(ch, string(ch))
	}






	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount := account{
		login: login,
		password: password,
		url: url,
	}
	// }
	// account2 := account{
	// 	login: login,
	// 	password: password,
	// }
	// account3 := account{}
	outputPassword(&myAccount)
}


func promptData(prompt string) string { 
	fmt.Print(prompt, ": ")
	var res string
	fmt.Scan(&res)
	return res
}


func outputPassword(acc *account){
	fmt.Println(acc.login, acc.password, acc.url)
}


func generatePassword(n int) string{
	
}