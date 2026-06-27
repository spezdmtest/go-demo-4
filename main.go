package main

import (
	"demo/password/account"
	"fmt"
)

func main() {
	login := promptData("Введіть логін")
	password := promptData("Введіть пароль")
	url := promptData("Введіть url")
	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Невірний формат URL або Login")
		return
	}
	myAccount.OutputPassword()
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
