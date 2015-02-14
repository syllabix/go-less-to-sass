/*
Simple command line tool to assist in converting less projects to sass
*/

package main

import (
	"flag"
	"fmt"
	"github.com/syllabix/go-less-to-sass/converter"
	"github.com/syllabix/go-less-to-sass/regexes"
	"io/ioutil"
	"os"
	"path/filepath"
)

var sassFile converter.DataStream

func main() {
	filename := flag.String("filename", "", "relative path to the less file you would like to convert to scss")
	flag.Parse()
	if *filename != "" {
		ch := converter.LessToSass(*filename)
		sassFile := <-ch
		if sassFile.Err != nil {
			fmt.Println(sassFile.Err)
		}
		writeSassFile(*filename, sassFile.Data)
	} else {
		fmt.Println("-- Converting Project to Sass --")
		err := filepath.Walk("../", inspect)
		fmt.Printf("File path return %v\n", err)
	}
}

func inspect(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	match, err := filepath.Match("*.less", info.Name())
	if match && err == nil {
		fmt.Printf("Less file found! -> %s\n", info.Name())
		ch := converter.LessToSass(path)
		sassFile := <-ch
		if sassFile.Err != nil {
			fmt.Println(sassFile.Err)
		}
		writeSassFile(path, sassFile.Data)
	}
	if err != nil {
		return err
	}
	return nil
}

func writeSassFile(filename string, file string) {
	newScssFile := regexes.LessFile.ReplaceAllLiteralString(filename, ".scss")
	err := ioutil.WriteFile(newScssFile, []byte(file), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
