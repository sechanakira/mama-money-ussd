package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	MESSAGE_1 = `Welcome to Mama Money!
Where would you like to send money today
1) Kenya
2) Malawi`
	MESSAGE_2 = `How much money (in Rands)
would you like to send to 
%s`
	MESSAGE_3 = `Your person you are sending to 
will receive %s
%s
1) OK`
	MESSAGE_4 = `Thank you for using Mama Money!`
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var req UssdRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	}

	s := UssdSession{
		sessionId: req.SessionId,
		msisdn:    req.Msisdn,
	}

	if s.hasSession() {
		s.refresh()

		if s.nextStage == "MENU_2" {
			handleMenu2(w, req, s)
		}

		if s.nextStage == "MENU_3" {
			handleMenu3(w, req, s)
		}

		if s.nextStage == "MENU_4" {
			s.clear()
			resp := UssdResponse{s.sessionId, MESSAGE_4}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&resp)
		}
	} else {
		s.init()
		resp := UssdResponse{s.sessionId, MESSAGE_1}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&resp)
	}
}

func handleMenu2(w http.ResponseWriter, req UssdRequest, s UssdSession) {
	var message string
	if req.UserEntry == "1" {
		s.countryName = "Kenya"
		s.nextStage = "MENU_3"
		message = fmt.Sprintf(MESSAGE_2, "Kenya")
		s.update()
	} else if req.UserEntry == "2" {
		s.countryName = "Malawi"
		s.nextStage = "MENU_3"
		message = fmt.Sprintf(MESSAGE_2, "Malawi")
		s.update()
	} else {
		message = "Invalid Response"
		s.clear()
	}
	resp := UssdResponse{s.sessionId, message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

func handleMenu3(w http.ResponseWriter, req UssdRequest, s UssdSession) {
	var msg string
	var resp UssdResponse
	amt, err := strconv.ParseFloat(req.UserEntry, 32)
	if err != nil {
		resp = UssdResponse{s.sessionId, "Invalid response"}
		s.clear()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&resp)
	}

	if s.countryName == "Malawi" {
		exchangeRate := 42.50
		toReceive := amt * exchangeRate
		s.foreignCurrencyCode = "MWK"
		msg = fmt.Sprintf(MESSAGE_3, toReceive, "MWK")
	}

	if s.countryName == "Kenya" {
		exchangeRate := 6.10
		toReceive := amt * exchangeRate
		s.foreignCurrencyCode = "KWS"
		msg = fmt.Sprintf(MESSAGE_3, toReceive, "KWS")
	}

	s.amount = float32(amt)
	s.nextStage = "MENU_4"
	s.update()

	resp = UssdResponse{s.sessionId, msg}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}
