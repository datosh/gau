package gau

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkgNode_DependOn(t *testing.T) {
	foo := NewPkgNode("foo")
	bar := NewPkgNode("bar")

	foo.DependOn(bar)

	assert.True(t, foo.IsDependingOn("bar"))
	assert.False(t, foo.IsDependedOnBy("bar"))

	assert.True(t, bar.IsDependedOnBy("foo"))
	assert.False(t, bar.IsDependingOn("foo"))
}

func TestPkgNode_IsIndirectlyDependingOn(t *testing.T) {
	// foo -> bar -> baz
	foo := NewPkgNode("foo")
	bar := NewPkgNode("bar")
	baz := NewPkgNode("baz")

	foo.DependOn(bar)
	bar.DependOn(baz)

	assert.True(t, foo.IsIndirectlyDependingOn("baz"))
}

func TestPkgNode_IsNotIndirectlyDependingOn(t *testing.T) {
	foo := NewPkgNode("foo")

	assert.False(t, foo.IsIndirectlyDependingOn("bar"))
}
