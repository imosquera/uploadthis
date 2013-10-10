package commands

type Commander interface {
	SetUploadFiles(uploadFiles []string)
	Prepare(workDir string)
	Run() ([]string, error)
}

type Command struct {
	UploadFiles []string
}

func (self *Command) SetUploadFiles(uploadFiles []string) {
	self.UploadFiles = uploadFiles
}

func (self *Command) Prepare(workDir string) {
	//uploadFiles = MoveToWorkDir(workDir, uploadFiles)
	//paths := monitor.GetUploadFiles(workDir)
	//uploadFiles = append(uploadFiles, paths...)
	//uploadFiles, _ = command.Run(uploadFiles)
}
