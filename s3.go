package uploadthis

import (
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
)

var s3Conn *s3.S3

func GetS3() *s3.S3 {
	if s3Conn == nil {
		auth := aws.Auth{Settings.Auth.AccessKey, Settings.Auth.SecretKey}
		s3Conn = s3.New(auth, aws.USEast)
	}
	return s3Conn
}

func UploadBuffer(bucket string, path string, data []byte) {
	s := GetS3()
        b := s.Bucket(bucket)
        b.Put(path, data, "content-type", s3.Private)
}

func DownloadBuffer(bucket string, path string) []byte {
	s := GetS3()
        b := s.Bucket(bucket)
	data, err := b.Get(path)
        if err != nil {
                return []byte{}
        } else {
                return data
        }
	return []byte{}
}
