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
	outputPassword(login, password, url)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(login, password, url string) {
	fmt.Println(login, password, url)
}
