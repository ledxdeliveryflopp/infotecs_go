package wallet

import (
	"encoding/json"
	"io"
	"log"
)

type Wallet struct {
	Number  string  `json:"number"`
	Balance float64 `json:"balance"`
}

func (w Wallet) decodeJson(body io.Reader) (Wallet, error) {
	err := json.NewDecoder(body).Decode(&w)
	if err != nil {
		log.Println("error while unmarshal json", err)
		return w, err
	}
	return w, nil
}

type SendMoneySchemas struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func (m SendMoneySchemas) decodeJson(body io.Reader) (SendMoneySchemas, error) {
	err := json.NewDecoder(body).Decode(&m)
	if err != nil {
		log.Println("error while unmarshal json", err)
		return m, err
	}
	return m, nil
}
