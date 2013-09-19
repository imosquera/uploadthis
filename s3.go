package uploadthis

import (
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
)

var s3Conn *s3.S3

var GetS3 = func() *s3.S3 {
	if s3Conn == nil {
		auth := aws.Auth{Settings.Auth.AccessKey, Settings.Auth.SecretKey}
		s3Conn = s3.New(auth, aws.USEast)
	}
	return s3Conn
}

func UploadBuffer(s *s3.S3, bucket string, path string, data []byte) error {
	b := s.Bucket(bucket)
	err := b.Put(path, data, "content-type", s3.Private)
	if err != nil {
		log.Println(err)
	}
	return err
}

func DownloadBuffer(s *s3.S3, bucket string, path string) ([]byte, error) {
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
