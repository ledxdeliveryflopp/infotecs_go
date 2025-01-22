package tests

import (
	"bytes"
	"encoding/json"
	"infotecs_go/src/wallet"
	"testing"
)

// Тестирование функции десериализации wallet.Wallet
func TestWalletStructDecode(t *testing.T) {
	var schemas wallet.Wallet
	schemas.Number = "fgreas543"
	schemas.Balance = 131.1
	expected, _ := json.Marshal(schemas)
	json.Unmarshal(expected, &schemas)
	result, _ := schemas.DecodeJson(bytes.NewReader(expected))
	if result != schemas {
		t.Errorf("Result was incorrect, got: \"%+v\\n\", want: \"%+v\\n\".", result, schemas)
	}
}

// Тестирование функции десериализации wallet.SendMoneySchemas
func TestSendMoneyStructDecode(t *testing.T) {
	var schemas wallet.SendMoneySchemas
	schemas.From = "fgreas543"
	schemas.To = "sdfa3"
	schemas.Amount = 131.1
	expected, _ := json.Marshal(schemas)
	json.Unmarshal(expected, &schemas)
	result, _ := schemas.DecodeJson(bytes.NewReader(expected))
	if result != schemas {
		t.Errorf("Result was incorrect, got: \"%+v\\n\", want: \"%+v\\n\".", result, schemas)
	}
}
