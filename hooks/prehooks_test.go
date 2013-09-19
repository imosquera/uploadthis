package hooks

import (
	"compress/gzip"
	"github.com/imosquera/uploadthis/hooks"
	"io/ioutil"
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

// test my gzip compressions
func (s *TestSuite) TestCompressionPrehook(c *C) {
	inFile, _ := os.Open("../fixtures/gziptest.txt")
	fileStat, _ := inFile.Stat()

	//create a compression prehook and run it
	compress := hooks.CompressPrehook{}
	reader := compress.RunPrehook(inFile, fileStat)

	//read the return bytes and using it to test the difference from source
	gzipReader, _ := gzip.NewReader(reader)
	allBytes, _ := ioutil.ReadAll(gzipReader)
	textFile := string(allBytes)

	inFile.Seek(0, 0)
	sourceBytes, _ := ioutil.ReadAll(inFile)
	sourceString := string(sourceBytes)
	c.Check(sourceString, Equals, textFile)
}

func (s *TestSuite) TestCompressionPrehookFail(c *C) {
 defer testinghelpers.Patch(mypkg.SomeMethod, func(*somepkg.SomeType, args) { 
        mock method 
    }).Restore() 
}
