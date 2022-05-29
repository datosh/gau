package gau

type PkgNode struct {
	dependsOn    map[string]*PkgNode
	dependedOnBy map[string]*PkgNode
	pkgPath      string
}

func NewPkgNode(pkgPath string) *PkgNode {
	n := &PkgNode{
		dependsOn:    make(map[string]*PkgNode),
		dependedOnBy: make(map[string]*PkgNode),
		pkgPath:      pkgPath,
	}
	return n
}

func (p *PkgNode) DependOn(other *PkgNode) {
	p.dependsOn[other.pkgPath] = other
	other.dependedOnBy[p.pkgPath] = p
}

func (p *PkgNode) IsDependingOn(pkgName string) bool {
	_, exists := p.dependsOn[pkgName]
	return exists
}

func (p *PkgNode) IsIndirectlyDependingOn(pkgName string) bool {
	if p.IsDependingOn(pkgName) {
		return true
	}

	for _, dependingOn := range p.dependsOn {
		if dependingOn.IsIndirectlyDependingOn(pkgName) {
			return true
		}
	}
	return false
}

func (p *PkgNode) IsDependedOnBy(pkgName string) bool {
	_, exists := p.dependedOnBy[pkgName]
	return exists
}
