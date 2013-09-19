package uploadthis

import (
	. "launchpad.net/gocheck"
)

type S3Suite struct{}

var _ = Suite(&S3Suite{})

func (s *S3Suite) TestS3(c *C) {
	bucket := "pointabout"
	path := "testS3"
	data := []byte{'h', 'e', 'l', 'l', 'o'}
	UploadBuffer(bucket, path, data)
	returnData := DownloadBuffer(bucket, path)
	c.Assert(string(returnData), Equals, string(data))
	c.Assert("a", Equals, "a")
}
