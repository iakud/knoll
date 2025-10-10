package codegen

import (
	"sort"
	"strings"
)

// importToGroup is a list of functions which map from an import path to
// a group number.
var importToGroup = []func(localPrefix, importPath string) (num int, ok bool){
	func(localPrefix, importPath string) (num int, ok bool) {
		if localPrefix == "" {
			return
		}
		for _, p := range strings.Split(localPrefix, ",") {
			if strings.HasPrefix(importPath, p) || strings.TrimSuffix(p, "/") == importPath {
				return 3, true
			}
		}
		return
	},
	func(_, importPath string) (num int, ok bool) {
		if strings.HasPrefix(importPath, "appengine") {
			return 2, true
		}
		return
	},
	func(_, importPath string) (num int, ok bool) {
		firstComponent := strings.Split(importPath, "/")[0]
		if strings.Contains(firstComponent, ".") {
			return 1, true
		}
		return
	},
}

func importGroup(localPrefix, importPath string) int {
	for _, fn := range importToGroup {
		if n, ok := fn(localPrefix, importPath); ok {
			return n
		}
	}
	return 0
}

type ImportSpec struct {
	Path         string
	Name         string
	SpacesBefore bool
}

func sortImports(localPrefix string, importSpecs []*ImportSpec) {
	sort.Sort(byImportSpec{localPrefix, importSpecs})
}

type byImportSpec struct {
	localPrefix string
	specs       []*ImportSpec
}

func (x byImportSpec) Len() int      { return len(x.specs) }
func (x byImportSpec) Swap(i, j int) { x.specs[i], x.specs[j] = x.specs[j], x.specs[i] }
func (x byImportSpec) Less(i, j int) bool {
	ipath := x.specs[i].Path
	jpath := x.specs[j].Path

	igroup := importGroup(x.localPrefix, ipath)
	jgroup := importGroup(x.localPrefix, jpath)
	if igroup != jgroup {
		return igroup < jgroup
	}

	if ipath != jpath {
		return ipath < jpath
	}
	return x.specs[i].Name < x.specs[j].Name
}
