package uploadthis

import (
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"log"
)

type S3Mocker interface {
	Bucket(name string) BucketMocker
}
type BucketMocker interface {
	Name string
	Get(path string) (data []byte, err error)
	Put(path string, data []byte, contType string, perm ACL) error
}

var s3Conn *s3.S3

var GetS3 = func() *s3.S3 {
	if s3Conn == nil {
		auth := aws.Auth{Settings.Auth.AccessKey, Settings.Auth.SecretKey}
		s3Conn = s3.New(auth, aws.USEast)
	}
	return s3Conn
}

func UploadBuffer(bucket string, path string, data []byte) {
	s := GetS3()
	b := s.Bucket(bucket)
	err := b.Put(path, data, "content-type", s3.Private)
	if err != nil {
		log.Panic(err)
	}
}

func DownloadBuffer(bucket string, path string) []byte {
	s := GetS3()
	b := s.Bucket(bucket)
	data, err := b.Get(path)
	if err != nil {
		log.Panic(err)
		return []byte{}
	} else {
		return data
	}
	return []byte{}
}
