# go-less-to-sass
simple go command line tool to quickly convert your less projects to sass
(this project is currently under development.) 

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




