package hooks

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test2(t *testing.T) { TestingT(t) }

type MySuite2 struct{}

var _ = Suite(&MySuite2{})

//this test makes sure that the access key and secret keys are set
func (s *MySuite2) TestAuthSet(c *C) {
}

//this test makes sure that are NOT set if one key is missing
//and if it's not it wont set the keys on the global settings packages
func (s *MySuite2) TestAuthNotSet(c *C) {
}
