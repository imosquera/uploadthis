package monitor

import (
	"code.google.com/p/gomock/gomock"
	"github.com/imosquera/uploadthis/util/mocks"
	mockio "io/ioutil" //mock
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func TestInit(t *testing.T) { TestingT(t) }

type MonitorTest struct{}

var _ = Suite(&MonitorTest{})

//this test makes sure that the access key and secret keys are set
func (s *MonitorTest) TestGetUploadFiles(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	filePath := mocks.NewMockFileInfo(mockCtrl)
	filePath.EXPECT().IsDir().Return(false)
	filePath.EXPECT().Name().Return("mockfile")

	mockio.MOCK().SetController(mockCtrl)
	mockio.EXPECT().ReadDir("mockpath").Return([]os.FileInfo{filePath}, nil)

	fileInfos := GetUploadFiles("mockpath")
	thePath := ""
	for _, fileInfo := range fileInfos {
		thePath = fileInfo.Path
	}
	c.Assert(len(fileInfos), Equals, 1)
	c.Assert(thePath, Equals, "mockpath/mockfile")
}
