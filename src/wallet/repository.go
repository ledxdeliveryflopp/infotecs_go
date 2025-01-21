package wallet

import (
	"errors"
	"fmt"
	"infotecs_go/src/settings"
)

func GetWalletByNumberRepository(number string) (Wallet, error) {
	db := settings.ConnectToBD()
	queryStr := fmt.Sprintf("SELECT number, balance FROM wallet WHERE number = '%s'", number)
	row := db.QueryRow(queryStr)
	var walletInfo Wallet
	err := row.Scan(&walletInfo.Number, &walletInfo.Balance)
	if err != nil {
		return walletInfo, err
	}
	return walletInfo, err
}

func UpdateSenderWalletRepository(wallet Wallet, amount float64) error {
	if wallet.Balance > amount {
		db := settings.ConnectToBD()
		newBalance := wallet.Balance - amount
		queryStr := fmt.Sprintf("UPDATE wallet SET balance = %f WHERE number = '%s'", newBalance, wallet.Number)
		_, err := db.Exec(queryStr)
		if err != nil {
			return err
		}
	} else {
		return errors.New("wallet balance is less than requested")
	}
	return nil
}

func UpdateRecipientWalletRepository(wallet Wallet, amount float64) error {
	db := settings.ConnectToBD()
	newBalance := wallet.Balance + amount
	queryStr := fmt.Sprintf("UPDATE wallet SET balance = %f WHERE number = '%s'", newBalance, wallet.Number)
	_, err := db.Exec(queryStr)
	if err != nil {
		return err
	}
	return nil
}

func SendMoneyToWalletRepository(from string, to string, amount float64) error {
	fromWallet, err := GetWalletByNumberRepository(from)
	if err != nil {
		return err
	}
	toWallet, err := GetWalletByNumberRepository(to)
	if err != nil {
		return err
	}
	err = UpdateSenderWalletRepository(fromWallet, amount)
	if err != nil {
		return err
	}
	err = UpdateRecipientWalletRepository(toWallet, amount)
	if err != nil {
		return err
	}
	return nil
}
