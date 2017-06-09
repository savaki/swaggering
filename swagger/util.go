package swagger

import (
	"fmt"
	"path/filepath"
	"strings"
)

func makeRef(name string) string {
	return fmt.Sprintf("#/definitions/%v", name)
}

type reflectType interface {
	PkgPath() string
	Name() string
}

func makeName(t reflectType) string {
	var name string
	pkgPath := t.PkgPath()
	if pkgPath == "" {
		name = t.Name()
	} else {
		name = filepath.Base(pkgPath) + t.Name()
	}
	return strings.Replace(name, "-", "_", -1)
}
