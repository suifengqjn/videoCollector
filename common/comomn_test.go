package common

import (
	"fmt"
	"io/ioutil"
	"myTool/ssrClient/check"
	"testing"
)

func TestClearTempFile(t *testing.T) {
	root := "/Users/qjn/Documents/2019"
	deep:= "/Users/qjn/Documents/2019/12/111/water"
	ClearTempFile(root, deep)
}


func TestLoadSSRAccounts(t *testing.T) {

	accs := LoadSSRAccounts()
	fmt.Println(len(accs))
	for _, a := range accs {
		fmt.Println(a)
		cli := check.MakeClient(a, target)

		if cli == nil {
			continue
		}
		res, err := cli.Get(target)

		if err != nil {
			continue
		}

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		fmt.Println(string(buf), err)
	}

}
