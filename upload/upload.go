package upload

import (
	"fmt"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/conf"
	"github.com/imosquera/uploadthis/util"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"os"
	"path"
)

//this upload command is the base for all types of uploading to a destination
type UploadCommand struct {
	*commands.Command
	uploader Uploader
}

//the only constructor for generating a upload command
func NewUploadCommand(monitorDir conf.MonitorDir) *UploadCommand {
	var contentType string = "text/plain"
	for _, item := range monitorDir.PreHooks {
		if item == "compress" {
			contentType = "application/gzip"
		}
	}
	return &UploadCommand{
		commands.NewFileStateCommand(),
		NewS3Uploader(monitorDir.Bucket, contentType),
	}
}

//the run command for the stand upload command
func (self *UploadCommand) Run() ([]string, error) {
	for _, file := range self.UploadFiles {
		self.uploader.Upload(file)
	}
	return self.UploadFiles, nil
}

type Uploader interface {
	Upload(filePath string)
}

type S3Uploader struct {
	bucket *s3.Bucket
	contentType string
}

func NewS3Uploader(bucket, contentType string) *S3Uploader {
	auth := aws.Auth{conf.Settings.Auth.AccessKey, conf.Settings.Auth.SecretKey}
	s3Conn := s3.New(auth, aws.USEast)
	return &S3Uploader{
		bucket: s3Conn.Bucket(bucket),
		contentType: contentType,
	}
}

func GeneratePathPrefix(fileInfo os.FileInfo) string {
	modifyTime := fileInfo.ModTime()
	return fmt.Sprintf("%04d/%02d/%02d/%02d", modifyTime.Year(), modifyTime.Month(), modifyTime.Day(), modifyTime.Hour())
}

//this method will upload to s3 based on the key strategy
func (self *S3Uploader) Upload(filePath string) {
	fileReader, err := util.Fs.Open(filePath)
	util.LogPanic(err)

	fileInfo, _ := fileReader.Stat()
	pathPrefix := GeneratePathPrefix(fileInfo)
	key := path.Join(pathPrefix, path.Base(filePath))

	err = self.bucket.PutReader(key, fileReader, fileInfo.Size(), self.contentType, s3.Private)
	util.LogPanic(err)
}
