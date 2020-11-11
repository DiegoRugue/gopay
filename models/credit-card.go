package models

import "time"

type CreditCard struct {
	card
	Limit float64
	ExpiringDate time.Time
}

func (c *CreditCard) Pay(value float64) bool {
	if value < c.Limit && c.ExpiringDate.Sub(time.Now()) > 0 {
		c.Limit -= value
		return true
	}
	return false
}
