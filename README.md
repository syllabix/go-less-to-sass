# go-less-to-sass
simple go command line tool to help quickly convert your less projects to sass. as it is basically an automated find and replace tool, certain oddities and more advanced LESS features will be up to you.

# usage 
Install Go. Clone the repository into your Go workspace and build the project. If your are unfamiliar with setting up your Go environment, [Golang.org](https://golang.org/doc/code.html) is super straight forward on getting up to speed. The tool can either convert an individual file or an entire project. In both cases, it will convert all *.less files and write a *.scss file with the converted syntax, named identically to the less file it converted. Commands are as follows:

```bash
go-less-to-sass -filename="path/to/your.less"
```
or
```bash
cd to/your/less/project
go-less-to-sass
```

Below is basic run down of of less-to-sass conversion.

# Variables:
Input:
```less
@myFavColor: #333;
```
Output:
```scss
$myFavColor: #333;
```

# Mixin Declarations
Input:
```less
.border-radius(@radius) {
	-webkit-border-radius: @radius;
	-moz-border-radius: @radius;
	-ms-border-radius: @radius;
	border-radius: @radius;
}
```
Output:
```scss
@mixin border-radius($radius) {
	-webkit-border-radius: $radius;
	-moz-border-radius: $radius;
	-ms-border-radius: $radius;
	border-radius: $radius;
}
```
# Mixins
Input
```less
.my-mixin(@width, @height) { //.... }
```
Output
```scss
@include my-mixin($width, $height) { //.... }
```

# LESS Namespaces
As scss does not support namespacing - when finding a less namespace declaration, the tool will attempt to use the namespace as a prefix delimited by a "-" for all of its respective encapsulated mixins. The declaration itself should be removed in the scss output. A lot simpler than it sounds... 
Input
```less
#font {
  #family {
    .serif() {
      font-family: @serifFontFamily;
    }
  }
  .shorthand(@size: @baseFontSize, @weight: normal, @lineHeight: @baseLineHeight) {
    font-size: @size;
    font-weight: @weight;
    line-height: @lineHeight;
  }
  .serif(@size: @baseFontSize, @weight: normal, @lineHeight: @baseLineHeight) {
    #font > #family > .serif;
    #font > .shorthand(@size, @weight, @lineHeight);
  }
}
```
Output
```scss
@mixin font-family-serif {
      font-family: $serifFontFamily;
    }

@mixin font-shorthand($size: $baseFontSize, $weight: normal, $lineHeight: $baseLineHeight) {
    font-size: $size;
    font-weight: $weight;
    line-height: $lineHeight;
  }
@mixin font-serif($size: $baseFontSize, $weight: normal, $lineHeight: $baseLineHeight) {
    @include font-family-serif;
    @include font-shorthand($size, $weight, $lineHeight);
  }
```

# Interpolated Strings
Input
```less
.grid@{index} { //styles }
```
Output
```scss
.gird#{$index} { //styles }
```

# Extend
The tool will currently only convert extend declarations that are nested. As scss does not support extend as a pseudo class, this will have to be cleaned up manually.
Input
```less
.my-style {
  &:extend(.their-style);
}
```
Output
```scss
.my-style {
  @ extend .their-style;
}
```
# String methods
The tool will convert LESS Printf style string methods with scss interpolated valued. It also (as it is included in this example) removes argb color methods as scss does not supporr them.

Input
```less
filter: e(%("progid:DXImageTransform.Microsoft.gradient(startColorstr='%d', endColorstr='%d', GradientType=1)",argb(@start-color),argb(@end-color)));
```

Output
```scss
filter: progid:DXImageTransform.Microsoft.gradient(startColorstr=#{$start-color}, endColorstr=#{$end-color}, GradientType=1);
```



*** Notes - this tool is meant to be thought of as an assistant - the real intention behind its development was to get more famliar with Go (amazing language!!) while making something potentially useful. As mentioned before, while being pretty effective, is not perfect. At minimum should save some time converting projects. If you find it useful - fork at will.




