package models

type Card interface {
	Pay(value float64) bool
}

type card struct {
	Name string
	Number int
}
