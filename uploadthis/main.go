package main

import (
	"github.com/imosquera/uploadthis"
)

func main() {
	uploadthis.ParseOpts()
	uploadthis.UploadFile("test")
}
