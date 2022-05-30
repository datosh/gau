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
