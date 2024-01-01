package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	MESSAGE_1 = `Welcome to Mama Money!
Where would you like to send money today
1) Kenya
2) Malawi`
	MESSAGE_2 = `How much money (in Rands)
would you like to send to 
%s`
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
			json.NewEncoder(w).Encode(resp)
		}
	} else {
		s.init()
		resp := UssdResponse{s.sessionId, MESSAGE_1}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&resp)
	}
}
