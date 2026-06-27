package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAd time.Time
	updatedAd time.Time
	account
}

func (acc *accountWithTimeStamp) outputPassword() {
	fmt.Println(*acc)
	// fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *accountWithTimeStamp) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")

	}
	newAcc := &accountWithTimeStamp{
		createdAd: time.Now(),
		updatedAd: time.Now(),
		account: account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
