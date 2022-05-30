package gau

import (
	"golang.org/x/tools/go/packages"
)

type pkgGraph struct {
	lut   map[string]*pkgNode
	roots []*pkgNode
}

func newPkgGraph() *pkgGraph {
	return &pkgGraph{
		lut: make(map[string]*pkgNode),
	}
}

func (g *pkgGraph) addNode(pkgPath string) {
	if _, exists := g.lut[pkgPath]; !exists {
		g.lut[pkgPath] = newPkgNode(pkgPath)
	}
}

func (g *pkgGraph) getNode(pkgPath string) *pkgNode {
	if node, exists := g.lut[pkgPath]; exists {
		return node
	}
	return nil
}

func (g *pkgGraph) size() int {
	return len(g.lut)
}

func (g *pkgGraph) load(pkg string) error {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedImports,
	}

	pkgs, err := packages.Load(cfg, pkg)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		g.addNode(pkg.PkgPath)
		for _, dependsOn := range pkg.Imports {
			g.addNode(dependsOn.PkgPath)
			g.getNode(pkg.PkgPath).dependOn(g.getNode(dependsOn.PkgPath))
		}
	}

	g.updateRoots()
	return nil
}

func (g *pkgGraph) updateRoots() {
	for _, node := range g.lut {
		if len(node.dependedOnBy) == 0 {
			g.roots = append(g.roots, node)
		}
	}
}
