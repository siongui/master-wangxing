package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

const dirPath = "../../userpages/content/articles"

func isMingyi(fpath string) bool {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		panic(err)
	}

	isMatch, err := regexp.Match("萬行法師", b)
	if err != nil {
		panic(err)
	}
	return isMatch
}

func main() {
	count := 0
	// walk all files in directory
	filepath.Walk(dirPath, func(fpath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if isMingyi(fpath) {
				newpath := path.Join("../content/articles/", path.Base(fpath))
				println(newpath)
				err := os.Rename(fpath, newpath)
				if err != nil {
					println(newpath)
					println(err)
				}
				count++
			}
		}
		return nil
	})
	println(count)
}
