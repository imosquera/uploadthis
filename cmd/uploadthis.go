package main

import "fmt"
import "github.com/imosquera/uploadthis"

func main() {
	uploadthis.LoadConfig("sample-config.yaml")
	fmt.Println("####")
	println(uploadthis.Settings.WatchDir)
}
