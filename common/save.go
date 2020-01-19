package common

import (
	"encoding/json"
	"myProject/videoCollector/library/cache/store"
)

func SaveRecord(video *VideoModel,dbFile string) {
	db, err := store.OpenDB(dbFile)
	if err != nil {
		return
	}
	defer db.Close()
	buf, err := json.Marshal(&video)
	if err != nil {
		_ = db.Save(video.SaveKey, []byte("1"))
	} else {
		_ = db.Save(video.SaveKey, buf)
	}

}
