// Package wallet предоставляет функции для работы с кошельками
package wallet

import (
	"database/sql"
	"fmt"
	"infotecs_go/src/settings"
	"infotecs_go/src/transaction"
)

// GetWalletByNumberRepository - Функция, для получения информации о кошельке по его номеру
//
// Аргументы: number string - номер кошелька
//
// Возвращаемые значения - error при ошибке получения кошелька, Wallet при удачном получении кошелька
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

// SendMoneyUpdateRecipientWallet - Функция, для обновления кошелька **получателя** при отправке денег
//
// Аргументы: tx *sql.Tx - транзакция бд, recipientWallet Wallet - информация о кошельке получателя,
// amount float64 - сумма перевода
//
// Возвращаемые значения - error при ошибке обновления кошелька, nil при удачном обновлении кошелька
func SendMoneyUpdateRecipientWallet(tx *sql.Tx, recipientWallet Wallet, amount float64) error {
	newBalance := recipientWallet.Balance + amount
	queryStr := fmt.Sprintf("UPDATE wallet SET balance = %f WHERE number = '%s'", newBalance,
		recipientWallet.Number)
	_, err := tx.Exec(queryStr)
	if err != nil {
		return err
	}
	return nil
}

// SendMoneyUpdateSenderWallet - Функция, для обновления кошелька отправителя при отправке денег
//
// Аргументы: tx *sql.Tx - транзакция бд, senderWallet Wallet - информация о кошельке отправителя,
// amount float64 - сумма перевода
//
// Возвращаемые значения - error при ошибке обновления кошелька, nil при удачном обновлении кошелька
func SendMoneyUpdateSenderWallet(tx *sql.Tx, senderWallet Wallet, amount float64) error {
	if senderWallet.Balance > amount {
		newBalance := senderWallet.Balance - amount
		queryStr := fmt.Sprintf("UPDATE wallet SET balance = %f WHERE number = '%s'", newBalance,
			senderWallet.Number)
		_, err := tx.Exec(queryStr)
		if err != nil {
			return err
		}
	} else {
		return settings.LowBalance
	}
	return nil
}

// SendMoneyToWalletRepository - Функция, для перевода денег с одного кошелька на другой
//
// Аргументы: sender string - номер кошелька отправителя, recipient string - номер кошелька - получателя,
// amount float64 - сумма перевода
//
// Возвращаемые значения - error при ошибке обновления кошелька, nil при удачном обновлении кошелька
func SendMoneyToWalletRepository(sender string, recipient string, amount float64) error {
	senderWallet, err := GetWalletByNumberRepository(sender)
	if err != nil {
		return err
	}
	recipientWallet, err := GetWalletByNumberRepository(recipient)
	if err != nil {
		return err
	}
	db := settings.ConnectToBD()
	tx, _ := db.Begin()
	defer tx.Rollback()
	err = SendMoneyUpdateSenderWallet(tx, senderWallet, amount)
	if err != nil {
		return err
	}
	err = SendMoneyUpdateRecipientWallet(tx, recipientWallet, amount)
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
