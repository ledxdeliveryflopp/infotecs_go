// Package wallet предоставляет функции для работы с кошельками
package wallet

import (
	"database/sql"
	"infotecs_go/src/settings"
	"infotecs_go/src/transaction"
)

func GetWalletByNumberRepository(number string) (Wallet, error) {
	walletFromCache, err := getWalletFromRedis(number)
	if err != nil {
		db := settings.ConnectToBD()
		row := db.QueryRow("SELECT number, balance FROM wallet WHERE number = $1", number)
		var walletInfo Wallet
		err := row.Scan(&walletInfo.Number, &walletInfo.Balance)
		if err != nil {
			return walletInfo, err
		}
		saveWalletInRedis(&walletInfo)
		return walletInfo, err
	} else {
		return walletFromCache, nil
	}
}

// __getWalletByNumberRepository - Функция, для получения информации о кошельке по его номеру
//
// Аргументы: number string - номер кошелька
//
// Возвращаемые значения - error при ошибке получения кошелька, Wallet при удачном получении кошелька
func __getWalletByNumberRepository(number string) (Wallet, error) {
	db := settings.ConnectToBD()
	row := db.QueryRow("SELECT number, balance FROM wallet WHERE number = $1", number)
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
func SendMoneyUpdateRecipientWallet(tx *sql.Tx, recipientWallet *Wallet, amount float64) error {
	recipientWallet.Balance = recipientWallet.Balance + amount
	_, err := tx.Exec("UPDATE wallet SET balance = $1 WHERE number = $2", recipientWallet.Balance,
		recipientWallet.Number)
	if err != nil {
		return err
	}
	saveWalletInRedis(recipientWallet)
	return nil
}

// SendMoneyUpdateSenderWallet - Функция, для обновления кошелька отправителя при отправке денег
//
// Аргументы: tx *sql.Tx - транзакция бд, senderWallet Wallet - информация о кошельке отправителя,
// amount float64 - сумма перевода
//
// Возвращаемые значения - error при ошибке обновления кошелька, nil при удачном обновлении кошелька
func SendMoneyUpdateSenderWallet(tx *sql.Tx, senderWallet *Wallet, amount float64) error {
	if senderWallet.Balance > amount {
		senderWallet.Balance = senderWallet.Balance - amount
		_, err := tx.Exec("UPDATE wallet SET balance = $1 WHERE number = $2", senderWallet.Balance,
			senderWallet.Number)
		if err != nil {
			return err
		}
	} else {
		return settings.LowBalance
	}
	saveWalletInRedis(senderWallet)
	return nil
}

// SendMoneyToWalletRepository - Функция, для перевода денег с одного кошелька на другой
//
// Аргументы: sender string - номер кошелька отправителя, recipient string - номер кошелька - получателя,
// amount float64 - сумма перевода
//
// Возвращаемые значения - error при ошибке обновления кошелька, nil при удачном обновлении кошелька
func SendMoneyToWalletRepository(sender string, recipient string, amount float64) error {
	senderWallet, err := __getWalletByNumberRepository(sender)
	if err != nil {
		return err
	}
	recipientWallet, err := __getWalletByNumberRepository(recipient)
	if err != nil {
		return err
	}
	db := settings.ConnectToBD()
	tx, _ := db.Begin()
	defer tx.Rollback()
	err = SendMoneyUpdateSenderWallet(tx, &senderWallet, amount)
	if err != nil {
		return err
	}
	err = SendMoneyUpdateRecipientWallet(tx, &recipientWallet, amount)
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
