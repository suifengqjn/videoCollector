package common

import (
	"myTool/file"
	"os"
	"path/filepath"
)

func ClearTempFile(rootDir string, deepDir string)  {

	files,err := file.GetCurrentFiles(deepDir)
	if err != nil {
		return
	}

	// 将最后的文件移到初始目录
	for _, f := range files {
		newPath := rootDir + "/" + filepath.Base(f)
		_= os.Rename(f, newPath)
	}

	// 删除临时文件夹
	_, dirs, err := file.GetAllFilesAndDirs(rootDir)
	if err != nil {
		return
	}

	for _, dir := range dirs {
		_ = os.RemoveAll(dir)
	}

}
