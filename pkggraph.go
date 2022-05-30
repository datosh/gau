package gau

import "golang.org/x/tools/go/packages"

type pkgGraph struct {
	lut   map[string]*pkgNode
	roots []*pkgNode
}

func newPkgGraph() *pkgGraph {
	g := &pkgGraph{}
	g.lut = make(map[string]*pkgNode)
	return g
}

func (p *pkgGraph) addNode(pkgPath string) {
	if _, exists := p.lut[pkgPath]; !exists {
		p.lut[pkgPath] = newPkgNode(pkgPath)
	}
}

func (p *pkgGraph) getNode(pkgPath string) *pkgNode {
	if node, exists := p.lut[pkgPath]; exists {
		return node
	} else {
		return nil
	}
}

func (p *pkgGraph) size() int {
	return len(p.lut)
}

func (p *pkgGraph) load(pkg string) error {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedImports,
	}

	pkgs, err := packages.Load(cfg, pkg)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		p.addNode(pkg.PkgPath)
		for _, dependsOn := range pkg.Imports {
			p.addNode(dependsOn.PkgPath)
			p.getNode(pkg.PkgPath).dependOn(p.getNode(dependsOn.PkgPath))
		}
	}

	p.updateRoots()
	return nil
}

func (p *pkgGraph) updateRoots() {
	p.roots = make([]*pkgNode, 0)

	for _, node := range p.lut {
		if len(node.dependedOnBy) == 0 {
			p.roots = append(p.roots, node)
		}
	}
}
