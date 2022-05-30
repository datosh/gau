package gau

import (
	"testing"

	"golang.org/x/tools/go/packages"
)

type Gau struct {
	graph      *pkgGraph
	t          *testing.T
	toBeLoaded []string

	should bool

	except   []string
	resideIn []string
}

func Packages(t *testing.T, pkgs ...string) *Gau {
	g := &Gau{
		graph:      newPkgGraph(),
		t:          t,
		toBeLoaded: pkgs,
	}
	return g
}

func (g *Gau) That() *Gau {
	for _, pkg := range g.toBeLoaded {
		g.graph.load(pkg)
	}
	return g
}

func (g *Gau) ResideIn(pkg string) *Gau {
	g.resideIn = expand(pkg)
	return g
}

func (g *Gau) Should() *Gau {
	g.should = true
	return g
}

func (g *Gau) ShouldNot() *Gau {
	return g
}

func (g *Gau) Except(pkg string) *Gau {
	g.except = append(g.except, expand(pkg)...)
	return g
}

func (g *Gau) DirectlyDependOn(pkg string) {
	for _, resideIn := range g.resideIn {
		g.directlyDependOn(resideIn, pkg)
	}
}

func (g *Gau) directlyDependOn(depender, dependee string) {
	if g.inExcept(depender) {
		return
	}
	if g.isDirectlyDependOn(depender, dependee) != g.should {
		g.t.Fail()
	}
}

func (g *Gau) isDirectlyDependOn(depender, dependee string) bool {
	return g.graph.getNode(depender).isDependingOn(dependee)
}

func (g *Gau) IndirectlyDependOn(pkg string) {
	for _, resideIn := range g.resideIn {
		if g.graph.getNode(resideIn).isIndirectlyDependingOn(pkg) {
			if !g.should {
				g.t.Fail()
			}
		}
	}
}

func (g *Gau) inExcept(pkg string) bool {
	for _, exception := range g.except {
		if exception == pkg {
			return true
		}
	}
	return false
}

func expand(pkg string) []string {
	var result []string
	cfg := packages.Config{
		Mode: packages.NeedName,
	}
	pkgs, _ := packages.Load(&cfg, pkg)
	for _, pkg := range pkgs {
		result = append(result, pkg.PkgPath)
	}
	return result
}
