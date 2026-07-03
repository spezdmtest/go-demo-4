package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	login     string `json:"login" xml:"test"`
	password  string
	url       string
	createdAd time.Time
	updatedAd time.Time
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")

	}
	newAcc := &Account{
		createdAd: time.Now(),
		updatedAd: time.Now(),
		url:       urlString,
		login:     login,
		password:  password,
	}
	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	fmt.Println(string(field.Tag))
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
