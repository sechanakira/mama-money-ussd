package internal

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
	nextStage           string
	countryName         string
	amount              string
	foreingCurrencyCode string
}
