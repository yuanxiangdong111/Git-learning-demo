package utils

import (
	"io"
	"os"
)

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func PathExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CreateDir(path string) error {
	if !PathExists(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func CreateFile(filePath string) (f *os.File, err error) {
	f, err = os.Create(filePath)
	return
}

func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

func LoadDataFrom(filePath string) ([]byte, error) {
	return LoadFromFileOrHTTP(filePath)
}

func WriteDataTo(filePath string, data string) error {
	var f *os.File
	var err error
	if FileExists(filePath) {
		f, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	} else {
		f, err = CreateFile(filePath)
	}
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.WriteString(f, data)
	return err
}
