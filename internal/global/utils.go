package global

import (
	"path"
	"runtime"
	"strings"
)

// Package is a helper class for retrieving local package information.
type Package int

// Name returns the name of the package which invoked the mothod.
func (r Package) Name() string {
	var pkgPath string
	pc, _, _, _ := runtime.Caller(1)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	if parts[len(parts)-2][0] == '(' {
		pkgPath = strings.Join(parts[0:len(parts)-2], ".")
	} else {
		pkgPath = strings.Join(parts[0:len(parts)-1], ".")
	}
	return path.Base(pkgPath)
}
