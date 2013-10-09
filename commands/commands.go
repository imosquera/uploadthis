package commands

type Commander interface {
	SetUploadFiles(uploadFiles []string)
	Prepare(workDir string)
	Run() ([]string, error)
}
