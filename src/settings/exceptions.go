package settings

import (
	"log"
	"net/http"
)

func RaiseError(writer http.ResponseWriter, request *http.Request, detail string, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson(detail)
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

func WalletDontFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("wallet with this number don't found")
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

func TransactionsDontFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("transactions don't found")
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

func EncodingError(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("error while encoding struct")
	if err != nil {
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while error write json", err)
		return
	}
}

func HttpWriteError(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("error while send response")
	if err != nil {
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while error write json", err)
		return
	}
}

func QueryParamConvertError(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("error while converting query param")
	if err != nil {
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while error write json", err)
		return
	}
}

func NotEnoughMoneyInWallet(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("wallet balance is less than requested")
	if err != nil {
		log.Println("error while build error json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		log.Println("error while error write json", err)
		return
	}
}

func NotFoundEndpoint(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	var errorSchemas ErrorSchemas
	json, err := errorSchemas.BuildJson("endpoint not found")
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
	json, err := errorSchemas.BuildJson("method Not Allowed")
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
