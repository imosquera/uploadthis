package conf

import (
	"github.com/imosquera/uploadthis/util"
	. "launchpad.net/gocheck"
	"path"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

//this test makes sure that the access key and secret keys are set
func (s *MySuite) TestAuthSet(c *C) {
	optsParser = func(interface{}) ([]string, error) {
		opts.AccesssKey = "MOCK KEY"
		opts.SecretKey = "MOCK SECRET"
		opts.ConfigPath = "MOCK PATH"
		return []string{}, nil
	}

	var isLoadConfigCalled bool
	defer util.Patch(&loadConfig, func(path string) {
		isLoadConfigCalled = true
		c.Check(path, Equals, "MOCK PATH")
	}).Restore()

	ParseOpts()

	c.Check(Settings.Auth.AccessKey, Equals, "MOCK KEY")
	c.Check(Settings.Auth.AccessKey, Equals, "MOCK KEY")
	c.Check(isLoadConfigCalled, Equals, true)
}

//this test makes sure that are NOT set if one key is missing
//and if it's not it wont set the keys on the global settings packages
func (s *MySuite) TestAuthNotSet(c *C) {
	optsParser = func(interface{}) ([]string, error) {
		opts.AccesssKey = "MOCK KEY"
		opts.SecretKey = ""
		return []string{}, nil
	}
	ParseOpts()
	c.Check(Settings.Auth.AccessKey, Equals, "")
	c.Check(Settings.Auth.SecretKey, Equals, "")
}

func (s *MySuite) TestLoadConfig(c *C) {
	sampleConfigPath := path.Join(util.RootProjectPath, "fixtures/sample-config.yaml")
	loadConfig(sampleConfigPath)
	c.Check(Settings.Auth.AccessKey, Equals, "myaccesskey")
	c.Check(Settings.Auth.SecretKey, Equals, "mysupersecretkey")
}
