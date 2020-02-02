package common

import (
	"fmt"
	"testing"
)

func TestExtractDesc(t *testing.T) {
	desc := "#萌医甜妻 #萌医甜妻19\n✅ ️《奔騰年代》 FULL: http://bit.ly/2PHIDFw\n✅ ️《光荣时代》 FULL: http://bit.ly/2PiNJbp\n✅ ️《麦香》FULL:  http://bit.ly/2HuI0KZ\n✅ ️《无名卫士》 FULL:  http://bit.ly/2WIpHus\n---\n✅  Please Like, Share and Subscribe for me:  http://bit.ly/2K0PNTs\n---\n▶ ️If my MV infringes your copyrights, please kindly pm me and I'll immediately delete it. I beg you not to file any complaints directly to Youtube since it would cause my channel to be closed. Thank you for your kind understanding and I sincerely apologize for any inconvenience caused."

	s := ExtractDesc(desc, 300)

	fmt.Println(s)
	fmt.Println(ChineseLen(s))
}