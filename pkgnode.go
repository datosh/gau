package gau

type pkgNode struct {
	dependsOn    map[string]*pkgNode
	dependedOnBy map[string]*pkgNode
	pkgPath      string
}

func newPkgNode(pkgPath string) *pkgNode {
	n := &pkgNode{
		dependsOn:    make(map[string]*pkgNode),
		dependedOnBy: make(map[string]*pkgNode),
		pkgPath:      pkgPath,
	}
	return n
}

func (p *pkgNode) dependOn(other *pkgNode) {
	p.dependsOn[other.pkgPath] = other
	other.dependedOnBy[p.pkgPath] = p
}

func (p *pkgNode) isDependingOn(pkgName string) bool {
	_, exists := p.dependsOn[pkgName]
	return exists
}

func (p *pkgNode) isIndirectlyDependingOn(pkgName string) bool {
	if p.isDependingOn(pkgName) {
		return true
	}

	for _, dependingOn := range p.dependsOn {
		if dependingOn.isIndirectlyDependingOn(pkgName) {
			return true
		}
	}
	return false
}

func (p *pkgNode) isDependedOnBy(pkgName string) bool {
	_, exists := p.dependedOnBy[pkgName]
	return exists
}
