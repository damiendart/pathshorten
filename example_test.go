// Copyright (C) Damien Dart, <damiendart@pobox.com>.
// This file is distributed under the MIT licence. For more information,
// please refer to the accompanying "LICENCE" file.

package pathshorten_test

import (
	"fmt"

	"github.com/damiendart/pathshorten"
)

func ExamplePathShorten() {
	fmt.Println(pathshorten.PathShorten("foo/bar/baz", "/", 2))
	fmt.Println(pathshorten.PathShorten("foo:bar:baz", ":", 2))
	fmt.Println(pathshorten.PathShorten("~/foo/bar/bazqux/", "/", 1))
	// Output:
	// fo/ba/baz
	// fo:ba:baz
	// ~/f/b/b/
}
