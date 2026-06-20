package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	str := "Привіт!)"
	for _, ch := range str {
		fmt.Println(ch, string(ch))
	}

	str1 := []rune("Привіт!)")
	for _, ch := range str1 {
		fmt.Println(ch, string(ch))
	}

	str2 := []int32("Привіт!)")
	for _, ch := range string(str2) {
		fmt.Println(ch, string(ch))
	}

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

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}
