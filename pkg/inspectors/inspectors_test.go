package inspectors_test

import (
	"github.com/proofpoint/kapprover/pkg/inspectors"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/proofpoint/kapprover/pkg/inspectors/group"
	_ "github.com/proofpoint/kapprover/pkg/inspectors/username"
)

func TestInspectors(t *testing.T) {
	var i inspectors.Inspectors
	assert := assert.New(t)

	actual := i.String()
	assert.Empty(actual, "default Inspectors.String()")
	assert.Len(i, 0, "default Inspectors")

	i.Set("group=system:serviceaccount")
	assert.Equal("group=system:serviceaccount", i.String(), "Inspectors.String()")
	assert.Len(i, 1, "Inspectors")
	assert.Equal("group", i[0].Name, "Inspectors[0].Name")

	i.Set("username")
	assert.Equal("group=system:serviceaccount,username", i.String(), "Inspectors.String()")
	assert.Len(i, 2, "Inspectors")
	assert.Equal("username", i[1].Name, "Inspectors[1].Name")

	i = inspectors.Inspectors{}
	i.Set("username,group")
	assert.Equal("username,group", i.String(), "Inspectors.String()")
	assert.Len(i, 2, "Inspectors")
	assert.Equal("username", i[0].Name, "Inspectors[0].Name")
	assert.Equal("group", i[1].Name, "Inspectors[1].Name")
}
