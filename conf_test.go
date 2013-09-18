package uploadthis

import (
	"testing"
)

func TestParseOptsAccessKey(t *testing.T) {
	optsParser = func(interface{}) ([]string, error) {
		opts.AccesssKey = "FAKE KEY"
		opts.SecretKey = ""
		return []string{}, nil
	}
	ParseOpts()
}
