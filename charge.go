package cashier

func NormalCharge(cost float32, member *Member) bool {
	//檢查會員平台幣餘額大於商品價格
	if (*member).Tokens < cost {
		return false
	}
	(*member).Tokens -= cost
	return true
}

func VipCharge(cost float32, member *Member) bool {
	//檢查會員為VIP
	if (*member).Vip != 0 {
		//檢查會員平台幣餘額大於商品價格
		if (*member).Tokens < cost*VipDiscount[(*member).Vip] {
			return false
		} else {
			(*member).Tokens -= cost * VipDiscount[(*member).Vip]
		}
	} else {
		if (*member).Tokens < cost {
			return false
		} else {
			(*member).Tokens -= cost
		}
	}
	return true
}

func PointsCharge(cost float32, member *Member, pointsWantToUse float32) bool {
	//檢查會員平台點數大於他設置的值
	if (*member).Points >= pointsWantToUse {
		//商品是否可以全部使用點數支付
		if cost > pointsWantToUse*ExchangeRate {
			cost -= pointsWantToUse * ExchangeRate
			//檢查會員平台幣餘額大於商品價格
			if (*member).Tokens < cost {
				return false
			} else {
				(*member).Points -= pointsWantToUse
				(*member).Tokens -= cost
			}
		} else {
			(*member).Points -= cost / ExchangeRate
		}
	} else {
		if (*member).Tokens < cost {
			return false
		} else {
			(*member).Tokens -= cost
		}
	}
	return true
}

func NewCharge(cost float32, member *Member, pointsWantToUse float32) bool {
	//檢查會員符合特殊活動條件
	if (*member).Points >= pointsWantToUse && pointsWantToUse >= 100 && (*member).Vip != 0 {
		//商品是否可以全部使用點數支付
		if cost > pointsWantToUse*ExchangeRate {
			cost -= pointsWantToUse * ExchangeRate
			cost *= 0.9
			//檢查會員平台幣餘額大於商品價格
			if (*member).Tokens < cost {
				return false
			} else {
				(*member).Points -= pointsWantToUse
				(*member).Tokens -= cost
			}
		} else {
			(*member).Points -= cost / ExchangeRate
		}
	} else {
		return false
	}
	return true
}
