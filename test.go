package main

import (
	"fmt"

	"github.com/roy105703007/go-cash-register/tree/main/cashier"
)

func main() {
	roy := cashier.Member{
		Name:   "roy",
		Vip:    2,
		Points: 1000,
		Tokens: 1000,
	}
	car := cashier.Item{
		Cost:       100,
		ChargeMode: 2,
	}

	cashier.VipDiscount = make(map[uint8]float32)
	cashier.VipDiscount[1] = 0.95
	cashier.VipDiscount[2] = 0.9
	cashier.VipDiscount[3] = 0.85
	cashier.ExchangeRate = 1

	success := false

	fmt.Println("Now have Point:1000 and Token:1000")

	fmt.Println("Use NormalCharge buy one car($100)")
	success = cashier.NormalCharge(car.Cost, &roy)
	if success {
		fmt.Printf("Success now have Point: %.2f and Token: %.2f\n", roy.Points, roy.Tokens)
	} else {
		fmt.Println("Fail")
	}

	fmt.Println("Use VipCharge buy one car($100)")
	success = cashier.VipCharge(car.Cost, &roy)
	if success {
		fmt.Printf("Success now have Point: %.2f and Token: %.2f\n", roy.Points, roy.Tokens)
	} else {
		fmt.Println("Fail")
	}

	fmt.Println("Use PointsCharge buy one car($100)")
	success = cashier.PointsCharge(car.Cost, &roy, 100)
	if success {
		fmt.Printf("Success now have Point: %.2f and Token: %.2f\n", roy.Points, roy.Tokens)
	} else {
		fmt.Println("Fail")
	}

	fmt.Println("Use NewCharge buy one car($100)")
	success = cashier.NewCharge(car.Cost, &roy, 100)
	if success {
		fmt.Printf("Success now have Point: %.2f and Token: %.2f\n", roy.Points, roy.Tokens)
	} else {
		fmt.Println("Fail")
	}

}
