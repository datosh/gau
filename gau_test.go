package gau

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGau_DependOnA(t *testing.T) {
	mockT := &testing.T{}

	Packages(mockT, "github.com/datosh/gau/tests/...").That().
		ResideIn("github.com/datosh/gau/tests/dependona").
		ShouldNot().DirectlyDependOn("github.com/datosh/gau/tests/a")

	assert.True(t, mockT.Failed())
}

func TestGau_IndirectOnA(t *testing.T) {
	mockT := &testing.T{}

	resideIn := Packages(mockT, "github.com/datosh/gau/tests/...").That().
		ResideIn("github.com/datosh/gau/tests/indirectona")

	resideIn.ShouldNot().DirectlyDependOn("github.com/datosh/gau/tests/a")
	assert.False(t, mockT.Failed())

	resideIn.ShouldNot().IndirectlyDependOn("github.com/datosh/gau/tests/a")
	assert.True(t, mockT.Failed())
}
