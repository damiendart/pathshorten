package pathshorten

import "testing"

// The following tests are based on Vim's "pathshorten" tests:
// <https://github.com/vim/vim/blob/master/src/testdir/test_functions.vim#L510>.
func TestPathShorten(t *testing.T) {
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
		if output := PathShorten(test[0]); output != test[1] {
			t.Errorf(
				"Input %#v: expected %#v, got %#v",
				test[0],
				test[1],
				output,
			)
		}
	}
}
