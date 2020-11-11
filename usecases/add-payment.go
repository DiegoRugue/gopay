package usecases

import "gopay/models"

func AddPayment(card models.Card, value float64) (success bool) {
	success = card.Pay(value)
	return
}
