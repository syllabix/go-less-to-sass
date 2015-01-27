# go-less-to-sass
simple go command line tool to quickly convert your less projects to sass

# usage 
the tool in its current form will take a .less file as input, convert the syntax to .scss, and output a .scss file. most of the conversion is what you might expect, but as there are some features of LESS that are not availbale in SCSS (mainly, namespaces). The expected output is outlined below.


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
# LESS Namespaces
As scss does not support namespacing - when finding a less namespace declaration, the name will be used as a prefix delimited by a "-" for all of its respective encapsulated mixins. The declaration itself will be removed in the scss output. A lot simpler than it sounds... 
Input
```less
#font {
  #family {
    .serif() {
      font-family: @serifFontFamily;
    }
    .sans-serif() {
      font-family: @sansFontFamily;
    }
    .monospace() {
      font-family: @monoFontFamily;
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
  .sans-serif(@size: @baseFontSize, @weight: normal, @lineHeight: @baseLineHeight) {
    #font > #family > .sans-serif;
    #font > .shorthand(@size, @weight, @lineHeight);
  }
  .monospace(@size: @baseFontSize, @weight: normal, @lineHeight: @baseLineHeight) {
    #font > #family > .monospace;
    #font > .shorthand(@size, @weight, @lineHeight);
  }
}
```
Output
```scss
@mixin font-family-serif {
      font-family: $serifFontFamily;
    }
@mixin font-family-sans-serif {
      font-family: $sansFontFamily;
    }
@mixin font-family-monospace {
      font-family: $monoFontFamily;
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
@mixin font-sans-serif($size: $baseFontSize, $weight: normal, $lineHeight: $baseLineHeight) {
    @include font-family-sans-serif;
    @include font-shorthand($size, $weight, $lineHeight);
  }
@mixin font-monospace($size: $baseFontSize, $weight: normal, $lineHeight: $baseLineHeight) {
    @include font-family-monospace;
    @include font-shorthand($size, $weight, $lineHeight);
  }
```

*** This project is under active development. 






