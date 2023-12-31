package internal

import (
	"fmt"
	"mama-money-ussd/internal/db"
)

type UssdSessionDb interface {
	init()
	update()
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
