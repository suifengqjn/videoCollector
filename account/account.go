package account

const (
	VCBASE = iota
	VCVIPMONTH
	VCVIPYEAR
	VCVIPSUPERVIP
)



var VcAccount *Account

func GetAccount(appId string) *Account  {
	VcAccount = GetAccountInfo(appId)
	return VcAccount
}

func (a *Account)DownloadAction()  {

	//if a.AccType <= 0 || a.AccType == VCVIPSUPERVIP{
	//	return
	//}
	//
	//a.Lock.Lock()
	//defer a.Lock.Unlock()
	//err := a.addCount()
	//if err == nil {
	//	a.Count --
	//}

}