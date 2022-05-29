package gau

import "golang.org/x/tools/go/packages"

type PkgGraph struct {
	lut   map[string]*PkgNode
	roots []*PkgNode
}

func NewPkgGraph() *PkgGraph {
	g := &PkgGraph{}
	g.lut = make(map[string]*PkgNode)
	return g
}

func (p *PkgGraph) addNode(pkgPath string) {
	if _, exists := p.lut[pkgPath]; !exists {
		p.lut[pkgPath] = NewPkgNode(pkgPath)
	}
}

func (p *PkgGraph) AddNode(pkgPath string) {
	p.addNode(pkgPath)
	p.updateRoots()
}

func (p *PkgGraph) GetNode(pkgPath string) *PkgNode {
	if node, exists := p.lut[pkgPath]; exists {
		return node
	} else {
		return nil
	}
}

func (p *PkgGraph) Size() int {
	return len(p.lut)
}

func (p *PkgGraph) Load(pkg string) error {
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
			p.GetNode(pkg.PkgPath).DependOn(p.GetNode(dependsOn.PkgPath))
		}
	}

	p.updateRoots()
	return nil
}

func (p *PkgGraph) Roots() []*PkgNode {
	return p.roots
}

func (p *PkgGraph) updateRoots() {
	p.roots = make([]*PkgNode, 0)

	for _, node := range p.lut {
		if len(node.dependedOnBy) == 0 {
			p.roots = append(p.roots, node)
		}
	}
}
