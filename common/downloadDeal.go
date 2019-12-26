package common

import (
	"myTool/annie/downloader"
)

//下载时过滤
func FilterBySize(data downloader.Data) bool  {

	con := ReadConfig().Condition
	//3. 大小
	min := con.Size[0] * 1024 * 1024
	max := con.Size[1] * 1024 * 1024
	for _, v := range data.Streams {

		if min > 0 {
			if v.Size <= int64(min) {
				return false
			}

		}

		if max > 0 {
			if v.Size >= int64(max) {
				return false
			}
		}
	}

	return true

}
