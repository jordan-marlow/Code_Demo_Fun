package credit_card_check

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	CardNumber string `json:"card_number"`
}

type Response struct {
	Valid bool   `json:"valid"`
	Maker string `json:"maker"`
}

// LuhnHandler handles the HTTP request to validate a credit card number.
func LuhnHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	valid := isValidCreditCard(req.CardNumber)
	maker := getCreditCardMaker(req.CardNumber)
	res := Response{Valid: valid, Maker: maker}
	json.NewEncoder(w).Encode(res)
}
