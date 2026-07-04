package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	createAccount()
}

func createAccount() {
	// files.ReadFile()
	login := promptData("Введіть логін")
	password := promptData("Введіть пароль")
	url := promptData("Введіть url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Невірний формат URL або Login")
		return
	}
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Помилка при конвертації в JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

// func outputPassword(acc *account) {
// 	fmt.Println(acc)
// 	fmt.Println(acc.login, acc.password, acc.url)
// }

// func generatePassword(n int) string {
// 	res := make([]rune, n)
// 	for i := range res {
// 		res[i] = letterRunes[rand.IntN(len(letterRunes))]
// 	}
// 	return string(res)
// }
