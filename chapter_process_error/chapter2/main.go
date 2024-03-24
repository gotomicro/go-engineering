package main

import (
	"fmt"

	"chapter_process_error/frame"
	"go.uber.org/zap"
)

func findMysqlFile() (*File, error) {
	err := fmt.Errorf("文件不存在")
	if err != nil {
		return nil, err
	}
	return &File{}, nil
}
func findServiceFile() (*File, error) {
	file, err := findMysqlFile()
	if err != nil {
		return nil, err
	}
	return file, nil
}
func findControllerFile() {
	_, err := findServiceFile()
	if err != nil {
		frame.LoggerError("findControllerFile fail", zap.Error(err))
		return
	}
}
func main() {
	findControllerFile()
}

type File struct {
	FileId string
}
