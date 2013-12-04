package conf

import (
	"code.google.com/p/gomock/gomock"
	"github.com/imosquera/uploadthis/util"
	"github.com/imosquera/uploadthis/util/mocks"
	"github.com/jessevdk/go-flags" //mock
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
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	flags.MOCK().SetController(mockCtrl)
	flags.EXPECT().Parse(&opts)

	opts.ConfigPath = "mockpath"
	opts.AccessKey = "MOCK KEY"
	opts.SecretKey = "MOCK SECRET"
	mockConfLoader := mocks.NewMockConfigLoader(mockCtrl)

	mockConfLoader.EXPECT().LoadConfig("mockpath", &Settings).Return()
	configLoader = mockConfLoader

	mockLogConfig := mocks.NewMockLoggerConfig(mockCtrl)
	mockLogConfig.EXPECT().ConfigLogger("")
	loggerConfig = mockLogConfig

	ParseOpts()

	c.Assert(Settings.Auth.AccessKey, Equals, "MOCK KEY")
	c.Assert(Settings.Auth.SecretKey, Equals, "MOCK SECRET")

	//make sure we clean up when we're done
	opts.ConfigPath = ""
	opts.AccessKey = ""
	opts.SecretKey = ""
}

func (s *MySuite) TestYamlLoadConfig(c *C) {
	yamlPath := path.Join(util.RootProjectPath, "fixtures/sample-config.yaml")
	yamlLoader := &YamlConfigLoader{}
	mockSettings := &UploadthisConfig{}
	yamlLoader.LoadConfig(yamlPath, mockSettings)
	c.Assert(mockSettings.Auth.AccessKey, Equals, "myaccesskey")
	c.Assert(mockSettings.Auth.SecretKey, Equals, "mysupersecretkey")
}
