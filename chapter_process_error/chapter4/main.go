package main

import (
	"fmt"

	"chapter_process_error/frame"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func findMysqlFile() (*File, error) {
	err := fmt.Errorf("文件不存在")
	if err != nil {
		return nil, err
	}
	return &File{}, nil
}
func findServiceFile1() (*File, error) {
	file, err := findMysqlFile()
	if err != nil {
		return nil, fmt.Errorf("findServiceFile fail, err: %w", err)
	}
	return file, nil
}
func findServiceFile2() (*File, error) {
	file, err := findMysqlFile()
	if err != nil {
		return nil, fmt.Errorf("findServiceFile2 fail, err: %w", err)
	}
	return file, nil
}
func findControllerFile(c *gin.Context) {
	var err error
	if c.Param("type") == "1" {
		_, err = findServiceFile1()
	} else {
		_, err = findServiceFile2()
	}
	if err != nil {
		frame.LoggerError("findControllerFile fail", zap.Error(err))
		return
	}
}
func main() {
	findControllerFile(&gin.Context{})
}

type File struct {
	FileId string
}
