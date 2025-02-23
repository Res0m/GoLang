package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (acc *Vault) ToBytes() ([]byte, error){
	file, err := json.Marshal(acc)
	if err != nil{
		return nil, err
	}
	return file, nil
}


func newVault()Vault{
	return &Vault{
		Accounts: []Account{},
		UpdatedAt: time.Now(),
	}, 
}