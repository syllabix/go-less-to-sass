package regexes

import (
	"regexp"
)

//keeping regular exp declarations out of the conversion logic
var LessNameSpace = regexp.MustCompile("#(\\w|\\d|-|_)+\\s{")
var OpenCurly = regexp.MustCompile("{")
var ClosedCurly = regexp.MustCompile("}")

var CssReservedWords = regexp.MustCompile("\\$(media|import|keyframes|-webkit|-moz|-o)")
var MixInDeclation = regexp.MustCompile(".(.)+\\((.)*\\)\\s{")
var EmptyParens = regexp.MustCompile("\\(\\)")
var OffByOneMixinConcat = regexp.MustCompile("-\\.")
var Hashtag = regexp.MustCompile("(#|{|\\s)")
