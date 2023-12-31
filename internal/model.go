package internal

import "time"

type UssdRequest struct {
	sessionId string
	msisdn    string
	userEntry string
}

type UssdResponse struct {
	sessionId string
	message   string
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
