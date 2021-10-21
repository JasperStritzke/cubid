package fileutil

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

//OpenFileOrCreate will open the file at the given part and will create it if it doesn't exist.
//The bool at the end identifies if a file was created or opened, where TRUE means CREATED and FALSE means OPENED
func OpenFileOrCreate(pth string) (*os.File, error, bool) {
	_, e := os.Stat(pth)

	if e == nil {
		file, err := os.OpenFile(pth, os.O_RDWR, os.ModeAppend)
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

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func ExistsFile(filepath string) bool {
	_, e := os.Stat(filepath)

	return !os.IsNotExist(e)
}
