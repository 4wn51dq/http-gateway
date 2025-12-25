package middleware

import (
	"building_api/api"
	"building_api/internal/tools"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedErr = errors.New("Invalid Username or Token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* all logic for authorizing http request */
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		var err error
		var errCode int

		switch {
		case username == "" || token == "":
			log.Error(UnAuthorizedErr)
			api.RequestErrorHandler(w, err, errCode)
		default:
			var db *tools.DatabaseInterface
			db, err = tools.NewDatabase()
			if err != nil {
				api.InternalErrorHandler(w, errCode)
			}

			var loginDetails *tools.LoginDetails
			loginDetails = (*db).GetUserLoginDetails(username)
			if loginDetails == nil || token != (*loginDetails).AuthToken {
				log.Error(UnAuthorizedErr)
				api.RequestErrorHandler(w, err, errCode)
				return
			}
		}

		next.ServeHTTP(w, r) // calls the next middleware in line or the handler function for the endpoints

	})
	// response writer to construct reponse to the caller
	// request contains information about the incoming request (headers, payload etc)
}
