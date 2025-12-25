package api

import (
	"encoding/json"
	"net/http"
)

// the parameter the API will take
type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int  // success code
	Balance uint // accountBalance
}

type Error struct {
	Code    int // error code
	Message string
}

func writeError(w http.ResponseWriter, errCode int, errMessage string) {
	resp := Error{
		Code:    errCode,
		Message: errMessage,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errCode)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error, errCode int) {
		writeError(w, errCode, err.Error())
	}
	InternalErrorHandler = func(w http.ResponseWriter, errCode int) {
		writeError(w, http.StatusInternalServerError, "Unexpected error occured")
	}
)
