package hooks

import (
	"code.google.com/p/gomock/gomock"
	mockgz "compress/gzip" //mock
	"compress/gzip"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/util" //mock
	"github.com/imosquera/uploadthis/util/mocks"
	"io" //mock
	. "launchpad.net/gocheck"
	"os"
	"path"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type HookSuite struct{}

var _ = Suite(&HookSuite{})

func (s *HookSuite) SetupTest(c *C) {
}

func (s *HookSuite) TearDownTest(c *C) {
}

func (s *HookSuite) TestGzipCompressFile(c *C) {
	//I THINK SHOULD BE CHANGED TO TEST LEST OF IMPLEMNTATION AND MORE OF AN INTEGRATION TEST
	//setup a mock controller
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	util.MOCK().SetController(mockCtrl)
	util.EXPECT().LogPanic(nil)
	util.EXPECT().LogPanic(nil)
	util.EXPECT().LogPanic(nil)

	mockgz.MOCK().SetController(mockCtrl)
	io.MOCK().SetController(mockCtrl)
	//we must open a real file in order to send it back.
	mockFile, _ := os.Open(path.Join(util.RootProjectPath, "fixtures", "gziptest.txt"))

	mockOS := util.OsFs{}

	mockOS.EXPECT().Open("mockpath").Return(mockFile, nil)
	mockOS.EXPECT().Create("mockpath.gz").Return(mockFile, nil)

	mockGzipWriter := &mockgz.Writer{}
	mockgz.EXPECT().NewWriterLevel(mockFile, gzip.BestCompression).Return(mockGzipWriter, nil)
	mockGzipWriter.EXPECT().Close()

	var mockInt int64 = 0
	io.EXPECT().Copy(mockGzipWriter, mockFile).Return(mockInt, nil)

	util.Fs = mockOS

	gzipCompressor := GzipFileCompressor{}

	gzipCompressor.Compress("mockpath")

}

func (s *HookSuite) TestRunCompressPrehook(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	mockCompressor := mocks.NewMockCompressor(mockCtrl)
	uploadFile := "mockpath"
	uploadFiles := []string{uploadFile}
	mockCompressor.EXPECT().Compress(uploadFile).Return("newmockpath", nil)
	compressionHook := NewCompressPrehook()
	compressionHook.compressor = mockCompressor
	compressionHook.SetUploadFiles(uploadFiles)
	println(&uploadFiles)
	newUploadFiles, _ := compressionHook.Run()
	c.Assert(len(newUploadFiles), Equals, 1)
	c.Assert(newUploadFiles[0], Equals, "newmockpath")
}

func (s *HookSuite) TestGetPrehooks(c *C) {

	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	prehooker := mocks.NewMockCommander(mockCtrl)
	mockPrehooks := []string{"mock_prehook"}
	prehookMap := make(map[string]commands.Commander, 0)
	RegisterPrehook("mock_prehook", prehooker)

	GetPrehookCommands(mockPrehooks, prehookMap)
	_, ok := prehookMap["mock_prehook"]

	c.Assert(len(prehookMap), Equals, 1)
	c.Assert(ok, Equals, true)
}

func (s *HookSuite) TestRegisterCompressHook(c *C) {
	prehook := &CompressPrehook{}
	RegisterPrehook("mock_name", prehook)
	c.Assert(registeredPrehooks["mock_name"], Equals, prehook)
}
