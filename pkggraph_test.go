package gau

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkgGraph_GetNode_NotAdded(t *testing.T) {
	graph := newPkgGraph()

	assert.Nil(t, graph.getNode("foo"))
}

func TestPkgGraph_Size(t *testing.T) {
	testCases := map[string]struct {
		path     string
		wantSize int
	}{
		"single package": {
			"github.com/datosh/gau/tests/a",
			1,
		},
		"single dependency": {
			"github.com/datosh/gau/tests/dependona",
			2,
		},
		"variadic path": {
			"github.com/datosh/gau/tests/...",
			7,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			graph := newPkgGraph()
			graph.load(tc.path)
			assert.Equal(t, tc.wantSize, graph.size())
		})
	}
}

func TestPkgGraph_IsDependingOn(t *testing.T) {
	assert := assert.New(t)
	graph := newPkgGraph()

	graph.load("github.com/datosh/gau/tests/dependona")

	assert.True(graph.getNode("github.com/datosh/gau/tests/dependona").
		isDependingOn("github.com/datosh/gau/tests/a"),
	)
}

func TestPkgGraph_IsDependedOnBy(t *testing.T) {
	assert := assert.New(t)
	graph := newPkgGraph()

	graph.load("github.com/datosh/gau/tests/dependona")

	assert.True(graph.getNode("github.com/datosh/gau/tests/a").
		isDependedOnBy("github.com/datosh/gau/tests/dependona"),
	)
}

func TestPkgGraph_Roots(t *testing.T) {
	assert := assert.New(t)
	graph := newPkgGraph()

	graph.load("github.com/datosh/gau/tests/a")

	assert.Len(graph.roots, 1)
}

func TestPkgGraph_RootsVariadic(t *testing.T) {
	assert := assert.New(t)
	graph := newPkgGraph()

	graph.load("github.com/datosh/gau/tests/...")

	assert.Len(graph.roots, 4)
}
