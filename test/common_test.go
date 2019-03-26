package test

import (
	"testing"

	"github.com/lijo-jose/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestParseAttributes(t *testing.T) {
	line := "TEST-ID=\"Help\",URI=\"http://test\",ID=33\n"
	mapAttr := m3u8.ParseAttributes(line)

	assert.NotNil(t, mapAttr)
	assert.Equal(t, "Help", mapAttr["TEST-ID"])
	assert.Equal(t, "http://test", mapAttr["URI"])
	assert.Equal(t, "33", mapAttr["ID"])
}
