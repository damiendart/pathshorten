// Copyright (C) Damien Dart, <damiendart@pobox.com>.
// This file is distributed under the MIT licence. For more information,
// please refer to the accompanying "LICENCE" file.

package pathshorten

import (
	"testing"
)

func TestPathShorten(t *testing.T) {
	t.Parallel()

	// The following tests cases are based on manual testing of the
	// "pathshorten" function in Vim 9.1.1591 and [Vim's unit tests].
	//
	// [Vim's unit tests]: https://github.com/vim/vim/blob/v9.1.1591/src/testdir/test_functions.vim#L517
	tests := []struct {
		input               string
		pathSeparator       string
		pathComponentLength uint
		want                string
	}{
		{"", "/", 1, ""},
		{".", "/", 1, "."},
		{"..", "/", 1, ".."},
		{"~", "/", 1, "~"},
		{"~~", "/", 1, "~~"},
		{"foo", "/", 1, "foo"},
		{"/foo", "/", 1, "/foo"},
		{"foo/", "/", 1, "f/"},
		{"foo/bar", "/", 1, "f/bar"},
		{"foo/bar/foobar", "/", 1, "f/b/foobar"},
		{".foo/bar", "/", 1, ".f/bar"},
		{"~foo/bar", "/", 1, "~f/bar"},
		{"~.foo/bar", "/", 1, "~.f/bar"},
		{".~foo/bar", "/", 1, ".~f/bar"},
		{"~/foo/bar", "/", 1, "~/f/bar"},
		{"~/föo/bar", "/", 1, "~/f/bar"},
		{"~/àéïöü/bar", "/", 1, "~/à/bar"},
		{"~/../bar", "/", 1, "~/../bar"},
		{"foo:bar", ":", 1, "f:bar"},
		{`C:\foo\bar`, `\`, 1, `C\f\bar`},

		{"", "/", 2, ""},
		{".", "/", 2, "."},
		{"..", "/", 2, ".."},
		{"~", "/", 2, "~"},
		{"~~", "/", 2, "~~"},
		{"foo", "/", 2, "foo"},
		{"/foo", "/", 2, "/foo"},
		{"foo/", "/", 2, "fo/"},
		{"foo/bar", "/", 2, "fo/bar"},
		{"foo/bar/foobar", "/", 2, "fo/ba/foobar"},
		{".foo/bar", "/", 2, ".fo/bar"},
		{"~foo/bar", "/", 2, "~fo/bar"},
		{"~.foo/bar", "/", 2, "~.fo/bar"},
		{".~foo/bar", "/", 2, ".~fo/bar"},
		{"~/foo/bar", "/", 2, "~/fo/bar"},
		{"~/föo/bar", "/", 2, "~/fö/bar"},
		{"~/àéïöü/bar", "/", 2, "~/àé/bar"},
		{"~/../bar", "/", 2, "~/../bar"},
		{"foo:bar", ":", 2, "fo:bar"},
		{`C:\foo\bar`, `\`, 2, `C:\fo\bar`},
	}

	for _, test := range tests {
		t.Run(
			test.input,
			func(t *testing.T) {
				t.Parallel()

				output := PathShorten(
					test.input,
					test.pathSeparator,
					test.pathComponentLength,
				)

				if output != test.want {
					t.Errorf(
						"PathShorten(%#v, %#v, %d) = %#v, want %#v",
						test.input,
						test.pathSeparator,
						test.pathComponentLength,
						output,
						test.want,
					)
				}
			},
		)
	}
}

func TestShortenPathComponent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input               string
		pathComponentLength uint
		want                string
	}{
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

	for _, test := range tests {
		t.Run(
			test.input,
			func(t *testing.T) {
				t.Parallel()

				output := shortenPathComponent(test.input, test.pathComponentLength)

				if output != test.want {
					t.Errorf(
						"shortenPathComponent(%#v, %d) = %#v, want %#v",
						test.input,
						test.pathComponentLength,
						output,
						test.want,
					)
				}
			},
		)
	}
}
