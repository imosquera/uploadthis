package upload

import (
	"code.google.com/p/gomock/gomock"
	"fmt"
	"github.com/imosquera/uploadthis/util"
	"github.com/imosquera/uploadthis/util/mocks"
	"launchpad.net/goamz/s3" //mock
	. "launchpad.net/gocheck"
	"path"
	"testing"
	"time"
)

// Hook up gocheck into the "go test" runner.
func TestMain(t *testing.T) { TestingT(t) }

type UploadSuite struct{}

var _ = Suite(&UploadSuite{})

func (s *UploadSuite) TestGeneratePathPrefix(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	modTime := time.Now()
	year, month, day := modTime.Date()

	newMockFileInfo := mocks.NewMockFileInfo(mockCtrl)
	newMockFileInfo.EXPECT().ModTime().Return(modTime)

	prefix := GeneratePathPrefix(newMockFileInfo)

	c.Assert(prefix, Equals, fmt.Sprintf("%04d-%02d-%02d", year, month, day))
}

func (s *UploadSuite) TestS3Upload(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	s3.MOCK().SetController(mockCtrl)

	fixturePath := path.Join(util.SetupRootProjectPath(), "fixtures/mockfile.log")

	bucket := &s3.Bucket{}
	bucket.EXPECT().PutReader("2013-10-09/mockfile.log", gomock.Any(), int64(79), "text/plain", s3.Private)

	s3Uploader := S3Uploader{bucket: bucket}
	s3Uploader.Upload(fixturePath)
}
func (s *UploadSuite) Run(c *C) {

	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()
	uploader := mocks.NewMockUploader(mockCtrl)
	uploader.EXPECT().Upload("mockpath")

	uploadCommand := UploadCommand{}
	uploadCommand.UploadFiles = []string{"mockpath"}
	uploadCommand.Run()
}

// func (self *UploadCommand) Run() ([]string, error) {
//     for _, file := range self.UploadFiles {
//         self.uploader.Upload(file)
//     }
//     return self.UploadFiles, nil
// }
