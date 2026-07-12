package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	output.PrintError(1)
	output.PrintError("sd")
	// 1. Створення акаунта
	// 2. Знайти акаунт
	// 3. Видалити акаунт
	// 4. Вихід
	fmt.Println("___Менеджер паролів___")
	vault := account.NewVault(files.NewJsonDb("data.json"))
	//  vault := account.NewVault(cloud.NewCloudDb("https://a.com"))
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func findAccount(vault *account.VaultWithDb) {
	// URL
	// Пошук
	// Вивід
	url := promptData("Введіть url для пошуку")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Акаунтів не знайдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	// URL
	// Метод видалення у Vault
	// Вивід видалено чи ні
	url := promptData("Введіть url для видалення")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Акаунт успішно видалено")
	} else {
		color.Red("Не вдалося видалити акаунт")
	}
}

func createAccount(vault *account.VaultWithDb) {
	// files.ReadFile()
	login := promptData("Введіть логін")
	password := promptData("Введіть пароль")
	url := promptData("Введіть url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Невірний формат URL або Login")
		return
	}
	vault = account.NewVault(files.NewJsonDb("data.json"))
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
