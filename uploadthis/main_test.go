package main

import (
	"code.google.com/p/gomock/gomock"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/conf" //mock
	"github.com/imosquera/uploadthis/util"
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

	calledCommandsForMon := false
	defer util.Patch(&executeCommandsForMonitors, func() {
		calledCommandsForMon = true
	}).Restore()

	main()

	c.Assert(calledCommandsForMon, Equals, true)
}

func (s *MainSuite) TestExecuteCommandsForMonitors(c *C) {
	//this test is getting pretty big already, we should consider breaking up the main function into smaller
	//testable functions
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	conf.Settings.MonitorDirs = []conf.MonitorDir{
		conf.MonitorDir{Path: "mockpath"},
	}

	calledCommandList := false
	defer util.Patch(&createCommandList, func(monitorDir *conf.MonitorDir) map[string]commands.Commander {
		calledCommandList = true
		return nil
	}).Restore()

	executeCommandsForMonitors()
	c.Assert(calledCommandList, Equals, true)
}
