package converter

import (
	"testing"
)

func TestSwapVars(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{".myMixin(@color, @size", ".myMixin($color, $size"},
		{"@favColor: #333", "$favColor: #333"},
		{"row@{index}", "row${index}"},
	}

	for _, c := range cases {
		got := swapVars(c.in)
		if got != c.out {
			t.Errorf("swapVars should have converted %s to %s, but instead output %s", c.in, c.out, got)
		}
	}
}

func TestSwapMixins(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{".primary-text() {", "@mixin primary-text {"},
		{".border-radius(@radius: 3px) {", "@mixin border-radius(@radius: 3px) {"},
	}

	for _, c := range cases {
		got := swapMixins(c.in)
		if got != c.out {
			t.Errorf("swapMixins should have converted '%s' to '%s', but instead output '%s'", c.in, c.out, got)
		}
	}
}

func TestConvertStringMethods(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"filter: ~\"alpha(opacity=@{opacity})\";", "filter: \"alpha(opacity=#{$opacity})\";"},
	}

	for _, c := range cases {
		got := convertStringMethods(c.in)
		if got != c.out {
			t.Errorf("swapMixins should have converted '%s' to '%s', but instead output '%s'", c.in, c.out, got)
		}
	}
}

func TestConvertInterpolatedStrings(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"row${idx}", "row#{$idx}"},
		{"hello${world}", "hello#{$world}"},
	}

	for _, c := range cases {
		got := convertInterpolatedStrings(c.in)
		if got != c.out {
			t.Errorf("swapMixins should have converted '%s' to '%s', but instead output '%s'", c.in, c.out, got)
		}
	}
}

func TestHandleLessNamespaces(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{".transition(all 0.2s ease-in-out);", "@include transition(all 0.2s ease-in-out);"},
		{"	.transition(all 0.2s ease-in-out);", "	@include transition(all 0.2s ease-in-out);"},
		{"  .transition(all 0.2s ease-in-out);", "  @include transition(all 0.2s ease-in-out);"},
	}

	for _, c := range cases {
		got := handleLessNamespaces(c.in)
		if got != c.out {
			t.Errorf("handleLessNamespaces should have converted '%s' to '%s', but instead output '%s'", c.in, c.out, got)
		}
	}
}
