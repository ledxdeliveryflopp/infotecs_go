// Package settings предоставляет функции для настройки приложения
package settings

import (
	"encoding/json"
	"log"
)

// ErrorSchemas - структура ошибки
type ErrorSchemas struct {
	Detail string `json:"detail"`
}

// BuildJson - функция для сериализации структуры ErrorSchemas
//
// # Аргументы - detail string - информация
//
// Возвращаемые значения - error при ошибке сериализации, []byte при удачной сериализации
func (s ErrorSchemas) BuildJson(detail string) ([]byte, error) {
	s.Detail = detail
	marshalDetail, err := json.Marshal(s)
	if err != nil {
		log.Println("error while marshaling json")
		return nil, err
	}
	return marshalDetail, nil
}
