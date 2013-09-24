package main

import (
	"github.com/imosquera/uploadthis"
	"fmt"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"

)

func main() {
	ACCESS := "x"
	SECRET := "y"
        REGION := aws.USEast
	BUCKET := "myBucket"
	FILENAME := "myPath"

	auth := aws.Auth{ACCESS, SECRET}
	s3Conn := s3.New(auth, REGION)	
	fmt.Println(uploadthis.UploadBuffer(s3Conn, BUCKET, FILENAME, []byte{'t','e','s','t'}))
	fmt.Println(uploadthis.DownloadBuffer(s3Conn, BUCKET, FILENAME))
}

