# Golang Cash Register

**A cash register program in Golang**

### Download
	go get -u github.com/roy105703007/cashier

### Usage
    ```go
    import "github.com/roy105703007/cashier"
    ```

### Setting
	//VipDiscount record discount rates for different membership levels
	cashier.VipDiscount = make(map[uint8]float32)
	cashier.VipDiscount[1] = 0.95
	cashier.VipDiscount[2] = 0.9
	cashier.VipDiscount[3] = 0.85
	cashier.ExchangeRate = 1 // token/point

### Create Member
	roy := cashier.Member{
		Name:   "roy",
		Vip:    2,
		Points: 1000,
		Tokens: 1000,
	}

### Create Item

	car := cashier.Item{
		Cost:       40,
		ChargeMode: 2,
	}

### Charge
	var success bool
	var point float32 // The user can choose how many points he want to use.
	success = cashier.NormalCharge(car.Cost, &roy)
	success = cashier.VipCharge(car.Cost, &roy)
	success = cashier.PointsCharge(car.Cost, &roy, point)
	success = cashier.NewCharge(car.Cost, &roy, point)

### Test
	package main
	import (
		"fmt"
		"github.com/roy105703007/cashier"
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

### Output
	Now have Point:1000 and Token:1000
	Use NormalCharge buy one car($100)
	Success now have Point: 1000.00 and Token: 900.00
	Use VipCharge buy one car($100)
	Success now have Point: 1000.00 and Token: 810.00
	Use PointsCharge buy one car($100)
	Success now have Point: 900.00 and Token: 810.00
	Use NewCharge buy one car($100)
	Success now have Point: 800.00 and Token: 810.00
