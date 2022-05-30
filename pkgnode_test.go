package gau

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkgNode_DependOn(t *testing.T) {
	foo := newPkgNode("foo")
	bar := newPkgNode("bar")

	foo.dependOn(bar)

	assert.True(t, foo.isDependingOn("bar"))
	assert.False(t, foo.isDependedOnBy("bar"))

	assert.True(t, bar.isDependedOnBy("foo"))
	assert.False(t, bar.isDependingOn("foo"))
}

func TestPkgNode_IsIndirectlyDependingOn(t *testing.T) {
	// foo -> bar -> baz
	foo := newPkgNode("foo")
	bar := newPkgNode("bar")
	baz := newPkgNode("baz")

	foo.dependOn(bar)
	bar.dependOn(baz)

	assert.True(t, foo.isIndirectlyDependingOn("baz"))
}

func TestPkgNode_IsNotIndirectlyDependingOn(t *testing.T) {
	foo := newPkgNode("foo")

	assert.False(t, foo.isIndirectlyDependingOn("bar"))
}
