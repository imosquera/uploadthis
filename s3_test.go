package uploadthis

import (
	"code.google.com/p/gomock/gomock"
	"launchpad.net/goamz/s3" //mock
	. "launchpad.net/gocheck"
)

type S3Suite struct{}

var _ = Suite(&S3Suite{})

func (s *S3Suite) TestUploadBuffer(c *C) {
	mockCtrl := gomock.NewController(c)
	defer mockCtrl.Finish()

	s3.MOCK().SetController(mockCtrl)

	kk := new(s3.S3)
	//println(kk)

	// // // We expect to see HandyMethod called, and we want to return true
	b := &s3.Bucket{}
	kk.EXPECT().Bucket("bucket").Return(b)
	b.EXPECT().Put("path", []byte{}, "content-type", s3.Private)

	xxkk := UploadBuffer(kk, "bucket", "path", []byte{})
	println(xxkk)

	// c.Assert(err, Not(Equals), nil)
}
