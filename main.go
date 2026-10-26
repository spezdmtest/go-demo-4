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
		variant := promptData([]string{
			"1. Створення акаунта",
			"2. Знайти акаунт",
			"3. Видалити акаунт",
			"4. Вихід",
			"Виберіть віріант",
		})

		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func findAccount(vault *account.VaultWithDb) {
	// URL
	// Пошук
	// Вивід
	url := promptData([]string{"Введіть url для пошуку"})
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
	url := promptData([]string{"Введіть url для видалення"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Акаунт успішно видалено")
	} else {
		output.PrintError("Не вдалося видалити акаунт")
	}
}

func createAccount(vault *account.VaultWithDb) {
	// files.ReadFile()
	login := promptData([]string{"Введіть логін"})
	password := promptData([]string{"Введіть пароль"})
	url := promptData([]string{"Введіть url"})
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Невірний формат URL або Login")
		return

	}
	vault = account.NewVault(files.NewJsonDb("data.json"))
	vault.AddAccount(*myAccount)
}

// Функція приймала slice любого типу
// Виводе строкою кожний елемент а останній - Printf додавая :
func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
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
