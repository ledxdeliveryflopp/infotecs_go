package wallet

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"reflect"
)

type BaseStruct struct {
	Detail string `json:"detail"`
}

func (b BaseStruct) buildJson(detail string) ([]byte, error) {
	b.Detail = detail
	marshalDetail, err := json.Marshal(b)
	if err != nil {
		log.Println("error while marshaling json")
		return nil, err
	}
	return marshalDetail, nil
}

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
	switch {
	case reflect.ValueOf(m.From).IsZero() == true:
		return m, errors.New("empty from field")
	case reflect.ValueOf(m.To).IsZero() == true:
		return m, errors.New("empty to field")
	case reflect.ValueOf(m.Amount).IsZero() == true:
		return m, errors.New("empty amount field")
	}
	return m, nil
}
