package converter

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/syllabix/go-less-to-sass/regexes"
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
	line = handleLessNamespaces(line)
	line = swapMixins(line)
	return line
}

func swapVars(line string) string {
	variables := regexp.MustCompile("@")
	line = variables.ReplaceAllString(line, "$")
	reserves := regexes.CssReservedWords.FindAllStringSubmatchIndex(line, -1)
	if len(reserves) > 0 {
		for i, _ := range reserves {
			ampersandIdx := reserves[i][0]
			line = line[:ampersandIdx] + "@" + line[ampersandIdx+1:]
		}
	}
	return line
}

func handleLessNamespaces(line string) string {
	nameSpaces := regexes.LessNameSpace.FindAllString(line, -1)
	if nameSpaces != nil {
		foundNameSpaces = append(foundNameSpaces, strings.Join(nameSpaces, ", "))
	}
	if len(foundNameSpaces) > 0 {
		if regexes.OpenCurly.MatchString(line) {
			nsCurlyCount++
		}
		if regexes.ClosedCurly.MatchString(line) {
			nsCurlyCount--
		}
	}
	if nsCurlyCount == 0 {
		foundNameSpaces = append(make([]string, 0))
	}
	nsMixInIdx := regexes.NamespacedMixins.FindAllStringSubmatchIndex(line, -1)
	nsMixIns := regexes.NamespacedMixins.FindAllString(line, -1)
	if len(nsMixIns) > 0 {
		for i, _ := range nsMixIns {
			fIdx := nsMixInIdx[i][0]
			lIdx := nsMixInIdx[i][len(nsMixInIdx[i])-1]
			fmtName := regexes.HashAndDot.ReplaceAllString(nsMixIns[i], "")
			fmtName = regexes.GreaterThan.ReplaceAllString(fmtName, "-")
			fmtName = regexes.Space.ReplaceAllString(fmtName, "")
			fmt.Println(fmtName)
			line = line[:fIdx] + "@include " + fmtName + line[lIdx:]
		}
		fmt.Println(line)
	}
	return line
}

func swapMixins(line string) string {
	if !regexes.MixInDeclation.MatchString(line) {
		return line
	}
	mixIns := regexes.MixInDeclation.FindAllStringSubmatchIndex(line, -1)
	var mixin string
	if len(foundNameSpaces) > 0 && nsCurlyCount > 0 {
		mixin = strings.Join(foundNameSpaces, "-")
		mixin = regexes.Hashtag.ReplaceAllString(mixin, "")
		mixin = "@mixin " + mixin + "-"
	} else {
		mixin = "@mixin "
	}
	if len(mixIns) > 0 {
		for i, _ := range mixIns {
			idx := mixIns[i][0]
			line = line[:idx] + mixin + strings.Trim(line[idx+1:], " ")
			line = regexes.EmptyParens.ReplaceAllString(line, "")
			line = regexes.OffByOneMixinConcat.ReplaceAllString(line, "-")
		}
	}
	return line
}
