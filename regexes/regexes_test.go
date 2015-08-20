/*unit tests for all expressions used to match and convert less syntax*/

package regexes

import (
	"testing"
)

func TestAt(t *testing.T) {
	cases := []string{
		"test@testing.com",
		"@ for you",
		"hello i am @ this place",
	}

	for _, c := range cases {
		matched := At.MatchString(c)
		if !matched {
			t.Errorf("Expression to match @ symbal failed when matching %s", c)
		}
	}
}

func TestLessNameSpace(t *testing.T) {
	cases := []struct {
		in          string
		shouldMatch bool
	}{
		{"#Namespace {", true},
		{"#Gradients()", false},
		{"Filters {", false},
	}

	for _, c := range cases {
		matched := LessNameSpace.MatchString(c.in)
		if matched != c.shouldMatch {
			t.Errorf("LessNameSpace returned %v. Expected %v when matching -> %s", matched, c.shouldMatch, c.in)
		}
	}
}

func TestMixInDeclation(t *testing.T) {
	cases := []struct {
		in          string
		shouldMatch bool
	}{
		{".my-mixin(@color) {", true},
		{"anotherMixin{blah]", false},
		{".mixit(@width, @color);", false},
		{".mixit(@width, @color) {", true},
		{".nospace-beforebracket(@width, @color){", true},
		{".space-before-parenthesis (){", true},
	}

	for _, c := range cases {
		matched := MixInDeclation.MatchString(c.in)
		if matched != c.shouldMatch {
			t.Errorf("MixInDeclaration returned %v. Expected %v when matching -> %s", matched, c.shouldMatch, c.in)
		}
	}
}

func TestLessMixin(t *testing.T) {
	cases := []struct {
		in          string
		shouldMatch bool
	}{
		{"color: blue; .myMixin(@funvar);", true},
		{".myClass { .my-Mixin(@funvar, @color)}", true},
		{"#cool { .box-shadow(1px 2px 4px rgba(0, 0, 0, 0.4); }", true},
		{".my-mixin(@color) {", false},
		{".no-parentheses;", true},
		{"width: .05em;", false},
	}
	for _, c := range cases {
		matched := LessMixin.MatchString(c.in)
		if matched != c.shouldMatch {
			t.Errorf("LessMixin returned %v. Expected %v when matching -> %s", matched, c.shouldMatch, c.in)
		}
	}
}
