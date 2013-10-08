package hooks

import (
	"code.google.com/p/gomock/gomock"
	"compress/gzip"
	"github.com/imosquera/uploadthis/monitor"
	"github.com/imosquera/uploadthis/util"
	"github.com/imosquera/uploadthis/util/mocks"
	"io/ioutil"
	. "launchpad.net/gocheck"
	"os"
	"path"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type HookSuite struct {
	mockCtrl *gomock.Controller
}

var _ = Suite(&HookSuite{})

func (s *HookSuite) SetupTest(c *C) {
	s.mockCtrl = gomock.NewController(c)
}

func (s *HookSuite) TearDownTest(c *C) {
	s.mockCtrl = nil
}

func (s *HookSuite) TestDoingPrehook(c *C) {
	scratchDir := path.Join(util.RootProjectPath, "tmp/fixtures/prehook_test")
	doingDir := path.Join(scratchDir, "doing")
	//cleanup previous test
	os.RemoveAll(scratchDir)
	os.MkdirAll(doingDir, 0755)

	mockFilePath := path.Join(scratchDir, "testfile.txt")
	os.Create(mockFilePath)
	mockFileStat, _ := os.Stat(mockFilePath)

	mockUploadFiles := []monitor.UploadFileInfo{
		monitor.UploadFileInfo{Path: mockFilePath, Info: mockFileStat},
	}

	prehooker := DoingPrehook{}
	prehooker.RunPrehook(mockUploadFiles)

	filePath, _ := os.Stat(path.Join(doingDir, "testfile.txt"))

	//monitor := monitor.UploadFileInfo{}
	//monitor.Path()
}
func (s *HookSuite) TestCompressFile(c *C) {
	sampleTxtPath := path.Join(util.RootProjectPath, "fixtures/monitordir/sample.txt")
	sampleFile, _ := os.Open(sampleTxtPath)
	originalBytes, _ := ioutil.ReadAll(sampleFile)
	info, _ := sampleFile.Stat()
	uploadFileInfo := monitor.UploadFileInfo{Path: sampleTxtPath, Info: info}
	gzipUploadFileInfo, _ := compressFile(uploadFileInfo)

	//lets read the gzipped file
	gzipFile, _ := os.Open(gzipUploadFileInfo.Path)
	gzipReader, _ := gzip.NewReader(gzipFile)
	unzippedBytes, _ := ioutil.ReadAll(gzipReader)

	c.Assert(string(originalBytes), Equals, string(unzippedBytes))
}

func (s *HookSuite) TestRunPrehook(c *C) {
	mockFileInfo := mocks.NewMockFileInfo(s.mockCtrl)

	mockUploadFileInfo := monitor.UploadFileInfo{Path: "MOCKPATH", Info: mockFileInfo}
	compressFileCalled := false
	defer util.Patch(&compressFile, func(uploadFile monitor.UploadFileInfo) (monitor.UploadFileInfo, error) {
		compressFileCalled = true
		c.Assert(uploadFile, Equals, mockUploadFileInfo)
		return monitor.UploadFileInfo{Path: "MOCKPATH", Info: mockFileInfo}, nil
	}).Restore()
	compress := CompressPrehook{}
	mockUploads := []monitor.UploadFileInfo{mockUploadFileInfo}
	compress.RunPrehook(mockUploads)
	c.Assert(compressFileCalled, Equals, true)
}

func (s *HookSuite) TestGetPrehooks(c *C) {
	prehooker := mocks.NewMockPrehooker(s.mockCtrl)
	mockPrehooks := []string{"mock_preehook"}
	RegisterPrehook("mock_prehook", prehooker)
	prehookers := GetPrehooks(mockPrehooks)
	c.Assert(len(prehookers), Equals, 1)
}

func (s *HookSuite) TestRegisterCompressHook(c *C) {
	prehook := CompressPrehook{}
	RegisterPrehook("mock_name", prehook)
	c.Assert(registeredPrehooks["mock_name"], Equals, prehook)
}
