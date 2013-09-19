package uploadthis

import (
	"launchpad.net/goamz/s3"
	. "launchpad.net/gocheck"
)

type S3Suite struct{}

var _ = Suite(&S3Suite{})

func (s *S3Suite) TestUploadBuffer(c *C) {
	myS3 := new(s3.S3)
	err := UploadBuffer(myS3, "bucket", "path", []byte{})
	c.Assert(err, Not(Equals), nil)
}
