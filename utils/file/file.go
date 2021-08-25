package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// 检查文件是否存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if CheckNotExist(src) {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// 创建目录
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {

	return nil, nil
}
