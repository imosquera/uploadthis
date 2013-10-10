package upload

import (
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/conf"
	"io"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
)

type UploadCommand struct {
	commands.Command
	uploader Uploader
}

func (self *UploadCommand) Run() ([]string, error) {
	//uploadfiles
	for _, file := range self.UploadFiles {
		println(file)
	}
	return []string{}, nil
}

type Uploader interface {
	Upload(bucket string, path string, data io.Reader, length int64) error
}

type S3Uploader struct{}

var s3Conn *s3.S3
var GetS3 = func() *s3.S3 {
	if s3Conn == nil {
		auth := aws.Auth{conf.Settings.Auth.AccessKey, conf.Settings.Auth.SecretKey}
		s3Conn = s3.New(auth, aws.USEast)
	}
	return s3Conn
}

func UploadBytes(s *s3.S3, bucket string, path string, data []byte) error {
	b := s.Bucket(bucket)
	err := b.Put(path, data, "content-type", s3.Private)
	if err != nil {
		log.Println(err)
	}
	return err
}

func UploadReader(s *s3.S3, bucket string, path string, data io.Reader, length int64) error {
	b := s.Bucket(bucket)
	err := b.PutReader(path, data, length, "content-type", s3.Private)
	if err != nil {
		log.Println(err)
	}
	return err
}

/*
func DownloadBytes(s *s3.S3, bucket string, path string) ([]byte, error) {
	b := s.Bucket(bucket)
	data, err := b.Get(path)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	} else {
		return data, err
	}
	return []byte{}, err
}

func DownloadReader(s *s3.S3, bucket string, path string) (io.ReadCloser, error) {
	b := s.Bucket(bucket)
	data, err := b.GetReader(path)
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		return data, err
	}
	return nil, err
}*/
