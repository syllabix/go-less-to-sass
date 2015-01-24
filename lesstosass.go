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
	filename := flag.String("filename", "", "relative path to the less file you would like to convert to sass")
	flag.Parse()
	fmt.Println(*filename)
	if *filename != "" {
		ch := converter.LessToSass(*filename)
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
