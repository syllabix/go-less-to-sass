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
var foundNameSpaces []string
var nsCurlyCount int = 0

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
	line = swapVars(line)
	lessNameSpace := regexp.MustCompile("#(\\w|\\d|-|_)+\\s{")
	openCurly := regexp.MustCompile("{")
	closedCurly := regexp.MustCompile("}")
	nameSpaces := lessNameSpace.FindAllString(line, -1)
	if nameSpaces != nil {
		foundNameSpaces = append(foundNameSpaces, strings.Join(nameSpaces, ", "))
	}
	if len(foundNameSpaces) > 0 {
		if openCurly.MatchString(line) {
			nsCurlyCount++
		}
		if closedCurly.MatchString(line) {
			nsCurlyCount--
		}
	}
	if nsCurlyCount == 0 {
		foundNameSpaces = append(make([]string, 0))
	}
	line = swapMixins(line)
	return line
}

func swapVars(line string) string {
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
	return line
}

func swapMixins(line string) string {
	mixInDeclation := regexp.MustCompile(".(.)+\\((.)*\\)\\s{")
	if !mixInDeclation.MatchString(line) {
		return line
	}
	mixIns := mixInDeclation.FindAllStringSubmatchIndex(line, -1)
	emptyParens := regexp.MustCompile("\\(\\)")
	offByOneMixinConcat := regexp.MustCompile("-\\.")
	var mixin string
	if len(foundNameSpaces) > 0 && nsCurlyCount > 0 {
		mixin = strings.Join(foundNameSpaces, "-")
		hashtag := regexp.MustCompile("(#|{|\\s)")
		//period := regexp.MustCompile("\\.")
		mixin = hashtag.ReplaceAllString(mixin, "")
		//mixin = period.ReplaceAllString(mixin, "-")
		mixin = "@mixin " + mixin + "-"
	} else {
		mixin = "@mixin "
	}
	if len(mixIns) > 0 {
		for i, _ := range mixIns {
			idx := mixIns[i][0]
			line = line[:idx] + mixin + strings.Trim(line[idx+1:], " ")
			line = emptyParens.ReplaceAllString(line, "")
			line = offByOneMixinConcat.ReplaceAllString(line, "-")
		}
	}
	return line
}
