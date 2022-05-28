# Golang Cash Register

**A cash register program in Golang**

###Download
	go get -u github.com/roy105703007/cashier

###Usage
	import "github.com/roy105703007/cashier"

###Setting
	cashier.VipDiscount = make(map[uint8]float32)
	cashier.VipDiscount[1] = 0.95
	cashier.VipDiscount[2] = 0.9
	cashier.VipDiscount[3] = 0.85
	cashier.ExchangeRate = 1 // token/point

###Create Member
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

###Charge
	var success bool
	var point float32 // The user can choose how many points he want to use.
	success := cashier.NormalCharge(car.Cost, &roy)
	success = cashier.VipCharge(car.Cost, &roy)
	success = cashier.PointsCharge(car.Cost, &roy, point)
	success = cashier.NewCharge(car.Cost, &roy, point)