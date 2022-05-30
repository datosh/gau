package gau

import (
	"testing"

	"golang.org/x/tools/go/packages"
)

type gau struct {
	graph      *PkgGraph
	t          *testing.T
	toBeLoaded []string

	should bool

	resideIn []string
}

func Packages(t *testing.T, pkgs ...string) *gau {
	g := &gau{
		graph:      NewPkgGraph(),
		t:          t,
		toBeLoaded: pkgs,
	}
	return g
}

func (g *gau) That() *gau {
	for _, pkg := range g.toBeLoaded {
		g.graph.Load(pkg)
	}
	return g
}

func (g *gau) ResideIn(pkg string) *gau {
	g.resideIn = expand(pkg)
	return g
}

func (g *gau) Should() *gau {
	g.should = true
	return g
}

func (g *gau) ShouldNot() *gau {
	return g
}

func (g *gau) DirectlyDependOn(pkg string) {
	for _, resideIn := range g.resideIn {
		g.directlyDependOn(resideIn, pkg)
	}
}

func (g *gau) directlyDependOn(depender, dependee string) {
	if xor(g.isDirectlyDependOn(depender, dependee), g.should) {
		g.t.Fail()
	}
}

func (g *gau) isDirectlyDependOn(depender, dependee string) bool {
	return g.graph.GetNode(depender).IsDependingOn(dependee)
}

func (g *gau) IndirectlyDependOn(pkg string) {
	for _, resideIn := range g.resideIn {
		if g.graph.GetNode(resideIn).IsIndirectlyDependingOn(pkg) {
			if !g.should {
				g.t.Fail()
			}
		}
	}
}

func xor(a, b bool) bool {
	return a != b
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
