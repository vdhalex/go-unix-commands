package main

import (
	"os"
	"path/filepath"
)

func check (e error) {
	if e != nil {
		panic(e)
	}
}

func main () {
	in := os.Args[1:]
	var flag byte = 0

	for _, val := range in {
		if val[0] == '-' && len(val) == 2 {
			flag = val[1]
			continue
		}

		if (flag != 0) {
			switch flag {
			case 'r':
				rrm(val)
			default:
				panic("dash cmd undefined")
			}
		} else {

			err := os.Remove(val)

			check(err)
		}
	}
}

func rrm(dir string) {
	var files[] string
	pwd, err := os.Getwd()

	check(err)

	err = os.Chdir(pwd+"/"+dir)

	check(err)
	pwd, err = os.Getwd()
	err = filepath.Walk(pwd, func (path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	
	check(err)
	curr_dir := files[0]
	files = files[1:]
	files = append(files, curr_dir)

	for _, file := range files {
		err = os.Remove(file)
		check(err)
	}
}