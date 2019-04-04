// Package techconferences allows you to find your next tech conference using the Open-source and crowd-sourced conference website https://confs.tech/
package techconferences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConference(t *testing.T) {
	confs, err := GetConferences(DevOps, 2019)
	assert.NoError(t, err)
	assert.NotZero(t, len(confs))

	_, err = GetConferences(DevOps, 2010)
	assert.EqualError(t, err, "error getting conferences: HTTP Status Code 404")
}
