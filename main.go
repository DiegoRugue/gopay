package main

import (
	"fmt"
	"gopay/models"
	"gopay/usecases"
	"time"
)

func main() {
	var creditCard models.CreditCard
	creditCard.Name = "Diego Rugue"
	creditCard.Number = 123456789
	creditCard.Limit = 5000
	creditCard.ExpiringDate = time.Date(2022, time.November, 10, 23, 0, 0, 0, time.UTC)

	if success := usecases.AddPayment(&creditCard, 4999); success {
		fmt.Println("Compra realizada com sucesso!")
	} else {
		fmt.Println("Não foi dessa vez =/")
	}
}
