package settings

import (
	json2 "encoding/json"
	"log"
)

type ErrorSchemas struct {
	Detail string `json:"detail"`
}

func (s ErrorSchemas) buildJson(detail string) ([]byte, error) {
	s.Detail = detail
	json, err := json2.Marshal(s)
	if err != nil {
		log.Println("error while marshaling json")
		return nil, err
	}
	return json, nil
}
