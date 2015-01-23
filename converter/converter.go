package converter

import (
	"bufio"
	"bytes"
	//	"fmt"
	"os"
	"regexp"
	"strings"
)

type DataStream struct {
	Data string
	Err  error
}

var convertedFile string
var stringBuffer bytes.Buffer

func LessToSass(filename string) chan DataStream {
	ch := make(chan DataStream)
	go func() {
		file, err := os.Open(filename)
		defer file.Close()
		if err == nil {
			convertedFile = convert(file)
		} else {
			convertedFile = ""
		}
		ch <- DataStream{convertedFile, err}
	}()
	return ch
}

func convert(file *os.File) string {
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		stringBuffer.WriteString(swapSyntax(scanner.Text()) + "\n")
	}
	return stringBuffer.String()
}

func swapSyntax(line string) string {
	variables := regexp.MustCompile("@")
	cssReserved := regexp.MustCompile("\\$(media|import|keyframes|-webkit|-moz|-o)")
	line = variables.ReplaceAllString(line, "$")
	reserves := cssReserved.FindAllStringSubmatchIndex(line, -1)
	if len(reserves) > 0 {
		for i, _ := range reserves {
			ampersandIdx := reserves[i][0]
			line = line[:ampersandIdx] + "@" + line[ampersandIdx+1:]
		}
	}
	mixInDeclation := regexp.MustCompile(".(.)+\\((.)*\\)\\s{")
	mixIns := mixInDeclation.FindAllStringSubmatchIndex(line, -1)
	emptyParens := regexp.MustCompile("\\(\\)")
	if len(mixIns) > 0 {
		for i, _ := range mixIns {
			idx := mixIns[i][0]
			line = line[:idx] + "@mixin " + strings.Trim(line[idx+1:], " ")
			line = emptyParens.ReplaceAllString(line, "")
		}
	}
	return line
}
