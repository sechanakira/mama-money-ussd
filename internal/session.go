package internal

import (
	"fmt"
	"mama-money-ussd/internal/db"
)

type UssdSessionDb interface {
	init()
	update()
	hasSession() bool
	refresh()
}

func (u *UssdSession) init() {
	db.InitUssdSession(u.sessionId, u.msisdn)
}

func (u *UssdSession) update() {
	values := make(map[string]string)
	values["nextStage"] = u.nextStage
	values["countryName"] = u.countryName
	values["amount"] = fmt.Sprintf("%f", u.amount)
	values["foreignCurrencyCode"] = u.foreignCurrencyCode
	db.UpdateUssdSession(u.sessionId, values)
}

func (u *UssdSession) hasSession() bool {
	us, err := db.FindSession(u.sessionId)

	if err != nil {
		return false
	}

	return us != nil
}

func (u *UssdSession) refresh() {
	us, _ := db.FindSession(u.sessionId)
	u.sessionId = us.SessionId
	u.msisdn = us.Msisdn
	u.nextStage = us.NextStage
	u.countryName = us.CountryName
	u.amount = us.Amount
	u.foreignCurrencyCode = us.ForeignCurrencyCode
	u.sessionStartTime = us.SessionStartTime
}
