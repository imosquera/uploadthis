package uploadthis

import (
	"errors"
	"testing"
)

func TestParseOpts(t *testing.T) {
	optsParser = func(interface{}) ([]string, error) {
		println("DID CALL")
		return []string{}, errors.New("")
	}
	ParseOpts()
}
