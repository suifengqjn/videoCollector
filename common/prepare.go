package common
//
//import (
//	"myProject/videoCollector/library/cache/store"
//	"myProject/videoCollector/library/filter/toutiao"
//	"myProject/videoCollector/library/logger"
//)
//
////判断是否已经下载过，过滤平台
//
//func Prepare(videos []*VideoModel, dbFile string, filter []string) []*VideoModel   {
//
//	vs := hasDownLoad(videos, dbFile)
//
//	return filterPlatform(vs, filter)
//
//}
//
//func hasDownLoad(videos []*VideoModel, dbFile string) []*VideoModel  {
//	db, err := store.OpenDB(dbFile)
//
//	if err != nil {
//		logger.Println("数据库打开失败")
//		return videos
//	}
//	defer db.Close()
//
//	var res []*VideoModel
//	for _, v := range videos {
//		exist := db.Has(v.SaveKey)
//		if !exist {
//			res = append(res, v)
//		}
//	}
//	return res
//}
//
//func filterPlatform(videos []*VideoModel, filter []string) []*VideoModel {
//
//	if len(filter) == 0 {
//		return videos
//	}
//
//	var res []*VideoModel
//	for _, v := range videos {
//		for _, f := range filter {
//			if f == "t" {
//				has := toutiao.CompareTitleExist(v.Title)
//				if !has {
//					res = append(res, v)
//				}
//
//			}
//		}
//	}
//
//	return res
//}