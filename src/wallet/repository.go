package wallet

import (
	"fmt"
	"infotecs_go/src/settings"
)

func GetWalletByNumber(number string) (Wallet, error) {
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
