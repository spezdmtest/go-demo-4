package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(*acc)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// 1. якщо лоніна немає, то помилка
// 2. якщо немає пароля, то генеруємо
func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("INVALID_")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")

	}
	newAcc := &account{
		url:      urlString,
		login:    login,
		password: password,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {
	//fmt.Println(generatePassword(12))
	login := promptData("Введіть логін")
	password := promptData("Введіть пароль")
	url := promptData("Введіть url")

	// account1 := account{
	// 	login,
	// 	"",
	// 	url,
	// }

	// account1 := account{}

	myAccount, err := newAccount(login, password, url)
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
