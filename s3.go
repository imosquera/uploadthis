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
