package internal

import (
	"encoding/json"
	"net/http"
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

	} else {
		s.init()
	}
}
