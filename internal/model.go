package internal

import "time"

type UssdRequest struct {
	SessionId string `json:"sessionId"`
	Msisdn    string `json:"msisdn"`
	UserEntry string `json:"userEntry"`
}

type UssdResponse struct {
	SessionId string `json:"sessionId"`
	Message   string `json:"message"`
}

type UssdSession struct {
	sessionId           string
	msisdn              string
	nextStage           string
	countryName         string
	amount              float32
	foreignCurrencyCode string
	sessionStartTime    time.Time
}
