package parser

type odds struct {
	price string
	num int32
	den int32
}

func Odds(price string) *odds {
	o := new(odds)
	o.price = price
	return o
}

func (o odds) getFractionalOdds() string {
	// returns fractional odds as string
	return ""
}

func (o odds) getDecimalOdds() float32 {
	// returns decimal price as double roundDown((num/den)+1, 2)
	return 0.00
}
