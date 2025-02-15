package tests

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"infotecs_go/src/wallet"
	"testing"
)

var (
	ctx = context.Background()
)

func TestRedisWallet(t *testing.T) {
	var schemasWallet wallet.Wallet
	schemasWallet.Balance = 100
	schemasWallet.Number = "test1"
	wallet.SaveWalletInRedis(&schemasWallet)
	walletFromRedis, err := wallet.GetWalletFromRedis("test1")
	if err != nil {
		t.Errorf("Error while testing redis: %s", err)
	}
	equal := cmp.Equal(schemasWallet, walletFromRedis)
	if equal == false {
		t.Errorf("Error while compare structs")
		fmt.Println("basic wallet: ", schemasWallet)
		fmt.Println("wallet from redis: ", walletFromRedis)

	}
}
