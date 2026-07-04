package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() (*Vault, error) {
	return &Vault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}, nil
}

func (vault *Vault) addAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
}

func (acc *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}
