package fileutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

//OpenFileOrCreate will open the file at the given part and will create it if it doesn't exist.
//The bool at the end identifies if a file was created or opened, where TRUE means CREATED and FALSE means OPENED
func OpenFileOrCreate(pth string) (*os.File, error, bool) {
	_, e := os.Stat(pth)

	if e == nil {
		file, err := os.Open(pth)
		return file, err, false
	}

	dir := path.Dir(pth)

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		return nil, err, false
	}

	file, err := os.Create(pth)
	return file, err, true
}

func CreateIfNotExists(pth string) error {
	_, e := os.Stat(pth)

	if e == nil {
		return nil
	}

	dir := path.Dir(pth)

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		return err
	}

	file, err := os.Create(pth)
	if file != nil {
		_ = file.Close()
	}
	return err
}

func WriteIfNotExists(pth, content string) bool {
	_, e := os.Stat(pth)

	if e == nil {
		return false
	}

	dir := path.Dir(pth)

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		fmt.Println("Unable to write to file " + pth + ": " + err.Error())
		return false
	}

	file, err := os.Create(pth)

	if err != nil {
		fmt.Println("Unable to write to file " + pth + ": " + err.Error())
		return false
	}

	_, _ = file.WriteString(content)
	file.Close()
	return true
}

func ReadString(pth string) string {
	bytes, err := ioutil.ReadFile(pth)

	if err != nil {
		return ""
	}

	return string(bytes)
}
