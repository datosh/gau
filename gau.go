package gau

import "testing"

type gau struct {
	graph      *PkgGraph
	t          *testing.T
	toBeLoaded []string

	should bool

	resideIn string
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
	g.resideIn = pkg
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
	if g.graph.GetNode(g.resideIn).IsDependingOn(pkg) {
		if g.should {
			return
		}
		g.t.Fail()
	}
}

func (g *gau) IndirectlyDependOn(pkg string) {
	if g.graph.GetNode(g.resideIn).IsIndirectlyDependingOn(pkg) {
		if g.should {
			return
		}
		g.t.Fail()
	}
}
