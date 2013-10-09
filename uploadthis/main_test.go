package main

import (
	"code.google.com/p/gomock/gomock"
	"github.com/imosquera/uploadthis/conf"      //mock
	"github.com/imosquera/uploadthis/execution" //mock
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func TestMain(t *testing.T) { TestingT(t) }

type MainSuite struct{}

var _ = Suite(&MainSuite{})

func (s *MainSuite) SetUpTest(c *C) {
}

func (s *MainSuite) TearDownTest(c *C) {
}

//this test makes sure that the access key and secret keys are set
func (s *MainSuite) TestMain(c *C) {
	//this test is getting pretty big already, we should consider breaking up the main function into smaller
	//testable functions
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	conf.MOCK().SetController(mockCtrl)
	conf.EXPECT().ParseOpts().Return()

	execution.MOCK().SetController(mockCtrl)

	mgr := execution.DefaultCommandManager{}
	mgr.EXPECT().ExecuteCommandsForMonitors()
	execution.EXPECT().NewDefaultCommandManager().Return(mgr)
	main()

}
