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
var LessMixin = regexp.MustCompile("\\.(\\D)*\\((\\$|\\w|\\d|,|\\s|-|_|\\.|rgba\\((.)*\\))*\\);")
var ScssInterpolatedValue = regexp.MustCompile("\\${(\\w|-|_)*}")
var DollarBracket = regexp.MustCompile("\\${")

var Tilde = regexp.MustCompile("~")
var TildeStringEscape = regexp.MustCompile(":(\\s)?~\\\"(.)*\\\"")
var RubyStringInterpolation = regexp.MustCompile("#{")
var LessEStringEscape = regexp.MustCompile(":(\\s)?e\\(")
var LessEscape = regexp.MustCompile("e\\({1}?")
var LessArgb = regexp.MustCompile("argb\\((@|\\w|\\d|-|_)*(\\)){1}?")
var ArgbDeclaration = regexp.MustCompile("argb\\(")
var ClosedPerenWithSemiColon = regexp.MustCompile("\\);")
var ClosedPeren = regexp.MustCompile("\\)")
var OpenPeren = regexp.MustCompile("\\(")
var LessStringFormat = regexp.MustCompile("%{1}?\\((.)*\\)")
var StringPlaceHolder = regexp.MustCompile("('|\\\")%(s|S|d|D|a|A)('|\\\")")
var StringReplaceArguments = regexp.MustCompile("\\$(\\w|\\d)[\\w\\d-_]*")
