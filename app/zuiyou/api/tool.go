package api

import (
	"fmt"
	"myProject/videoCollector/app/zuiyou/model"
	"myProject/videoCollector/common"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func TrimTitle(title string) string {

	re, _ := regexp.Compile(`#.+#`)

	str := re.ReplaceAllString(title, "")

	return strings.TrimSpace(str)
}

func DownloadDir() string {

	return "./" + time.Now().Format(common.DownloadTimeFormat) + "/zuiyou"
}

func SaveKey(url string) []byte {
	return []byte(fmt.Sprintf("zuiyou_%v", url))
}

func ParseVideoDetail(vs []*model.ZyVideo) []*common.VideoModel {

	var res []*common.VideoModel
	for _, v := range vs {

		vm := common.VideoModel{}
		vm.Title = TrimTitle(v.Content)
		vm.ID = fmt.Sprintf("%v", v.ID)
		vm.DownLoadDir = DownloadDir()
		vm.Platform = common.PlatZY
		for _, video := range v.Videos {
			if vm.Url == "" {
				vm.Url = video.Urlsrc
			}
		}
		vm.SaveKey = SaveKey(vm.Url)
		detail := common.VideoDetail{}
		detail.Tags = append(detail.Tags, v.Topic.Topic)
		detail.Tags = append(detail.Tags, v.Topic.AttsTitle)
		detail.Tags = append(detail.Tags, v.Member.Name)

		if len(v.GodReviews) > 0 {
			for _, g := range v.GodReviews {
				detail.Desc = detail.Desc + "-" + g.Review
			}
		}

		vm.Detail = &detail

		vm.Up = v.Up
		vm.Down = v.Down
		vm.ShareCount = v.Share
		vm.CommentCount = v.Reviews
		if len(v.Imgs) > 0 {
			for _, img := range v.Imgs[0].Urls {
				if vm.Pic == "" {
					if len(img.Urls) > 0 {
						vm.Pic = img.Urls[0]
						break
					}

				}
			}
		}

		res = append(res, &vm)

	}

	return res
}

func GetHts()string  {
	return strconv.Itoa(int(time.Now().UnixNano()/1000000))
}