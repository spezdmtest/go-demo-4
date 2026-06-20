package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Введіть логін")
	password := promptData("Введіть пароль")
	url := promptData("Введіть url")

	// account1 := account{
	// 	login,
	// 	"",
	// 	url,
	// }

	// account1 := account{}

	myAccount := account{
		password: password,
		login:    login,
		url:      url,
	}

	outputPassword(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc account) {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}
