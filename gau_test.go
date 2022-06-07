package gau

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGau(t *testing.T) {
	testCases := map[string]struct {
		applier  func(*gau)
		wantFail bool
	}{
		"fail: dependona should not directly depend on a": {
			applier: func(g *gau) {
				g.ResideIn("github.com/datosh/gau/tests/dependona").
					ShouldNot().DirectlyDependOn("github.com/datosh/gau/tests/a")
			},
			wantFail: true,
		},
		"success: dependona should directly depend on a": {
			applier: func(g *gau) {
				g.ResideIn("github.com/datosh/gau/tests/dependona").
					Should().DirectlyDependOn("github.com/datosh/gau/tests/a")
			},
		},
		"fail: indirectona should not indirectly depend on a": {
			applier: func(g *gau) {
				g.ResideIn("github.com/datosh/gau/tests/indirectona").
					ShouldNot().IndirectlyDependOn("github.com/datosh/gau/tests/a")
			},
			wantFail: true,
		},
		"success: indirectona should indirectly depend on a": {
			applier: func(g *gau) {
				g.ResideIn("github.com/datosh/gau/tests/indirectona").
					Should().IndirectlyDependOn("github.com/datosh/gau/tests/a")
			},
		},
		"success: indirectona should not directly depend on a": {
			applier: func(g *gau) {
				g.ResideIn("github.com/datosh/gau/tests/indirectona").
					ShouldNot().DirectlyDependOn("github.com/datosh/gau/tests/a")
			},
		},
		"fail: indirectona should directly depend on a": {
			applier: func(g *gau) {
				g.ResideIn("github.com/datosh/gau/tests/indirectona").
					Should().DirectlyDependOn("github.com/datosh/gau/tests/a")
			},
			wantFail: true,
		},
		"success: no one should directly depend on nodependency": {
			applier: func(g *gau) {
				g.ResideIn("github.com/datosh/gau/tests/...").
					ShouldNot().DirectlyDependOn("github.com/datosh/gau/tests/nodependency")
			},
		},
		"success: no one except indirectona should directly depend on dependona": {
			applier: func(g *gau) {
				g.ResideIn("github.com/datosh/gau/tests/...").
					Except("github.com/datosh/gau/tests/indirectona").
					ShouldNot().DirectlyDependOn("github.com/datosh/gau/tests/dependona")
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockT := &testing.T{}

			pkgs := Packages(mockT, "github.com/datosh/gau/tests/...").That()
			tc.applier(pkgs)

			if tc.wantFail {
				assert.True(t, mockT.Failed())
			} else {
				assert.False(t, mockT.Failed())
			}
		})
	}
}

func Test_expand(t *testing.T) {
	testCases := map[string]struct {
		pkg      string
		wantPkgs []string
	}{
		"no expanse required": {
			"github.com/datosh/gau",
			[]string{"github.com/datosh/gau"},
		},
		"expand test": {
			"github.com/datosh/gau/tests/...",
			[]string{
				"github.com/datosh/gau/tests/a",
				"github.com/datosh/gau/tests/b",
				"github.com/datosh/gau/tests/dependona",
				"github.com/datosh/gau/tests/dependonaandb",
				"github.com/datosh/gau/tests/dependonb",
				"github.com/datosh/gau/tests/indirectona",
				"github.com/datosh/gau/tests/nodependency",
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.wantPkgs, expand(tc.pkg))
		})
	}
}

func TestSimple(t *testing.T) {
	Packages(t, "github.com/datosh/gau/tests/...").That().
		ResideIn("github.com/datosh/gau/tests/dependona").
		Should().DirectlyDependOn("github.com/datosh/gau/tests/a")
}
