package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	year := req.URL.Query().Get("year")
	res.Header().Set("Content-Type", "application/json")
	if year == "" {
		winners, err := data.ListAllJSON()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Write(winners)
	} else {
		var filteredWinners, err = data.ListAllByYear(year)
		if err!= nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		res.Write(filteredWinners)

	}

}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("X-ACCESS-TOKEN")
	isAccessTokenValid := data.IsAccessTokenValid(accessToken)
	if !isAccessTokenValid {
		res.WriteHeader(http.StatusUnauthorized)
	} else {
		err := data.AddNewWinner(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		res.WriteHeader(http.StatusCreated)

	}
}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		ListWinners(res,req)
	case http.MethodPost:
		AddNewWinner(res,req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}
