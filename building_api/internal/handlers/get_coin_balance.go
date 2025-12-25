package handlers

import (
	"building_api/api"
	"building_api/internal/tools"
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Response) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	var errCode int

	err = decoder.Decode(&params, r.Request.URL.Query())
	if err != nil {
		api.InternalErrorHandler(w, errCode)
		return
	}
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w, errCode)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w, errCode)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, errCode)
		return
	}
}
