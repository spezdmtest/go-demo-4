package main

import (
	"demo/password/account"
	"fmt"
)

func main() {
	// 1. Створення акаунта
	// 2. Знайти акаунт
	// 3. Видалити акаунт
	// 4. Вихід
	fmt.Println("___Менеджер паролів___")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Виберіть віріант: ")
	fmt.Println("1. Створення акаунта")
	fmt.Println("2. Знайти акаунт")
	fmt.Println("3. Видалити акаунт")
	fmt.Println("4. Вихід")
	fmt.Scanln(&variant)
	return variant
}

func findAccount() {
}

func deleteAccount() {
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
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
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
