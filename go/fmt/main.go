package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes.ReplaceAll(data, []byte("\t"), []byte("    "))))
}
