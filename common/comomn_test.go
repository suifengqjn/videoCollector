package common

import (
	"fmt"
	"testing"
)

func TestClearTempFile(t *testing.T) {
	root := "/Users/qjn/Documents/2019"
	deep:= "/Users/qjn/Documents/2019/12/111/water"
	ClearTempFile(root, deep)
}


func TestLoadSSRAccounts(t *testing.T) {



}

func TestExtractDesc2(t *testing.T) {
	str := "🤩 民視優質戲劇搶先看 🤩\n🔔免費訂閱【民視戲劇館】👉🏻 https://reurl.cc/NE1O9\n\n🎁 更多精彩官方影片，請關注我們 🎁\n四季線上影視FB👉🏻 https://goo.gl/xzMmw3\n民視娛樂FB👉🏻 https://goo.gl/EFflxq\n\n📺 四季線上影視4gTV 免費戲劇新上架 📺\n四季線上影視4gTV：https://4gtv.tv\nAPP下載：https://reurl.cc/zyW56y"
	s := ExtractDesc(str,300)
	fmt.Println(s)
}
