package gau

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkgGraph_AddNode(t *testing.T) {
	graph := NewPkgGraph()

	graph.AddNode("foo")

	assert.Equal(t, "foo", graph.GetNode("foo").pkgPath)
}

func TestPkgGraph_GetNode_NotAdded(t *testing.T) {
	graph := NewPkgGraph()

	assert.Nil(t, graph.GetNode("foo"))
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
			"github.com/datosh/gau/tests/dependOnA",
			2,
		},
		"variadic path": {
			"github.com/datosh/gau/tests/...",
			3,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			graph := NewPkgGraph()
			graph.Load(tc.path)
			assert.Equal(t, tc.wantSize, graph.Size())
		})
	}
}

func TestPkgGraph_IsDependingOn(t *testing.T) {
	assert := assert.New(t)
	graph := NewPkgGraph()

	graph.Load("github.com/datosh/gau/tests/dependOnA")

	assert.True(graph.GetNode("github.com/datosh/gau/tests/dependOnA").
		IsDependingOn("github.com/datosh/gau/tests/a"),
	)
}

func TestPkgGraph_IsDependedOnBy(t *testing.T) {
	assert := assert.New(t)
	graph := NewPkgGraph()

	graph.Load("github.com/datosh/gau/tests/dependOnA")

	assert.True(graph.GetNode("github.com/datosh/gau/tests/a").
		IsDependedOnBy("github.com/datosh/gau/tests/dependOnA"),
	)
}

func TestPkgGraph_Roots(t *testing.T) {
	assert := assert.New(t)
	graph := NewPkgGraph()

	graph.Load("github.com/datosh/gau/tests/a")

	assert.Len(graph.Roots(), 1)
}
