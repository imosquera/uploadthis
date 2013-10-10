package upload

import (
	"code.google.com/p/gomock/gomock"
	"fmt"
	"github.com/imosquera/uploadthis/util"
	"github.com/imosquera/uploadthis/util/mocks"
	"launchpad.net/goamz/s3" //mock
	. "launchpad.net/gocheck"
	"os"
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

	newMockFileInfo := mocks.NewMockFileInfo(mockCtrl)
	newMockFileInfo.EXPECT().ModTime().Return(modTime)

	prefix := GeneratePathPrefix(newMockFileInfo)

	dateFormat := formatDate(modTime)
	c.Assert(prefix, Equals, dateFormat)
}

func formatDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}
func (s *UploadSuite) TestS3Upload(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	s3.MOCK().SetController(mockCtrl)

	fixturePath := path.Join(util.SetupRootProjectPath(), "fixtures/mockfile.log")
	if file, err := os.Open(fixturePath); err == nil {
		fileInfo, _ := file.Stat()
		keyPath := path.Join(formatDate(fileInfo.ModTime()), path.Base(fixturePath))
		bucket := &s3.Bucket{}
		bucket.EXPECT().PutReader(keyPath, gomock.Any(), int64(79), "text/plain", s3.Private)
		s3Uploader := S3Uploader{bucket: bucket}
		s3Uploader.Upload(fixturePath)
	}

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
