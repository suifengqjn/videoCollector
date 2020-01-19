package common

import (
	"testing"
)

func TestClearTempFile(t *testing.T) {
	root := "/Users/qjn/Documents/2019"
	deep:= "/Users/qjn/Documents/2019/12/111/water"
	ClearTempFile(root, deep)
}


func TestLoadSSRAccounts(t *testing.T) {



}
