package hooks

import (
	"compress/gzip"
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/util"
	"io"
	"log"
	"os"
	"path"
)

var registeredPrehooks map[string]commands.Commander

func RegisterPrehook(name string, prehook commands.Commander) {

	if registeredPrehooks == nil {
		registeredPrehooks = make(map[string]commands.Commander, 5)
	}
	registeredPrehooks[name] = prehook
}

// func RunPrehooks(uploadFiles []string, monitorDir conf.MonitorDir) []string {
// 	prehooks := getPrehooks(monitorDir.PreHooks)
// 	for name, prehook := range prehooks {
// 		workDir := path.Join(monitorDir.Path, name)
// 		uploadFiles = MoveToWorkDir(workDir, uploadFiles)
// 		//update uploadfiles with resume files
// 		uploadFiles = updateResumeFiles(workDir, uploadFiles)
// 		uploadFiles, _ = prehook.RunPrehook(uploadFiles)
// 	}
// 	return uploadFiles
// }

type Prehook struct{}

func GetPrehookCommands(prehooks []string, prehookCommands map[string]commands.Commander) {
	for _, prehook := range prehooks {
		prehookCommands[prehook] = registeredPrehooks[prehook]
	}
}

var compressFile = func(filepath string) (string, error) {
	inFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Error for file " + filepath + " " + err.Error())
	}
	log.Println("Working on file", filepath)

	gzipPath := path.Join(path.Dir(filepath), path.Base(filepath)+".gz")
	outFile, err := os.Create(gzipPath)
	util.LogPanic(err)

	gzipWriter, err := gzip.NewWriterLevel(outFile, gzip.BestCompression)
	util.LogPanic(err)

	_, err = io.Copy(gzipWriter, inFile)
	util.LogPanic(err)

	gzipWriter.Close()

	return gzipPath, err
}

type CompressPrehook struct {
	*Prehook
}

//this function will compress the infile and
//return a file pointer that is readible
func (c CompressPrehook) RunPrehook(uploadFiles []string) ([]string, error) {
	var err error
	newFiles := make([]string, len(uploadFiles))
	for _, uploadFile := range uploadFiles {
		uploadFile, _ := compressFile(uploadFile)
		newFiles = append(newFiles, uploadFile)
	}
	return newFiles, err
}
