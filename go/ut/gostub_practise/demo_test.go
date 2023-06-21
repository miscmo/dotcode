package gostub_practise

import (
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTag(t *testing.T) {

	assert.Equal(t, "tag", GetTag())

	stubs := gostub.Stub(&Tag, "new_tag")
	defer stubs.Reset()

	assert.Equal(t, "new_tag", GetTag())
}

