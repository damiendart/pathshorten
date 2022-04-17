package pathshorten

import "testing"

func TestPathShorten(t *testing.T) {
	// The following tests are based on Vim's "pathshorten" tests:
	// <https://github.com/vim/vim/blob/master/src/testdir/test_functions.vim#L510>.
	basicTests := [][]string{
		{"", ""},
		{".", "."},
		{"..", ".."},
		{"~", "~"},
		{"~~", "~~"},
		{"foo", "foo"},
		{"/foo", "/foo"},
		{"foo/", "f/"},
		{"foo/bar", "f/bar"},
		{"foo/bar/foobar", "f/b/foobar"},
		{".foo/bar", ".f/bar"},
		{"~foo/bar", "~f/bar"},
		{"~.foo/bar", "~.f/bar"},
		{".~foo/bar", ".~f/bar"},
		{"~/foo/bar", "~/f/bar"},
		{"~/föo/bar", "~/f/bar"},
		{"~/àéïöü/bar", "~/à/bar"},
	}

	for _, test := range basicTests {
		if output := PathShorten(test[0], 1); output != test[1] {
			t.Errorf(
				"Input %#v: expected %#v, got %#v",
				test[0],
				test[1],
				output,
			)
		}
	}
}

func TestShortenToken(t *testing.T) {
	basicTests := [][]string{
		{"", ""},
		{".", "."},
		{"..", ".."},
		{"~", "~"},
		{"~~", "~~"},
		{"a", "a"},
		{"ab", "a"},
		{"foo", "f"},
		{"foobar", "f"},
		{".a", ".a"},
		{"~a", "~a"},
		{".~a", ".~a"},
		{"~.a", "~.a"},
		{"~.foo", "~.f"},
		{"~~.foo", "~~.f"},
	}

	for _, test := range basicTests {
		if output := shortenToken(test[0], 1); output != test[1] {
			t.Errorf(
				"Input %#v: expected %#v, got %#v",
				test[0],
				test[1],
				output,
			)
		}
	}

	trimmingTests := [][]string{
		{"", ""},
		{".", "."},
		{"..", ".."},
		{"~", "~"},
		{"~~", "~~"},
		{"a", "a"},
		{"ab", "ab"},
		{"foo", "fo"},
		{"foobar", "fo"},
		{".a", ".a"},
		{"~a", "~a"},
		{".~a", ".~a"},
		{"~.a", "~.a"},
		{"~.foo", "~.fo"},
		{"~~.foo", "~~.fo"},
	}

	for _, test := range trimmingTests {
		if output := shortenToken(test[0], 2); output != test[1] {
			t.Errorf(
				"Input %#v: expected %#v, got %#v",
				test[0],
				test[1],
				output,
			)
		}
	}
}
