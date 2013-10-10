package hooks

import (
	"code.google.com/p/gomock/gomock"
	mockgz "compress/gzip" //mock
	"compress/gzip"
	"github.com/imosquera/uploadthis/util" //mock
	"github.com/imosquera/uploadthis/util/mocks"
	"io" //mock
	. "launchpad.net/gocheck"
	"os"
	"path"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func TestPrehooks(t *testing.T) { TestingT(t) }

type PreHookSuite struct{}

var _ = Suite(&PreHookSuite{})

func (s *PreHookSuite) SetupTest(c *C) {
}

func (s *PreHookSuite) TearDownTest(c *C) {
}

func (s *PreHookSuite) TestGzipCompressFile(c *C) {
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

func (s *PreHookSuite) TestRunCompressPrehook(c *C) {
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
