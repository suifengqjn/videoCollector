package account

const (
	VCFREE = iota
	VCVIPDAY
	VCVIPWEEk
	VCVIPMONTH
	VCVIPYEAR
	)



var VcAccount *Account

func GetAccount(appId string) *Account  {
	VcAccount = GetAccountInfo(appId)
	return VcAccount
}

func (a *Account)DownloadAction()  {

	a.Count --
	a.Add()
}