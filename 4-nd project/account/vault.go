package account

import (
	"Golang/password/files"
	"Golang/password/output"
	"encoding/json"
	"strings"
	"time"
)

// Разделение интерфейсов:
type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface { // Просто вписываем интерфейсы в один общий, если требуется
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.json")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := checker(account, str)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	db := files.NewJsonDb("data.json")
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		output.PrintError("Не удалось преобразовать")
	}
	db.Write(data)
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		output.PrintError("Не удалось преобразовать")
	}
	db := files.NewJsonDb("data.json")
	db.Write(data)
	return isDeleted
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
