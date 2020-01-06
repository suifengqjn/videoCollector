package account

const (
	VCFREE = iota
	VCVIPDAY
	VCVIPWEEk
	VCVIPMONTH
	VCVIPYEAR
	)



var VcAccount *Account

type Account struct {
	AccType int
	Count int64
	Start string
	End string
}

func GetAccount() *Account  {
	VcAccount = &Account{AccType:VCVIPDAY}

	return VcAccount
}

