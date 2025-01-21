package settings

import (
	"log"
	"net/http"
)

func RaiseError(writer http.ResponseWriter, request *http.Request, detail string, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.buildJson(detail)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while error write json", err)
		return
	}
}

func NotFoundEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.buildJson("endpoint not found")
	if err != nil {
		log.Println("error while build not found json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while error write json", err)
		return
	}
}

func MethodNotAllowed(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusMethodNotAllowed)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.buildJson("method Not Allowed")
	if err != nil {
		log.Println("error while build not allowed json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while not allowed write json", err)
		return
	}
}
