package hooks

import (
	"github.com/imosquera/uploadthis/commands"
	"github.com/imosquera/uploadthis/util"
	log "github.com/cihub/seelog"
	"time"
	"os"
)

func NewArchiveCommand() *ArchiveCommand {
	return &ArchiveCommand{
		commands.NewFileStateCommand(),
	}
}

//The archive command moves files into the archive folder for later deletion from some other process
type ArchiveCommand struct {
	*commands.Command
}

func (self *ArchiveCommand) Run() ([]string, error) {
	threshold := self.GetMonitor().CleanupThreshold
	if threshold >= 0 {
		expirationThreshold := time.Now().Unix() - int64(threshold * 60 * 60)
		var err error
		newFiles := make([]string, 0, len(self.UploadFiles))
		for _, arcFilePath := range self.UploadFiles {
			arcFile, err := os.Stat(arcFilePath)
			util.LogPanic(err)
			if !arcFile.IsDir() && arcFile.ModTime().Unix() < expirationThreshold {
				log.Infof("%s is too old (%s). Deleting", arcFilePath, arcFile.ModTime().Format("2006-01-02 15:04:05"))
				err := os.Remove(arcFilePath)
				util.LogPanic(err)
			} else {
				newFiles = append(newFiles, arcFilePath)
			}
		}
		return newFiles, err
	}
	return self.UploadFiles, nil
}
