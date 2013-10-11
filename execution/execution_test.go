package execution

import (
	"code.google.com/p/gomock/gomock"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/hooks" //mock
	"github.com/imosquera/uploadthis/util"  //mock
	"github.com/imosquera/uploadthis/util/mocks"
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func TestMain(t *testing.T) { TestingT(t) }

type ExecutionSuite struct{}

var _ = Suite(&ExecutionSuite{})

func (s *ExecutionSuite) SetUpTest(c *C) {
}

func (s *ExecutionSuite) TearDownTest(c *C) {
}

func (s *ExecutionSuite) GetMockMonitorDirs() []conf.MonitorDir {
	dirs := []conf.MonitorDir{conf.MonitorDir{Path: "mockpath"}}
	return dirs
}

func (s *ExecutionSuite) TestExecuteCommands(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	mockUploadFiles := []string{"mockpaths"}
	util.MOCK().SetController(mockCtrl)
	util.EXPECT().GetFilesFromDir("mockpath").Return(mockUploadFiles)

	hooks.MOCK().SetController(mockCtrl)

	mockMonitor := conf.MonitorDir{Path: "mockpath"}

	mockCommand := mocks.NewMockCommander(mockCtrl)
	mockCommand.EXPECT().SetUploadFiles(mockUploadFiles)
	mockCommand.EXPECT().SetName("mock")
	mockCommand.EXPECT().Prepare()
	mockCommand.EXPECT().Run()

	mockCommands := map[string]commands.Commander{
		"mock": mockCommand,
	}
	//mockCommand.EXPECT().Run(mockUploadFiles).Return(mockUploadFiles, nil)
	commandExecutor := SequentialCommandExecutor{}
	commandExecutor.ExecuteCommands(mockCommands, mockMonitor)
}

func (s *ExecutionSuite) TestExecutionForMonitors(c *C) {
	//this test is getting pretty big already, we should consider breaking up the main function into smaller
	//testable functions
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	conf.Settings.MonitorDirs = s.GetMockMonitorDirs()

	mockProducer := mocks.NewMockCommandProducer(mockCtrl)

	hooks.MOCK().SetController(mockCtrl)

	mockCommand := mocks.NewMockCommander(mockCtrl)
	mockList := map[string]commands.Commander{
		"mock": mockCommand,
	}

	mockMonitorDir := s.GetMockMonitorDirs()[0]
	mockProducer.EXPECT().CreateCommandList(&mockMonitorDir).Return(mockList)

	mockExecutor := mocks.NewMockCommandExecutor(mockCtrl)
	mockExecutor.EXPECT().ExecuteCommands(mockList, mockMonitorDir)

	mgr := DefaultCommandManager{
		Executor: mockExecutor,
		Producer: mockProducer,
	}
	mgr.ExecuteCommandsForMonitors()
}
