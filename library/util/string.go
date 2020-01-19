package util

import (
	"crypto/md5"
	"fmt"
)

func Md5String(str string)string  {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(str))
	Result := Md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x", Result)
}
