package wallet

import (
	"database/sql"
	"fmt"
	"infotecs_go/src/settings"
	"infotecs_go/src/transaction"
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

func SendMoneyUpdateToWallet(tx *sql.Tx, toWallet Wallet, amount float64) error {
	newBalance := toWallet.Balance + amount
	queryStr := fmt.Sprintf("UPDATE wallet SET balance = %f WHERE number = '%s'", newBalance, toWallet.Number)
	_, err := tx.Exec(queryStr)
	if err != nil {
		return err
	}
	return nil
}

func SendMoneyUpdateFromWallet(tx *sql.Tx, wallet Wallet, amount float64) error {
	if wallet.Balance > amount {
		newBalance := wallet.Balance - amount
		queryStr := fmt.Sprintf("UPDATE wallet SET balance = %f WHERE number = '%s'", newBalance, wallet.Number)
		_, err := tx.Exec(queryStr)
		if err != nil {
			return err
		}
	} else {
		return settings.LowBalance
	}
	return nil
}

func SendMoneyToWalletRepository(sender string, recipient string, amount float64) error {
	fromWallet, err := GetWalletByNumberRepository(sender)
	if err != nil {
		return err
	}
	toWallet, err := GetWalletByNumberRepository(recipient)
	if err != nil {
		return err
	}
	db := settings.ConnectToBD()
	tx, _ := db.Begin()
	defer tx.Rollback()
	err = SendMoneyUpdateFromWallet(tx, fromWallet, amount)
	if err != nil {
		return err
	}
	err = SendMoneyUpdateToWallet(tx, toWallet, amount)
	if err != nil {
		return err
	}
	err = transaction.CreateTransactionRepository(tx, sender, recipient, amount)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
