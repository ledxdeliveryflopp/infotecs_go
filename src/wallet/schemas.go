// Package wallet предоставляет функции для работы с кошельками
package wallet

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"reflect"
)

// BaseSchemas - Базовая информационна структура
type BaseSchemas struct {
	Detail string `json:"detail"`
}

// BuildJson - функция для сериализации структуры BaseSchemas
//
// # Аргументы - detail string - информация
//
// Возвращаемые значения - error при ошибке сериализации, []byte при удачной сериализации
func (b BaseSchemas) BuildJson(detail string) ([]byte, error) {
	b.Detail = detail
	marshaled, err := json.Marshal(b)
	if err != nil {
		log.Println("error while marshaling json")
		return marshaled, err
	}
	return marshaled, nil
}

// Wallet - Структура кошелька
type Wallet struct {
	Number  string  `json:"number"`
	Balance float64 `json:"balance"`
}

// MarshalBinary - функция для сериализации структуры *Wallet для Redis
//
// Возвращаемые значения - error при ошибке сериализации, []byte при удачной сериализации
func (w *Wallet) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}

// UnmarshalBinary - функция для десериализации структуры *Wallet для Redis
//
// # Аргументы - data []byte - строка из Redis
//
// Возвращаемые значения - error при ошибке сериализации
func (w *Wallet) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &w)
}

// SendMoneySchemas - Структура перевода денег с кошелька
type SendMoneySchemas struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

// DecodeJson - функция для десериализации структуры SendMoneySchemas
//
// Аргументы - body io.Reader - тело запроса(json)
//
// Возвращаемые значения - error при ошибке десериализации, SendMoneySchemas при удачной десериализации
func (m *SendMoneySchemas) DecodeJson(body io.Reader) error {
	err := json.NewDecoder(body).Decode(&m)
	if err != nil {
		log.Println("error while unmarshal json", err)
		return err
	}
	switch {
	case reflect.ValueOf(m.From).IsZero() == true:
		return errors.New("empty from field")
	case reflect.ValueOf(m.To).IsZero() == true:
		return errors.New("empty to field")
	case reflect.ValueOf(m.Amount).IsZero() == true:
		return errors.New("empty amount field")
	}
	return nil
}
