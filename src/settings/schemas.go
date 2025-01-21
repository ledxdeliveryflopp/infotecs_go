package settings

import (
	"encoding/json"
	"log"
)

type ErrorSchemas struct {
	Detail string `json:"detail"`
}

func (s ErrorSchemas) buildJson(detail string) ([]byte, error) {
	s.Detail = detail
	marshalDetail, err := json.Marshal(s)
	if err != nil {
		log.Println("error while marshaling json")
		return nil, err
	}
	return marshalDetail, nil
}
