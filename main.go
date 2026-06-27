package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	//fmt.Println(generatePassword(12))
	login := promptData("Введіть логін")
	password := promptData("Введіть пароль")
	url := promptData("Введіть url")
	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Невірний формат URL або Login")
		return
	}
	// myAccount.generatePassword(12)
	myAccount.outputPassword()
	// fmt.Println(myAccount)
	// outputPassword(&myAccount)

	// str := "Привіт!)"
	// for _, ch := range str {
	// 	fmt.Println(ch, string(ch))
	// }

	// str1 := []rune("Привіт!)")
	// for _, ch := range str1 {
	// 	fmt.Println(ch, string(ch))
	// }

	// str2 := []int32("Привіт!)")
	// for _, ch := range string(str2) {
	// 	fmt.Println(ch, string(ch))
	// }
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}

func generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(res)
}
