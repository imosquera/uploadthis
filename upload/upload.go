package upload

import (
	"fmt"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/conf"
	//"github.com/imosquera/uploadthis/aws/iam"
	"github.com/imosquera/uploadthis/util"
	//"launchpad.net/goamz/aws"
	//"launchpad.net/goamz/s3"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
	"os"
	"path"
	"time"
	log "github.com/cihub/seelog"
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
	exptdate := time.Now()
	auth, err := aws.GetAuth(conf.Settings.Auth.AccessKey, conf.Settings.Auth.SecretKey, "", exptdate)
	util.LogPanic(err)

	//Mon Jan 2 15:04:05 MST 2006 (MST is GMT-0700)
	//"2014-01-30T21:31:47Z"
	//exp, err := time.Parse("2006-01-02T15:04:05MST", "2014-01-30T21:31:47Z")
	//util.LogPanic(err)
	//auth := aws.Auth{"ASIAJHKSBHMMUFFPORVA", "9hCnULX/nfQlNZg0b5lp1EMBjHPiOjVg+EoL//3b", "AQoDYXdzENn//////////wEa8AI98/cg/ATURlliR1p6P6Enm83cPIkgC7J4tTOLQjG7TJAEATNSuHxZCxd/1AmNm79iShElr3Ejiy2HigZQ4RRcTEtWP5ZE44L45t9CbroUhUgmZvhzkEqUAkUgOk6srFQ+5WjqMX4e8K/5a/iBBRt2o25/Nb/HRgCVOuVEZvnSkZuTkLKzdZQSqpAR/jEFu/C+AxJgB4OWbwGpYRUL2HrvVsNMLNVz4yxJoX0NTFlFzJK3affLnc3X+SdR0WCMIgOAjweau98IClPVTL7yPtbElne/j8gtIKmD9zr//o391zx+5oqunCsGJn2Tc2qWHjyJgkiQiivrT0r8QN1afToR9GmrgWd50KN8FCgjrauxHJTltzeyRSIelASfnwRMR/+gfpy3NcUX0IhP+ZpHIWfpVyc7bW82QA/5S546HCJMu/w48T4cyszrQAjiSCnhoinHNJ11kcIV1lRuItpo7vl2J37TrE6nsbYFBRa32BBXjCCZ3qmXBQ==", time.Now()}

	s3Conn := s3.New(auth, aws.USEast)
	return &S3Uploader{
		bucket: s3Conn.Bucket(bucket),
		contentType: contentType,
	}
}

func GeneratePathPrefix(fileInfo os.FileInfo) string {
	modifyTime := fileInfo.ModTime().UTC()
	return fmt.Sprintf("%04d/%02d/%02d/%02d", modifyTime.Year(), modifyTime.Month(), modifyTime.Day(), modifyTime.Hour())
}

//this method will upload to s3 based on the key strategy
func (self *S3Uploader) Upload(filePath string) {
	fileReader, err := util.Fs.Open(filePath)
	util.LogPanic(err)

	fileInfo, _ := fileReader.Stat()
	pathPrefix := GeneratePathPrefix(fileInfo)
	key := path.Join(pathPrefix, path.Base(filePath))
	log.Info("Upload to: ", key)

	err = self.bucket.PutReader(key, fileReader, fileInfo.Size(), self.contentType, s3.Private, s3.Options{})
	util.LogPanic(err)
}
