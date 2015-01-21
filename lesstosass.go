package main

import (
	"flag"
	"fmt"
	"github.com/syllabix/go-less-to-sass/converter"
	"io/ioutil"
	//"os"
)

var sassFile converter.DataStream

func main() {
	ch := make(chan converter.DataStream)
	filename := flag.String("filename", "", "relative path to the less file you would like to convert to sass")
	flag.Parse()
	if *filename != "" {
		converter.LessToSass(*filename, ch)
		sassFile := <-ch
		if sassFile.Err != nil {
			fmt.Println(sassFile.Err)
		}
		writeSassFile(sassFile.Data)
	}
}

func writeSassFile(file string) {
	err := ioutil.WriteFile("test.scss", []byte(file), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
