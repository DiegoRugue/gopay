package models

type DebitCard struct {
	card
	Balance float64
}

func (c *DebitCard) Pay(value float64) bool {
	if value < c.Balance  {
		c.Balance -= value
		return true
	}
	return false
}

func (c *DebitCard) Deposit(value float64) float64 {
	c.Balance += value
	return  c.Balance
}
