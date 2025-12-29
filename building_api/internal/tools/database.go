package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    uint
	Username string
}

// methods required for API
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &mockDB{}

	var err error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &db, nil
}
