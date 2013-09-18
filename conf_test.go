package uploadthis

import (
	"launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { gocheck.TestingT(t) }

type MySuite struct{}

var _ = gocheck.Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *gocheck.C) {
	optsParser = func(interface{}) ([]string, error) {
		opts.AccesssKey = "MOCK KEY"
		opts.SecretKey = "MOCK SECRET"
		return []string{}, nil
	}
	ParseOpts()
	c.Check(opts.AccesssKey, gocheck.Equals, "MOCK KEY")
	c.Check(opts.SecretKey, gocheck.Equals, "MOCK SECRET")
}
