package hooks

import (
	"compress/gzip"
	"github.com/imosquera/uploadthis/testhelper"
	"io"
	"io/ioutil"
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct {
	InFile   *os.File
	FileStat os.FileInfo
}

var _ = Suite(&TestSuite{})

func (s *TestSuite) SetUpTest(c *C) {
	s.InFile, _ = os.Open("../fixtures/gziptest.txt")
	s.FileStat, _ = s.InFile.Stat()
}

func (s *TestSuite) TearDownTest(c *C) {
	s.InFile = nil
	s.FileStat = nil
}

//test my gzip compressions
func (s *TestSuite) TestCompressionPrehook(c *C) {
	//create a compression prehook and run it
	compress := CompressPrehook{}
	reader, _ := compress.RunPrehook(s.InFile, s.FileStat)

	//read the return bytes and using it to test the difference from source
	gzipReader, _ := gzip.NewReader(reader)
	allBytes, _ := ioutil.ReadAll(gzipReader)
	textFile := string(allBytes)

	s.InFile.Seek(0, 0)
	sourceBytes, _ := ioutil.ReadAll(s.InFile)
	sourceString := string(sourceBytes)
	c.Check(sourceString, Equals, textFile)
}

func (s *TestSuite) TestCompressionPrehookFail(c *C) {

	defer testhelper.Patch(&Copy, func(io.Writer, io.Reader) (int64, error) {
		return 0, nil
	}).Restore()

	compress := CompressPrehook{}
	_, err := compress.RunPrehook(s.InFile, s.FileStat)
	c.Check("Bytes written does not match InFile byte size", Equals, err.Error())
	compress.RunPrehook(s.InFile, s.FileStat)
}
