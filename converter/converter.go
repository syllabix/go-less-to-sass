package converter

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

type DataStream struct {
	Data string
	Err  error
}

var convertedFile string

func LessToSass(filename string, ch chan DataStream) {
	go func() {
		contents, err := ioutil.ReadFile(filename)
		convertedFile := convertMixins(string(contents))
		//convertedFile = convertVars(string(convertedFile))
		ch <- DataStream{convertedFile, err}
	}()
}

func convertVars(file string) string {
	re := regexp.MustCompile("@+")
	mediaRE := regexp.MustCompile("\\$media+")
	importRE := regexp.MustCompile("\\$import+")
	result := re.ReplaceAllString(file, "$")
	result = mediaRE.ReplaceAllString(result, "@media")
	result = importRE.ReplaceAllString(result, "@import")
	return result
}

func convertMixins(file string) string {
	re := regexp.MustCompile("\\.([a-zA-Z-_0-9])+\\(([a-zA-Z-_0-9])*\\)(\\s)*{")
	result := re.FindAllString(file, -1)
	for _, match := range result {
		regex := regexp.MustCompile(match)
		fmt.Println("@mixin " + match)
		file = regex.ReplaceAllString(file, "@mixin"+match)
	}
	return file
}
