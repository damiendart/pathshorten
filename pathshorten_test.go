package pathshorten

import "testing"

type TestCase struct {
	input               string
	pathComponentLength uint
	expected            string
}

func TestPathShorten(t *testing.T) {
	// The following tests are based on Vim's "pathshorten" tests:
	// <https://github.com/vim/vim/blob/master/src/testdir/test_functions.vim#L510>.
	testCases := []TestCase{
		{"", 1, ""},
		{".", 1, "."},
		{"..", 1, ".."},
		{"~", 1, "~"},
		{"~~", 1, "~~"},
		{"foo", 1, "foo"},
		{"/foo", 1, "/foo"},
		{"foo/", 1, "f/"},
		{"foo/bar", 1, "f/bar"},
		{"foo/bar/foobar", 1, "f/b/foobar"},
		{".foo/bar", 1, ".f/bar"},
		{"~foo/bar", 1, "~f/bar"},
		{"~.foo/bar", 1, "~.f/bar"},
		{".~foo/bar", 1, ".~f/bar"},
		{"~/foo/bar", 1, "~/f/bar"},
		{"~/föo/bar", 1, "~/f/bar"},
		{"~/àéïöü/bar", 1, "~/à/bar"},
		{"~/../bar", 1, "~/../bar"},
	}

	for _, testCase := range testCases {
		output := PathShorten(testCase.input, testCase.pathComponentLength)

		if output != testCase.expected {
			t.Errorf(
				"Input %#v: expected %#v, got %#v",
				testCase.input,
				testCase.expected,
				output,
			)
		}
	}
}

func TestShortenPathComponent(t *testing.T) {
	testCases := []TestCase{
		{"", 1, ""},
		{"", 2, ""},
		{".", 1, "."},
		{".", 2, "."},
		{"..", 1, ".."},
		{"..", 2, ".."},
		{"..", 3, ".."},
		{"~", 1, "~"},
		{"~", 2, "~"},
		{"~~", 1, "~~"},
		{"~~", 2, "~~"},
		{"~~", 3, "~~"},
		{"a", 1, "a"},
		{"a", 2, "a"},
		{"ab", 1, "a"},
		{"ab", 2, "ab"},
		{"foo", 1, "f"},
		{"foo", 2, "fo"},
		{"foobar", 1, "f"},
		{"foobar", 2, "fo"},
		{"foobar", 3, "foo"},
		{".a", 1, ".a"},
		{".a", 2, ".a"},
		{"~a", 1, "~a"},
		{"~a", 2, "~a"},
		{".~a", 1, ".~a"},
		{".~a", 2, ".~a"},
		{"~.a", 1, "~.a"},
		{"~.a", 2, "~.a"},
		{"~.foo", 1, "~.f"},
		{"~.foo", 2, "~.fo"},
		{"~~.foo", 1, "~~.f"},
		{"~~.foo", 2, "~~.fo"},
	}

	for _, testCase := range testCases {
		output := shortenPathComponent(
			testCase.input,
			testCase.pathComponentLength,
		)

		if output != testCase.expected {
			t.Errorf(
				"Input %#v: expected %#v, got %#v",
				testCase.input,
				testCase.expected,
				output,
			)
		}
	}
}
