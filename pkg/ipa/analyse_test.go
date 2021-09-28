package ipa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBuildNumberFromIPA(t *testing.T) {
	ipaLocalPath := "./test.ipa"
	buildNumber, err := GetBuildNumberFromIPA(ipaLocalPath)
	assert.Nil(t, err)
	t.Logf("build number is->%s", buildNumber)
}
