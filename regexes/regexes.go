package regexes

import (
	"regexp"
)

//keeping regular exp declarations out of the conversion logic
var At = regexp.MustCompile("@")
var LessNameSpace = regexp.MustCompile("#(\\w|\\d|-|_)+\\s{")
var OpenCurly = regexp.MustCompile("{")
var ClosedCurly = regexp.MustCompile("}")
var OneClosedCurly = regexp.MustCompile("^(.*?)}")
var NamespacedMixins = regexp.MustCompile("(#(.)*\\s>\\s)\\.(\\w|\\d|-|_)*")
var HashAndDot = regexp.MustCompile("(#|\\.)")
var GreaterThan = regexp.MustCompile(">")
var Space = regexp.MustCompile("\\s")

var CssReservedWords = regexp.MustCompile("\\$(media|import|keyframes|-webkit|-moz|-o|font-face|page|supports|document|charset)")
var MixInDeclation = regexp.MustCompile(".(.)+\\((.)*\\)\\s{")
var EmptyParens = regexp.MustCompile("\\(\\)")
var OffByOneMixinConcat = regexp.MustCompile("-\\.")
var Hashtag = regexp.MustCompile("(#|{|\\s)")

var ScssMixin = regexp.MustCompile("@mixin")
var LessMixin = regexp.MustCompile("\\.\\D((.)*\\((.)*\\)|(.)*);")

var Tilde = regexp.MustCompile("~")
var TildeStringEscape = regexp.MustCompile("~\\\"(.)*\\\"")
var RubyStringInterpolation = regexp.MustCompile("#{")
