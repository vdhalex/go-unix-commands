package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main () {
	file := os.Args[1]

	f, err := os.Open(file)
	check(err)
	fi, err := f.Stat()
	sz := fi.Size()

	in := bufio.NewReader(f)
	out, err := in.Peek(int(sz))
	check(err)

	fmt.Println(string(out))
	f.Close()
}