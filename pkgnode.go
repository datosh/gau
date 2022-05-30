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

func (n *pkgNode) dependOn(other *pkgNode) {
	n.dependsOn[other.pkgPath] = other
	other.dependedOnBy[n.pkgPath] = n
}

func (n *pkgNode) isDependingOn(pkgName string) bool {
	_, exists := n.dependsOn[pkgName]
	return exists
}

func (n *pkgNode) isIndirectlyDependingOn(pkgName string) bool {
	if n.isDependingOn(pkgName) {
		return true
	}

	for _, dependingOn := range n.dependsOn {
		if dependingOn.isIndirectlyDependingOn(pkgName) {
			return true
		}
	}
	return false
}

func (n *pkgNode) isDependedOnBy(pkgName string) bool {
	_, exists := n.dependedOnBy[pkgName]
	return exists
}
