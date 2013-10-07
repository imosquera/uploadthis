package main

import (
	"code.google.com/p/gomock/gomock"
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/hooks"
	"github.com/imosquera/uploadthis/monitor"
	"github.com/imosquera/uploadthis/util"
	"github.com/imosquera/uploadthis/util/mocks"
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

	mockMonitorDir := conf.MonitorDir{}
	mockMonitorDir.Path = "mockpath"
	mockUploadFileInfos := []monitor.UploadFileInfo{}

	//MOCK PARSE OPTS
	calledParseOpt := false
	defer util.Patch(&conf.ParseOpts, func() {
		calledParseOpt = true
		//in this parse opts we'll also need to configure the monitordirs because
		//the monitorDirs struct is private
		conf.Settings.MonitorDirs = []conf.MonitorDir{mockMonitorDir}
	}).Restore()

	//MOCK MAKE WORK DIRS
	calledMakeWorkDirs := false
	defer util.Patch(&conf.MakeWorkDirs, func(dir string) {
		calledMakeWorkDirs = true
		c.Assert(dir, Equals, mockMonitorDir.Path)
	}).Restore()

	//MOCK GET UPLOADFILES
	calledGetUploadFiles := false
	defer util.Patch(&monitor.GetUploadFiles, func(dirPath string) []monitor.UploadFileInfo {
		calledGetUploadFiles = true
		c.Assert(dirPath, Equals, mockMonitorDir.Path)
		return mockUploadFileInfos
	}).Restore()

	//MOCK GET PREHOOKS
	calledPrehooks := false
	mockPrehook := mocks.NewMockPrehooker(mockCtrl)
	mockPrehook.EXPECT().RunPrehook(mockUploadFileInfos).Return(mockUploadFileInfos, nil)
	defer util.Patch(&hooks.GetPrehooks, func(prehooks []string) []hooks.Prehooker {
		calledPrehooks = true
		return []hooks.Prehooker{mockPrehook}
	}).Restore()

	//call main function and run test
	main()
	c.Assert(calledParseOpt, Equals, true)
	c.Assert(calledMakeWorkDirs, Equals, true)
	c.Assert(calledGetUploadFiles, Equals, true)
	c.Assert(calledPrehooks, Equals, true)
}
