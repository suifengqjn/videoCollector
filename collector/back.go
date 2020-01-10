package collector

import (
	"myProject/videoCollector/common"
	"myTool/ffmpeg"
	"os"
	"path/filepath"
	"strings"
)
// 下载后处理
func (c *Collector)CheckRemove(path string) {

	con := c.getCondition()
	if con == nil {
		return
	}

	var remove = false
	if common.IsVideo(path) {

		info, err := ffmpeg.GetVideoInfo("", path)
		if err == nil && info != nil {
			//1. 宽高
			if con.Width > 0 {
				if info.W <= con.Width {
					remove = true
				}
			}

			if con.Height > 0 {
				if info.H <= con.Height {
					remove = true
				}
			}

			//2. 横竖屏
			if len(con.Direction) > 0 {
				if con.Direction == "h" { //横版视频
					if info.W <= info.H {
						remove = true
					}

				} else if con.Direction == "v" {
					if info.H <= info.W {
						remove = true
					}
				}

			}

			//3. 大小
			min := con.Size[0] * 1024 * 1024
			max := con.Size[1] * 1024 * 1024
			if min > 0 {
				if info.Capacity <= int64(min) {
					remove = true
				}

			}

			if max > 0 {
				if info.Capacity >= int64(max) {
					remove = true
				}
			}

		}

	}

	if remove {
		fName := filepath.Base(path)
		descPath := filepath.Dir(path) + "/" + strings.Split(fName, ".")[0] + ".txt"
		_ = os.Remove(path)
		_ = os.Remove(descPath)
	}

}

func(c *Collector) getCondition() *common.Condition  {
	return common.ReadConfig().Condition
}