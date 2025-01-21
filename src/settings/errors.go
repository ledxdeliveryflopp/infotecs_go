package settings

import "errors"

var LowBalance = errors.New("wallet balance is less than requested")

var TransactionNotFound = errors.New("transactions not found")
