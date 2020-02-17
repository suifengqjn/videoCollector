package common

import (
	"myTool/ffmpeg"
	"myTool/file"
	"os"
	"path/filepath"
	"strings"
)
func CoverToMp4(dir string) string {

	files, _ := file.GetCurrentFiles(dir)

	resultPath := ""
	for _, f := range files {
		re := ffmpeg.CoverToMp4("", f)

		if resultPath == "" {
			resultPath = re
		}
	}

	return filepath.Dir(resultPath)

}



func GetVideoFiles(dir string) []string {
	files, err := file.GetCurrentFiles(dir)
	if err != nil {
		return nil
	}

	var vfiles []string
	for _, f := range files {
		if ffmpeg.IsVideo(f) {
			vfiles = append(vfiles, f)
		}
	}

	return vfiles

}

//删除不合适的视频
func VideoCutPrepare(con Condition, dir string) {

	files, err := file.GetCurrentFiles(dir)
	if err != nil {
		return
	}

	var removeArr []string
	for _, f := range files {
		if ffmpeg.IsVideo(f) {

			info, err := ffmpeg.GetVideoInfo("", f)
			if err == nil && info != nil {
				//1. 宽高
				if con.Width > 0 {
					if info.W <= con.Width {
						removeArr = append(removeArr, f)
						continue
					}
				}

				if con.Height > 0 {
					if info.H <= con.Height {
						removeArr = append(removeArr, f)
						continue
					}
				}

				//2. 横竖屏
				if len(con.Direction) > 0 {
					if con.Direction == "h" { //横版视频
						if info.W <= info.H {
							removeArr = append(removeArr, f)
							continue
						}

					} else if con.Direction == "v" {
						if info.H <= info.W {
							removeArr = append(removeArr, f)
							continue
						}
					}

				}

				//3. 大小
				min := con.Size[0] * 1024 * 1024
				max := con.Size[1] * 1024 * 1024
				if min > 0 {
					if info.Capacity <= int64(min) {
						removeArr = append(removeArr, f)
						continue
					}

				}

				if max > 0 {
					if info.Capacity >= int64(max) {
						removeArr = append(removeArr, f)
						continue
					}
				}

			}

		}
	}

	//删除
	for _, f := range removeArr {
		fName := filepath.Base(f)
		descPath := filepath.Dir(f) + "/" + strings.Split(fName, ".")[0] + ".txt"
		_ = os.Remove(f)
		_ = os.Remove(descPath)
	}

}
