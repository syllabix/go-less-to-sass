package converter

import (
	"bufio"
	"bytes"
	//	"fmt"
	"github.com/syllabix/go-less-to-sass/regexes"
	"os"
	"regexp"
	"strings"
)

type DataStream struct {
	Data string
	Err  error
}

type lessNameSpace struct {
	name       string
	curlyCount int
	verified   int //0 = just found; 1 = pending; 2 = false; 3 = true
}

var convertedFile string
var stringBuffer bytes.Buffer
var foundNameSpaces []lessNameSpace
var capturedNameSpaces []string

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
	if len(capturedNameSpaces) > 0 {
		return removeNameSpaces(stringBuffer.String())
	}
	return stringBuffer.String()
}

func swapSyntax(line string) string {
	line = swapVars(line)
	line = swapMixins(line)
	line = handleLessNamespaces(line)
	return line
}

func swapVars(line string) string {
	variables := regexp.MustCompile("@")
	line = variables.ReplaceAllLiteralString(line, "$")
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
		for _, nameSpace := range nameSpaces {
			ns := lessNameSpace{name: nameSpace, curlyCount: 0, verified: 0}
			foundNameSpaces = append(foundNameSpaces, ns)
		}
		capturedNameSpaces = append(capturedNameSpaces, strings.Join(nameSpaces, ", "))
	}
	line = trackNameSpaceClosures(line)
	nsMixInIdx := regexes.NamespacedMixins.FindAllStringSubmatchIndex(line, -1)
	nsMixIns := regexes.NamespacedMixins.FindAllString(line, -1)
	if len(nsMixIns) > 0 {
		for i, _ := range nsMixIns {
			fIdx := nsMixInIdx[i][0]
			lIdx := nsMixInIdx[i][len(nsMixInIdx[i])-1]
			fmtName := regexes.HashAndDot.ReplaceAllLiteralString(nsMixIns[i], "")
			fmtName = regexes.GreaterThan.ReplaceAllLiteralString(fmtName, "-")
			fmtName = regexes.Space.ReplaceAllLiteralString(fmtName, "")
			line = line[:fIdx] + "@include " + fmtName + line[lIdx:]
		}
	}
	if len(foundNameSpaces) > 0 {
		verifyNameSpaces(line)
	}
	if regexes.LessMixin.MatchString(line) {
		mixIns := regexes.LessMixin.FindAllStringSubmatchIndex(line, -1)
		for i, _ := range mixIns {
			idx := mixIns[i][0]
			line = line[:idx] + "@include " + line[idx+1:]
		}
	}
	return line
}

func trackNameSpaceClosures(line string) string {
	if len(foundNameSpaces) > 0 {
		if regexes.OpenCurly.MatchString(line) {
			for i, _ := range foundNameSpaces {
				foundNameSpaces[i].curlyCount++
			}
		}
		if regexes.ClosedCurly.MatchString(line) {
			var idxToRemove []int
			for i := 0; i < len(foundNameSpaces); i++ {
				foundNameSpaces[i].curlyCount--
				if foundNameSpaces[i].curlyCount == 0 {
					line = regexes.OneClosedCurly.ReplaceAllLiteralString(line, "")
					idxToRemove = append(idxToRemove, int(i))
				}
			}
			for _, idx := range idxToRemove {
				foundNameSpaces[idx] = foundNameSpaces[len(foundNameSpaces)-1]
				foundNameSpaces = foundNameSpaces[:len(foundNameSpaces)-1]
			}
		}
	}
	return line
}

func verifyNameSpaces(line string) {
	for i := 0; i < len(foundNameSpaces); i++ {
		switch foundNameSpaces[i].verified {
		case 3:
			break
		case 0:
			foundNameSpaces[i].verified = 2
		case 2:
			if regexes.LessNameSpace.MatchString(line) {
				foundNameSpaces[i].verified = 1
			}
			if regexes.ScssMixin.MatchString(line) {
				foundNameSpaces[i].verified = 3
				return
			}
			if foundNameSpaces[i].verified == 2 {
				currentNameSpaces := strings.Join(capturedNameSpaces, ",")
				nsToRemove := regexp.MustCompile(foundNameSpaces[i].name)
				currentNameSpaces = nsToRemove.ReplaceAllLiteralString(currentNameSpaces, "")
				capturedNameSpaces = strings.Split(currentNameSpaces, ",")
				foundNameSpaces[i] = foundNameSpaces[len(foundNameSpaces)-1]
				foundNameSpaces = foundNameSpaces[:len(foundNameSpaces)-1]
			}
		}
	}
}

func removeNameSpaces(filecontent string) string {
	var nsExp = "(" + strings.Join(capturedNameSpaces, "|") + ")"
	nsExp = regexes.Space.ReplaceAllLiteralString(nsExp, "\\s")
	var nsRegExp = regexp.MustCompile(nsExp)
	return nsRegExp.ReplaceAllLiteralString(filecontent, "")
}

func swapMixins(line string) string {
	if !regexes.MixInDeclation.MatchString(line) {
		return line
	}
	mixIns := regexes.MixInDeclation.FindAllStringSubmatchIndex(line, -1)
	var mixin string
	if len(foundNameSpaces) > 0 {
		var mixinNames []string
		for _, ns := range foundNameSpaces {
			mixinNames = append(mixinNames, ns.name)
		}
		mixin = strings.Join(mixinNames, "-")
		mixin = regexes.Hashtag.ReplaceAllLiteralString(mixin, "")
		mixin = "@mixin " + mixin + "-"
	} else {
		mixin = "@mixin "
	}
	if len(mixIns) > 0 {
		for i, _ := range mixIns {
			idx := mixIns[i][0]
			line = line[:idx] + mixin + strings.Trim(line[idx+1:], " ")
			line = regexes.EmptyParens.ReplaceAllLiteralString(line, "")
			line = regexes.OffByOneMixinConcat.ReplaceAllLiteralString(line, "-")
		}
	}
	return line
}
