package delaunay

import (
	pkg "github.com/alexeydott/geom/cmp"
)

var cmp = pkg.HiCMP

var oldCmp = pkg.SetDefault(pkg.HiCMP)
